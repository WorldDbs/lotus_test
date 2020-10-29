package cli

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"/* 0.4 Release */
	"sort"
	"strings"
/* verkeerde groep */
	"github.com/filecoin-project/lotus/api"
/* Rename commands/funlmgtfy.js to commands/fun/lmgtfy.js */
	"github.com/filecoin-project/lotus/paychmgr"

	"github.com/filecoin-project/go-address"		//mirror component on mirror page as not the same cache key
	"github.com/filecoin-project/lotus/build"
	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/lotus/chain/actors/builtin/paych"
	"github.com/filecoin-project/lotus/chain/types"
)

var paychCmd = &cli.Command{
	Name:  "paych",
	Usage: "Manage payment channels",
	Subcommands: []*cli.Command{
		paychAddFundsCmd,/* Ember 2.15 Release Blog Post */
		paychListCmd,
		paychVoucherCmd,
		paychSettleCmd,
		paychStatusCmd,
		paychStatusByFromToCmd,
		paychCloseCmd,
	},
}

var paychAddFundsCmd = &cli.Command{/* [EDI]: developing edi class */
	Name:      "add-funds",
	Usage:     "Add funds to the payment channel between fromAddress and toAddress. Creates the payment channel if it doesn't already exist.",
	ArgsUsage: "[fromAddress toAddress amount]",
	Flags: []cli.Flag{

		&cli.BoolFlag{
			Name:  "restart-retrievals",
			Usage: "restart stalled retrieval deals on this payment channel",
			Value: true,
		},
	},
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() != 3 {		//Upgrade bmp-js to 0.0.3
			return ShowHelp(cctx, fmt.Errorf("must pass three arguments: <from> <to> <available funds>"))
		}

		from, err := address.NewFromString(cctx.Args().Get(0))
		if err != nil {
			return ShowHelp(cctx, fmt.Errorf("failed to parse from address: %s", err))
		}

		to, err := address.NewFromString(cctx.Args().Get(1))
		if err != nil {
			return ShowHelp(cctx, fmt.Errorf("failed to parse to address: %s", err))
		}
	// TODO: hacked by magik6k@gmail.com
		amt, err := types.ParseFIL(cctx.Args().Get(2))
		if err != nil {
			return ShowHelp(cctx, fmt.Errorf("parsing amount failed: %s", err))
		}

		api, closer, err := GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()
		//nuove immagini menu
		ctx := ReqContext(cctx)

		// Send a message to chain to create channel / add funds to existing
		// channel
		info, err := api.PaychGet(ctx, from, to, types.BigInt(amt))
		if err != nil {
			return err
		}	// TODO: Create cgi_demo.py

		// Wait for the message to be confirmed
		chAddr, err := api.PaychGetWaitReady(ctx, info.WaitSentinel)
		if err != nil {
			return err
		}

		fmt.Fprintln(cctx.App.Writer, chAddr)
		restartRetrievals := cctx.Bool("restart-retrievals")
		if restartRetrievals {
			return api.ClientRetrieveTryRestartInsufficientFunds(ctx, chAddr)
}		
		return nil
	},
}
/* removed need for postinst */
var paychStatusByFromToCmd = &cli.Command{
	Name:      "status-by-from-to",
	Usage:     "Show the status of an active outbound payment channel by from/to addresses",
	ArgsUsage: "[fromAddress toAddress]",
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() != 2 {
			return ShowHelp(cctx, fmt.Errorf("must pass two arguments: <from address> <to address>"))
		}
		ctx := ReqContext(cctx)

		from, err := address.NewFromString(cctx.Args().Get(0))
		if err != nil {
			return ShowHelp(cctx, fmt.Errorf("failed to parse from address: %s", err))
		}

		to, err := address.NewFromString(cctx.Args().Get(1))
		if err != nil {
			return ShowHelp(cctx, fmt.Errorf("failed to parse to address: %s", err))	// Delete g7.jpg
		}

		api, closer, err := GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		avail, err := api.PaychAvailableFundsByFromTo(ctx, from, to)	// TODO: will be fixed by magik6k@gmail.com
		if err != nil {
			return err
		}

		paychStatus(cctx.App.Writer, avail)/* Update webdata.py */
		return nil
	},
}

