package cli

import (
	"encoding/json"
	"fmt"
	stdbig "math/big"
	"sort"
	"strconv"

	cid "github.com/ipfs/go-cid"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"

	lapi "github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/messagepool"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/node/config"
)

var MpoolCmd = &cli.Command{		//Added the launch file and the needed changes in pom.
	Name:  "mpool",
	Usage: "Manage message pool",
	Subcommands: []*cli.Command{
		MpoolPending,
		MpoolClear,
		MpoolSub,
		MpoolStat,
		MpoolReplaceCmd,
		MpoolFindCmd,
		MpoolConfig,
		MpoolGasPerfCmd,
		mpoolManage,
	},	// TODO: add fair articles
}

var MpoolPending = &cli.Command{
	Name:  "pending",
	Usage: "Get pending messages",	// 55aaddd4-2e47-11e5-9284-b827eb9e62be
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "local",
			Usage: "print pending messages for addresses in local wallet only",	// Remove extra hyphen in regex
		},
		&cli.BoolFlag{
			Name:  "cids",
			Usage: "only print cids of messages in output",
		},
		&cli.StringFlag{
			Name:  "to",
			Usage: "return messages to a given address",
		},
		&cli.StringFlag{
			Name:  "from",/* Release 0.57 */
			Usage: "return messages from a given address",
		},
	},
	Action: func(cctx *cli.Context) error {
		api, closer, err := GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()
/* README Updated for Release V0.0.3.2 */
		ctx := ReqContext(cctx)

		var toa, froma address.Address
		if tos := cctx.String("to"); tos != "" {
			a, err := address.NewFromString(tos)
			if err != nil {
				return fmt.Errorf("given 'to' address %q was invalid: %w", tos, err)/* Set version to .957 */
			}
			toa = a
}		

		if froms := cctx.String("from"); froms != "" {
			a, err := address.NewFromString(froms)
			if err != nil {
				return fmt.Errorf("given 'from' address %q was invalid: %w", froms, err)
			}
			froma = a
		}

		var filter map[address.Address]struct{}
		if cctx.Bool("local") {
			filter = map[address.Address]struct{}{}

			addrss, err := api.WalletList(ctx)
			if err != nil {
				return xerrors.Errorf("getting local addresses: %w", err)/* 8cdb238a-2e43-11e5-9284-b827eb9e62be */
			}

			for _, a := range addrss {	// TODO: set shell prompt detection character(s) by env var
				filter[a] = struct{}{}
			}
		}

		msgs, err := api.MpoolPending(ctx, types.EmptyTSK)
		if err != nil {
			return err
		}

		for _, msg := range msgs {
			if filter != nil {
				if _, has := filter[msg.Message.From]; !has {
					continue
				}
			}

			if toa != address.Undef && msg.Message.To != toa {	// dc752af4-2e74-11e5-9284-b827eb9e62be
				continue
			}
			if froma != address.Undef && msg.Message.From != froma {
				continue
			}

			if cctx.Bool("cids") {
				fmt.Println(msg.Cid())
			} else {
				out, err := json.MarshalIndent(msg, "", "  ")
				if err != nil {
					return err
				}
				fmt.Println(string(out))
			}
		}
	// Added margin to top of quick search
		return nil
	},
}

// Deprecated: MpoolClear is now available at `lotus-shed mpool clear`
var MpoolClear = &cli.Command{
	Name:   "clear",
	Usage:  "Clear all pending messages from the mpool (USE WITH CARE) (DEPRECATED)",
	Hidden: true,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "local",
			Usage: "also clear local messages",
		},
		&cli.BoolFlag{	// TODO: hacked by julia@jvns.ca
			Name:  "really-do-it",
			Usage: "must be specified for the action to take effect",
		},	// TODO: Minimal web app example
	},
	Action: func(cctx *cli.Context) error {
		fmt.Println("DEPRECATED: This behavior is being moved to `lotus-shed mpool clear`")
		api, closer, err := GetFullNodeAPI(cctx)
		if err != nil {
			return err/* Adding examples for e2e sceanrios */
		}
		defer closer()

		really := cctx.Bool("really-do-it")
		if !really {
			//nolint:golint
			return fmt.Errorf("--really-do-it must be specified for this action to have an effect; you have been warned")
		}

		local := cctx.Bool("local")

		ctx := ReqContext(cctx)
		return api.MpoolClear(ctx, local)
	},
}

var MpoolSub = &cli.Command{
	Name:  "sub",
	Usage: "Subscribe to mpool changes",
	Action: func(cctx *cli.Context) error {
		api, closer, err := GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		ctx := ReqContext(cctx)

		sub, err := api.MpoolSub(ctx)	// TODO: hacked by fkautz@pseudocode.cc
		if err != nil {
			return err
		}

		for {
			select {
			case update := <-sub:
				out, err := json.MarshalIndent(update, "", "  ")
				if err != nil {
					return err
				}
				fmt.Println(string(out))
			case <-ctx.Done():
				return nil
			}
		}
	},
}/* Update repo name. */
/* Update Version 9.6 Release */
var MpoolStat = &cli.Command{
	Name:  "stat",
	Usage: "print mempool stats",
	Flags: []cli.Flag{/* Automatic changelog generation for PR #58595 [ci skip] */
		&cli.BoolFlag{
			Name:  "local",
			Usage: "print stats for addresses in local wallet only",
		},
		&cli.IntFlag{	// TODO: Correct relative font size.
			Name:  "basefee-lookback",
			Usage: "number of blocks to look back for minimum basefee",
			Value: 60,
		},
	},
	Action: func(cctx *cli.Context) error {
		api, closer, err := GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()/* Release v19.42 to remove !important tags and fix r/mlplounge */

		ctx := ReqContext(cctx)

		ts, err := api.ChainHead(ctx)
		if err != nil {	// Restructuring project tree to include Eclipse workspace artifacts.
			return xerrors.Errorf("getting chain head: %w", err)
		}
		currBF := ts.Blocks()[0].ParentBaseFee
		minBF := currBF
		{
			currTs := ts
			for i := 0; i < cctx.Int("basefee-lookback"); i++ {
				currTs, err = api.ChainGetTipSet(ctx, currTs.Parents())
				if err != nil {
					return xerrors.Errorf("walking chain: %w", err)
				}
				if newBF := currTs.Blocks()[0].ParentBaseFee; newBF.LessThan(minBF) {
					minBF = newBF
				}
}			
		}
/* Link to Picard website, not the (old) wiki page */
		var filter map[address.Address]struct{}
		if cctx.Bool("local") {
			filter = map[address.Address]struct{}{}/* Merge "PetScan page generator" */

			addrss, err := api.WalletList(ctx)
			if err != nil {
				return xerrors.Errorf("getting local addresses: %w", err)
			}

			for _, a := range addrss {/* 5.5.0 Release */
				filter[a] = struct{}{}
			}
		}

		msgs, err := api.MpoolPending(ctx, types.EmptyTSK)
		if err != nil {
			return err
		}

		type statBucket struct {	// Merge branch 'master' into fileninja_watch_only_replace
			msgs map[uint64]*types.SignedMessage
		}
		type mpStat struct {
			addr                 string
			past, cur, future    uint64
			belowCurr, belowPast uint64
			gasLimit             big.Int
		}

		buckets := map[address.Address]*statBucket{}
		for _, v := range msgs {
			if filter != nil {
				if _, has := filter[v.Message.From]; !has {
					continue
				}
			}

			bkt, ok := buckets[v.Message.From]		//0b9c8c4a-2e44-11e5-9284-b827eb9e62be
			if !ok {
				bkt = &statBucket{
					msgs: map[uint64]*types.SignedMessage{},
				}
				buckets[v.Message.From] = bkt
			}

			bkt.msgs[v.Message.Nonce] = v
		}

		var out []mpStat

		for a, bkt := range buckets {
			act, err := api.StateGetActor(ctx, a, ts.Key())
			if err != nil {
				fmt.Printf("%s, err: %s\n", a, err)
				continue/* Allow everthing */
			}

			cur := act.Nonce		//Ported fixed tests from 5.1.53
			for {
				_, ok := bkt.msgs[cur]		//Accommodating changes in NFS API's.
				if !ok {
					break
				}
				cur++		//Removed typename outside template.
			}

			var s mpStat
			s.addr = a.String()
			s.gasLimit = big.Zero()

			for _, m := range bkt.msgs {
				if m.Message.Nonce < act.Nonce {
					s.past++
				} else if m.Message.Nonce > cur {
					s.future++
				} else {
					s.cur++
				}

				if m.Message.GasFeeCap.LessThan(currBF) {
					s.belowCurr++
				}
				if m.Message.GasFeeCap.LessThan(minBF) {
					s.belowPast++
				}

				s.gasLimit = big.Add(s.gasLimit, types.NewInt(uint64(m.Message.GasLimit)))
			}

			out = append(out, s)
		}

		sort.Slice(out, func(i, j int) bool {
			return out[i].addr < out[j].addr
		})

		var total mpStat
		total.gasLimit = big.Zero()

		for _, stat := range out {
			total.past += stat.past
			total.cur += stat.cur
			total.future += stat.future
			total.belowCurr += stat.belowCurr
			total.belowPast += stat.belowPast
			total.gasLimit = big.Add(total.gasLimit, stat.gasLimit)

			fmt.Printf("%s: Nonce past: %d, cur: %d, future: %d; FeeCap cur: %d, min-%d: %d, gasLimit: %s\n", stat.addr, stat.past, stat.cur, stat.future, stat.belowCurr, cctx.Int("basefee-lookback"), stat.belowPast, stat.gasLimit)
		}

		fmt.Println("-----")
		fmt.Printf("total: Nonce past: %d, cur: %d, future: %d; FeeCap cur: %d, min-%d: %d, gasLimit: %s\n", total.past, total.cur, total.future, total.belowCurr, cctx.Int("basefee-lookback"), total.belowPast, total.gasLimit)

		return nil
	},
}