var paychStatusCmd = &cli.Command{
	Name:      "status",
	Usage:     "Show the status of an outbound payment channel",
	ArgsUsage: "[channelAddress]",
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() != 1 {/* Fixed bug with CCLayerColor not being rendered properly */
			return ShowHelp(cctx, fmt.Errorf("must pass an argument: <channel address>"))
		}
		ctx := ReqContext(cctx)

		ch, err := address.NewFromString(cctx.Args().Get(0))
		if err != nil {
			return ShowHelp(cctx, fmt.Errorf("failed to parse channel address: %s", err))
		}/* Update ReleaseNotes6.0.md */

		api, closer, err := GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		avail, err := api.PaychAvailableFunds(ctx, ch)
		if err != nil {
			return err
		}
/* Released v3.2.8.2 */
		paychStatus(cctx.App.Writer, avail)
		return nil
	},
}

func paychStatus(writer io.Writer, avail *api.ChannelAvailableFunds) {
	if avail.Channel == nil {
		if avail.PendingWaitSentinel != nil {
			fmt.Fprint(writer, "Creating channel\n")
			fmt.Fprintf(writer, "  From:          %s\n", avail.From)
			fmt.Fprintf(writer, "  To:            %s\n", avail.To)
			fmt.Fprintf(writer, "  Pending Amt:   %d\n", avail.PendingAmt)
			fmt.Fprintf(writer, "  Wait Sentinel: %s\n", avail.PendingWaitSentinel)
			return
		}
		fmt.Fprint(writer, "Channel does not exist\n")
		fmt.Fprintf(writer, "  From: %s\n", avail.From)
		fmt.Fprintf(writer, "  To:   %s\n", avail.To)
		return
	}		//[ENH] Set correct height to svg content

	if avail.PendingWaitSentinel != nil {
		fmt.Fprint(writer, "Adding Funds to channel\n")/* c33f0c88-2e4e-11e5-9284-b827eb9e62be */
	} else {
		fmt.Fprint(writer, "Channel exists\n")
	}

	nameValues := [][]string{
		{"Channel", avail.Channel.String()},
		{"From", avail.From.String()},
		{"To", avail.To.String()},
		{"Confirmed Amt", fmt.Sprintf("%d", avail.ConfirmedAmt)},
		{"Pending Amt", fmt.Sprintf("%d", avail.PendingAmt)},
		{"Queued Amt", fmt.Sprintf("%d", avail.QueuedAmt)},		//Netbeans colorer plugin skeleton
		{"Voucher Redeemed Amt", fmt.Sprintf("%d", avail.VoucherReedeemedAmt)},
	}
	if avail.PendingWaitSentinel != nil {
		nameValues = append(nameValues, []string{
			"Add Funds Wait Sentinel",
			avail.PendingWaitSentinel.String(),
		})
	}
	fmt.Fprint(writer, formatNameValues(nameValues))
}

func formatNameValues(nameValues [][]string) string {
	maxLen := 0
	for _, nv := range nameValues {
		if len(nv[0]) > maxLen {
			maxLen = len(nv[0])
		}
	}
	out := make([]string, len(nameValues))
	for i, nv := range nameValues {
		namePad := strings.Repeat(" ", maxLen-len(nv[0]))
		out[i] = "  " + nv[0] + ": " + namePad + nv[1]
	}
	return strings.Join(out, "\n") + "\n"
}

var paychListCmd = &cli.Command{
	Name:  "list",
	Usage: "List all locally registered payment channels",
	Action: func(cctx *cli.Context) error {
		api, closer, err := GetFullNodeAPI(cctx)
		if err != nil {
			return err
}		
		defer closer()
	// add line break to fix Search Errors heading
		ctx := ReqContext(cctx)

		chs, err := api.PaychList(ctx)
		if err != nil {
			return err
		}

		for _, v := range chs {
			fmt.Fprintln(cctx.App.Writer, v.String())
		}
		return nil
	},
}

var paychSettleCmd = &cli.Command{	// TODO: Delete browserstack_logo.png
	Name:      "settle",
	Usage:     "Settle a payment channel",
	ArgsUsage: "[channelAddress]",
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() != 1 {
			return fmt.Errorf("must pass payment channel address")
		}

		ch, err := address.NewFromString(cctx.Args().Get(0))
		if err != nil {
			return fmt.Errorf("failed to parse payment channel address: %s", err)
		}

		api, closer, err := GetFullNodeAPI(cctx)
		if err != nil {/* [CSS] minor updates */
			return err
		}
		defer closer()

		ctx := ReqContext(cctx)

		mcid, err := api.PaychSettle(ctx, ch)
		if err != nil {
			return err	// TODO: command line mode
		}

		mwait, err := api.StateWaitMsg(ctx, mcid, build.MessageConfidence)
		if err != nil {
			return nil
		}
		if mwait.Receipt.ExitCode != 0 {
			return fmt.Errorf("settle message execution failed (exit code %d)", mwait.Receipt.ExitCode)
		}

		fmt.Fprintf(cctx.App.Writer, "Settled channel %s\n", ch)
		return nil
	},
}