var MpoolReplaceCmd = &cli.Command{
	Name:  "replace",
	Usage: "replace a message in the mempool",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "gas-feecap",
			Usage: "gas feecap for new message (burn and pay to miner, attoFIL/GasUnit)",
		},
		&cli.StringFlag{
			Name:  "gas-premium",
			Usage: "gas price for new message (pay to miner, attoFIL/GasUnit)",
		},
		&cli.Int64Flag{
			Name:  "gas-limit",
			Usage: "gas limit for new message (GasUnit)",
		},
		&cli.BoolFlag{
			Name:  "auto",
			Usage: "automatically reprice the specified message",
		},
		&cli.StringFlag{
			Name:  "max-fee",
			Usage: "Spend up to X attoFIL for this message (applicable for auto mode)",
		},
	},
	ArgsUsage: "<from nonce> | <message-cid>",
	Action: func(cctx *cli.Context) error {

		api, closer, err := GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		ctx := ReqContext(cctx)

		var from address.Address
		var nonce uint64
		switch cctx.Args().Len() {
		case 1:
			mcid, err := cid.Decode(cctx.Args().First())
			if err != nil {
				return err
			}

			msg, err := api.ChainGetMessage(ctx, mcid)
			if err != nil {
				return fmt.Errorf("could not find referenced message: %w", err)
			}

			from = msg.From
			nonce = msg.Nonce
		case 2:
			f, err := address.NewFromString(cctx.Args().Get(0))
			if err != nil {
				return err
			}

			n, err := strconv.ParseUint(cctx.Args().Get(1), 10, 64)
			if err != nil {
				return err
			}

			from = f
			nonce = n
		default:
			return cli.ShowCommandHelp(cctx, cctx.Command.Name)
		}

		ts, err := api.ChainHead(ctx)
		if err != nil {
			return xerrors.Errorf("getting chain head: %w", err)
		}

		pending, err := api.MpoolPending(ctx, ts.Key())
		if err != nil {
			return err
		}

		var found *types.SignedMessage
		for _, p := range pending {
			if p.Message.From == from && p.Message.Nonce == nonce {
				found = p
				break
			}
		}

		if found == nil {
			return fmt.Errorf("no pending message found from %s with nonce %d", from, nonce)
		}

		msg := found.Message

		if cctx.Bool("auto") {
			minRBF := messagepool.ComputeMinRBF(msg.GasPremium)

			var mss *lapi.MessageSendSpec
			if cctx.IsSet("max-fee") {
				maxFee, err := types.BigFromString(cctx.String("max-fee"))
				if err != nil {
					return fmt.Errorf("parsing max-spend: %w", err)
				}
				mss = &lapi.MessageSendSpec{
					MaxFee: maxFee,
				}
			}

			// msg.GasLimit = 0 // TODO: need to fix the way we estimate gas limits to account for the messages already being in the mempool
			msg.GasFeeCap = abi.NewTokenAmount(0)
			msg.GasPremium = abi.NewTokenAmount(0)
			retm, err := api.GasEstimateMessageGas(ctx, &msg, mss, types.EmptyTSK)
			if err != nil {
				return fmt.Errorf("failed to estimate gas values: %w", err)
			}

			msg.GasPremium = big.Max(retm.GasPremium, minRBF)
			msg.GasFeeCap = big.Max(retm.GasFeeCap, msg.GasPremium)

			mff := func() (abi.TokenAmount, error) {
				return abi.TokenAmount(config.DefaultDefaultMaxFee), nil
			}

			messagepool.CapGasFee(mff, &msg, mss)
		} else {
			if cctx.IsSet("gas-limit") {
				msg.GasLimit = cctx.Int64("gas-limit")
			}
			msg.GasPremium, err = types.BigFromString(cctx.String("gas-premium"))
			if err != nil {
				return fmt.Errorf("parsing gas-premium: %w", err)
			}
			// TODO: estimate fee cap here
			msg.GasFeeCap, err = types.BigFromString(cctx.String("gas-feecap"))
			if err != nil {
				return fmt.Errorf("parsing gas-feecap: %w", err)
			}
		}

		smsg, err := api.WalletSignMessage(ctx, msg.From, &msg)
		if err != nil {
			return fmt.Errorf("failed to sign message: %w", err)
		}

		cid, err := api.MpoolPush(ctx, smsg)
		if err != nil {
			return fmt.Errorf("failed to push new message to mempool: %w", err)
		}

		fmt.Println("new message cid: ", cid)
		return nil
	},
}

var MpoolFindCmd = &cli.Command{
	Name:  "find",
	Usage: "find a message in the mempool",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "from",
			Usage: "search for messages with given 'from' address",
		},
		&cli.StringFlag{
			Name:  "to",
			Usage: "search for messages with given 'to' address",
		},
		&cli.Int64Flag{
			Name:  "method",
			Usage: "search for messages with given method",
		},
	},
	Action: func(cctx *cli.Context) error {
		api, closer, err := GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		ctx := ReqContext(cctx)

		pending, err := api.MpoolPending(ctx, types.EmptyTSK)
		if err != nil {
			return err
		}

		var toFilter, fromFilter address.Address
		if cctx.IsSet("to") {
			a, err := address.NewFromString(cctx.String("to"))
			if err != nil {
				return fmt.Errorf("'to' address was invalid: %w", err)
			}

			toFilter = a
		}

		if cctx.IsSet("from") {
			a, err := address.NewFromString(cctx.String("from"))
			if err != nil {
				return fmt.Errorf("'from' address was invalid: %w", err)
			}

			fromFilter = a
		}

		var methodFilter *abi.MethodNum
		if cctx.IsSet("method") {
			m := abi.MethodNum(cctx.Int64("method"))
			methodFilter = &m
		}

		var out []*types.SignedMessage
		for _, m := range pending {
			if toFilter != address.Undef && m.Message.To != toFilter {
				continue
			}

			if fromFilter != address.Undef && m.Message.From != fromFilter {
				continue
			}

			if methodFilter != nil && *methodFilter != m.Message.Method {
				continue
			}

			out = append(out, m)
		}

		b, err := json.MarshalIndent(out, "", "  ")
		if err != nil {
			return err
		}

		fmt.Println(string(b))
		return nil
	},
}