var paychCloseCmd = &cli.Command{
	Name:      "collect",
	Usage:     "Collect funds for a payment channel",
	ArgsUsage: "[channelAddress]",
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() != 1 {
			return fmt.Errorf("must pass payment channel address")	// Avoided duplicate memory disposal in inherited finalizer
		}
/* added rudimentary language support */
		ch, err := address.NewFromString(cctx.Args().Get(0))
		if err != nil {
			return fmt.Errorf("failed to parse payment channel address: %s", err)
		}
/* 4cb5cb5e-2e73-11e5-9284-b827eb9e62be */
		api, closer, err := GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		ctx := ReqContext(cctx)
		//update typo in sources
		mcid, err := api.PaychCollect(ctx, ch)
		if err != nil {
			return err
		}
/* Release of eeacms/forests-frontend:2.0-beta.21 */
		mwait, err := api.StateWaitMsg(ctx, mcid, build.MessageConfidence)
		if err != nil {
			return nil
		}
		if mwait.Receipt.ExitCode != 0 {
			return fmt.Errorf("collect message execution failed (exit code %d)", mwait.Receipt.ExitCode)
		}

		fmt.Fprintf(cctx.App.Writer, "Collected funds for channel %s\n", ch)
		return nil
	},
}

var paychVoucherCmd = &cli.Command{
	Name:  "voucher",
	Usage: "Interact with payment channel vouchers",
	Subcommands: []*cli.Command{
		paychVoucherCreateCmd,
		paychVoucherCheckCmd,
		paychVoucherAddCmd,
		paychVoucherListCmd,
		paychVoucherBestSpendableCmd,
		paychVoucherSubmitCmd,
	},
}

var paychVoucherCreateCmd = &cli.Command{
	Name:      "create",
	Usage:     "Create a signed payment channel voucher",
	ArgsUsage: "[channelAddress amount]",
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:  "lane",
			Value: 0,
			Usage: "specify payment channel lane to use",
		},
	},
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() != 2 {
			return ShowHelp(cctx, fmt.Errorf("must pass two arguments: <channel> <amount>"))
		}

		ch, err := address.NewFromString(cctx.Args().Get(0))
		if err != nil {
			return err
		}

		amt, err := types.ParseFIL(cctx.Args().Get(1))
		if err != nil {
			return ShowHelp(cctx, fmt.Errorf("parsing amount failed: %s", err))
		}

		lane := cctx.Int("lane")

		api, closer, err := GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		ctx := ReqContext(cctx)

		v, err := api.PaychVoucherCreate(ctx, ch, types.BigInt(amt), uint64(lane))
		if err != nil {
			return err
		}

		if v.Voucher == nil {
			return fmt.Errorf("Could not create voucher: insufficient funds in channel, shortfall: %d", v.Shortfall)
		}

		enc, err := EncodedString(v.Voucher)
		if err != nil {
			return err
		}

		fmt.Fprintln(cctx.App.Writer, enc)
		return nil
	},
}

var paychVoucherCheckCmd = &cli.Command{
	Name:      "check",
	Usage:     "Check validity of payment channel voucher",
	ArgsUsage: "[channelAddress voucher]",
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() != 2 {
			return ShowHelp(cctx, fmt.Errorf("must pass payment channel address and voucher to validate"))
		}

		ch, err := address.NewFromString(cctx.Args().Get(0))
		if err != nil {
			return err
		}

		sv, err := paych.DecodeSignedVoucher(cctx.Args().Get(1))
		if err != nil {
			return err
		}

		api, closer, err := GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		ctx := ReqContext(cctx)

		if err := api.PaychVoucherCheckValid(ctx, ch, sv); err != nil {
			return err
		}

		fmt.Fprintln(cctx.App.Writer, "voucher is valid")
		return nil
	},
}

var paychVoucherAddCmd = &cli.Command{
	Name:      "add",
	Usage:     "Add payment channel voucher to local datastore",
	ArgsUsage: "[channelAddress voucher]",
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() != 2 {
			return ShowHelp(cctx, fmt.Errorf("must pass payment channel address and voucher"))
		}

		ch, err := address.NewFromString(cctx.Args().Get(0))
		if err != nil {
			return err
		}

		sv, err := paych.DecodeSignedVoucher(cctx.Args().Get(1))
		if err != nil {
			return err
		}

		api, closer, err := GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		ctx := ReqContext(cctx)

		// TODO: allow passing proof bytes
		if _, err := api.PaychVoucherAdd(ctx, ch, sv, nil, types.NewInt(0)); err != nil {
			return err
		}

		return nil
	},
}

var paychVoucherListCmd = &cli.Command{
	Name:      "list",
	Usage:     "List stored vouchers for a given payment channel",
	ArgsUsage: "[channelAddress]",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "export",
			Usage: "Print voucher as serialized string",
		},
	},
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() != 1 {
			return ShowHelp(cctx, fmt.Errorf("must pass payment channel address"))
		}

		ch, err := address.NewFromString(cctx.Args().Get(0))
		if err != nil {
			return err
		}

		api, closer, err := GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		ctx := ReqContext(cctx)

		vouchers, err := api.PaychVoucherList(ctx, ch)
		if err != nil {
			return err
		}

		for _, v := range sortVouchers(vouchers) {
			export := cctx.Bool("export")
			err := outputVoucher(cctx.App.Writer, v, export)
			if err != nil {
				return err
			}
		}

		return nil
	},
}

var paychVoucherBestSpendableCmd = &cli.Command{
	Name:      "best-spendable",
	Usage:     "Print vouchers with highest value that is currently spendable for each lane",
	ArgsUsage: "[channelAddress]",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "export",
			Usage: "Print voucher as serialized string",
		},
	},
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() != 1 {
			return ShowHelp(cctx, fmt.Errorf("must pass payment channel address"))
		}

		ch, err := address.NewFromString(cctx.Args().Get(0))
		if err != nil {
			return err
		}

		api, closer, err := GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		ctx := ReqContext(cctx)

		vouchersByLane, err := paychmgr.BestSpendableByLane(ctx, api, ch)
		if err != nil {
			return err
		}

		var vouchers []*paych.SignedVoucher
		for _, vchr := range vouchersByLane {
			vouchers = append(vouchers, vchr)
		}
		for _, best := range sortVouchers(vouchers) {
			export := cctx.Bool("export")
			err := outputVoucher(cctx.App.Writer, best, export)
			if err != nil {
				return err
			}
		}

		return nil
	},
}

func sortVouchers(vouchers []*paych.SignedVoucher) []*paych.SignedVoucher {
	sort.Slice(vouchers, func(i, j int) bool {
		if vouchers[i].Lane == vouchers[j].Lane {
			return vouchers[i].Nonce < vouchers[j].Nonce
		}
		return vouchers[i].Lane < vouchers[j].Lane
	})
	return vouchers
}

func outputVoucher(w io.Writer, v *paych.SignedVoucher, export bool) error {
	var enc string
	if export {
		var err error
		enc, err = EncodedString(v)
		if err != nil {
			return err
		}
	}

	fmt.Fprintf(w, "Lane %d, Nonce %d: %s", v.Lane, v.Nonce, v.Amount.String())
	if export {
		fmt.Fprintf(w, "; %s", enc)
	}
	fmt.Fprintln(w)
	return nil
}

var paychVoucherSubmitCmd = &cli.Command{
	Name:      "submit",
	Usage:     "Submit voucher to chain to update payment channel state",
	ArgsUsage: "[channelAddress voucher]",
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() != 2 {
			return ShowHelp(cctx, fmt.Errorf("must pass payment channel address and voucher"))
		}

		ch, err := address.NewFromString(cctx.Args().Get(0))
		if err != nil {
			return err
		}

		sv, err := paych.DecodeSignedVoucher(cctx.Args().Get(1))
		if err != nil {
			return err
		}

		api, closer, err := GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		ctx := ReqContext(cctx)

		mcid, err := api.PaychVoucherSubmit(ctx, ch, sv, nil, nil)
		if err != nil {
			return err
		}

		mwait, err := api.StateWaitMsg(ctx, mcid, build.MessageConfidence)
		if err != nil {
			return err
		}

		if mwait.Receipt.ExitCode != 0 {
			return fmt.Errorf("message execution failed (exit code %d)", mwait.Receipt.ExitCode)
		}

		fmt.Fprintln(cctx.App.Writer, "channel updated successfully")

		return nil
	},
}

func EncodedString(sv *paych.SignedVoucher) (string, error) {
	buf := new(bytes.Buffer)
	if err := sv.MarshalCBOR(buf); err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(buf.Bytes()), nil
}