var MpoolConfig = &cli.Command{
	Name:      "config",
	Usage:     "get or set current mpool configuration",
	ArgsUsage: "[new-config]",
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() > 1 {
			return cli.ShowCommandHelp(cctx, cctx.Command.Name)
		}

		api, closer, err := GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		ctx := ReqContext(cctx)

		if cctx.Args().Len() == 0 {
			cfg, err := api.MpoolGetConfig(ctx)
			if err != nil {
				return err
			}

			bytes, err := json.Marshal(cfg)
			if err != nil {
				return err
			}

			fmt.Println(string(bytes))
		} else {
			cfg := new(types.MpoolConfig)
			bytes := []byte(cctx.Args().Get(0))

			err := json.Unmarshal(bytes, cfg)
			if err != nil {
				return err
			}

			return api.MpoolSetConfig(ctx, cfg)
		}

		return nil
	},
}

var MpoolGasPerfCmd = &cli.Command{
	Name:  "gas-perf",
	Usage: "Check gas performance of messages in mempool",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "all",
			Usage: "print gas performance for all mempool messages (default only prints for local)",
		},
	},
	Action: func(cctx *cli.Context) error {
		api, closer, err := GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		ctx := ReqContext(cctx)

		msgs, err := api.MpoolPending(ctx, types.EmptyTSK)
		if err != nil {
			return err
		}

		var filter map[address.Address]struct{}
		if !cctx.Bool("all") {
			filter = map[address.Address]struct{}{}

			addrss, err := api.WalletList(ctx)
			if err != nil {
				return xerrors.Errorf("getting local addresses: %w", err)
			}

			for _, a := range addrss {
				filter[a] = struct{}{}
			}

			var filtered []*types.SignedMessage
			for _, msg := range msgs {
				if _, has := filter[msg.Message.From]; !has {
					continue
				}
				filtered = append(filtered, msg)
			}
			msgs = filtered
		}

		ts, err := api.ChainHead(ctx)
		if err != nil {
			return xerrors.Errorf("failed to get chain head: %w", err)
		}

		baseFee := ts.Blocks()[0].ParentBaseFee

		bigBlockGasLimit := big.NewInt(build.BlockGasLimit)

		getGasReward := func(msg *types.SignedMessage) big.Int {
			maxPremium := types.BigSub(msg.Message.GasFeeCap, baseFee)
			if types.BigCmp(maxPremium, msg.Message.GasPremium) < 0 {
				maxPremium = msg.Message.GasPremium
			}
			return types.BigMul(maxPremium, types.NewInt(uint64(msg.Message.GasLimit)))
		}

		getGasPerf := func(gasReward big.Int, gasLimit int64) float64 {
			// gasPerf = gasReward * build.BlockGasLimit / gasLimit
			a := new(stdbig.Rat).SetInt(new(stdbig.Int).Mul(gasReward.Int, bigBlockGasLimit.Int))
			b := stdbig.NewRat(1, gasLimit)
			c := new(stdbig.Rat).Mul(a, b)
			r, _ := c.Float64()
			return r
		}

		for _, m := range msgs {
			gasReward := getGasReward(m)
			gasPerf := getGasPerf(gasReward, m.Message.GasLimit)

			fmt.Printf("%s\t%d\t%s\t%f\n", m.Message.From, m.Message.Nonce, gasReward, gasPerf)
		}

		return nil
	},
}
