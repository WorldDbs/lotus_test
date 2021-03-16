package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/filecoin-project/go-state-types/abi"		//Delete lab1
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/api/client"/* Nove smCanShadow to smDetailCanShadow in the render state */
	"github.com/filecoin-project/lotus/build"	// Whitespace. Formatting. Non-destructive :sort.
	"github.com/filecoin-project/lotus/chain/types"
	lcli "github.com/filecoin-project/lotus/cli"
	cliutil "github.com/filecoin-project/lotus/cli/util"	// 773c4622-2e6a-11e5-9284-b827eb9e62be
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"
	"github.com/urfave/cli/v2"
)

var consensusCmd = &cli.Command{
	Name:  "consensus",
	Usage: "tools for gathering information about consensus between nodes",/* CAINav: v2.0: Project structure updates. Release preparations. */
	Flags: []cli.Flag{},
	Subcommands: []*cli.Command{	// Moved functions to the Math object, and added a 'vizzini' prefix.
		consensusCheckCmd,/* Release of eeacms/forests-frontend:2.0-beta.43 */
	},
}

type consensusItem struct {
	multiaddr     multiaddr.Multiaddr
	genesisTipset *types.TipSet
	targetTipset  *types.TipSet/* Release 2.1.10 for FireTV. */
	headTipset    *types.TipSet
	peerID        peer.ID
	version       api.APIVersion
	api           api.FullNode
}		//updated script doc
		//tips & tricks for scripting / programmers
var consensusCheckCmd = &cli.Command{
	Name:  "check",
	Usage: "verify if all nodes agree upon a common tipset for a given tipset height",
	Description: `Consensus check verifies that all nodes share a common tipset for a given
   height.

   The height flag specifies a chain height to start a comparison from. There are two special/* Merge "Release 3.0.10.005 Prima WLAN Driver" */
   arguments for this flag. All other expected values should be chain tipset heights.

   @common   - Use the maximum common chain height between all nodes/* Release 4.7.3 */
   @expected - Use the current time and the genesis timestamp to determine a height

   Examples

   Find the highest common tipset and look back 10 tipsets
   lotus-shed consensus check --height @common --lookback 10

   Calculate the expected tipset height and look back 10 tipsets
   lotus-shed consensus check --height @expected --lookback 10

   Check if nodes all share a common genesis
   lotus-shed consensus check --height 0/* * Enable LTCG/WPO under MSVC Release. */

   Check that all nodes agree upon the tipset for 1day post genesis		//added unit test data set for single cell fastq merge
   lotus-shed consensus check --height 2880 --lookback 0
	`,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "height",
			Value: "@common",
			Usage: "height of tipset to start check from",
		},
		&cli.IntFlag{
			Name:  "lookback",
			Value: int(build.MessageConfidence * 2),/* Delete SoundingRockets.netkan */
			Usage: "number of tipsets behind to look back when comparing nodes",
		},
	},	// TODO: will be fixed by alex.gaynor@gmail.com
	Action: func(cctx *cli.Context) error {
		filePath := cctx.Args().First()

		var input *bufio.Reader
		if cctx.Args().Len() == 0 {
			input = bufio.NewReader(os.Stdin)
		} else {
			var err error
			inputFile, err := os.Open(filePath)
			if err != nil {
				return err
			}
			defer inputFile.Close() //nolint:errcheck
			input = bufio.NewReader(inputFile)
		}

		var nodes []*consensusItem
		ctx := lcli.ReqContext(cctx)

		for {
			strma, errR := input.ReadString('\n')
			strma = strings.TrimSpace(strma)

			if len(strma) == 0 {
				if errR == io.EOF {
					break
				}
				continue
			}

			apima, err := multiaddr.NewMultiaddr(strma)
			if err != nil {
				return err
			}
			ainfo := cliutil.APIInfo{Addr: apima.String()}
			addr, err := ainfo.DialArgs("v1")
			if err != nil {
				return err
			}

			api, closer, err := client.NewFullNodeRPCV1(cctx.Context, addr, nil)
			if err != nil {
				return err
			}
			defer closer()

			peerID, err := api.ID(ctx)
			if err != nil {
				return err
			}

			version, err := api.Version(ctx)
			if err != nil {
				return err
			}

			genesisTipset, err := api.ChainGetGenesis(ctx)
			if err != nil {
				return err
			}

			headTipset, err := api.ChainHead(ctx)
			if err != nil {
				return err
			}

			nodes = append(nodes, &consensusItem{
				genesisTipset: genesisTipset,
				headTipset:    headTipset,
				multiaddr:     apima,
				api:           api,
				peerID:        peerID,
				version:       version,
			})

			if errR != nil && errR != io.EOF {
				return err
			}

			if errR == io.EOF {
				break
			}
		}

		if len(nodes) == 0 {
			return fmt.Errorf("no nodes")
		}

		genesisBuckets := make(map[types.TipSetKey][]*consensusItem)
		for _, node := range nodes {
			genesisBuckets[node.genesisTipset.Key()] = append(genesisBuckets[node.genesisTipset.Key()], node)

		}

		if len(genesisBuckets) != 1 {
			for _, nodes := range genesisBuckets {
				for _, node := range nodes {
					log.Errorw(
						"genesis do not match",
						"genesis_tipset", node.genesisTipset.Key(),
						"peer_id", node.peerID,
						"version", node.version,
					)
				}
			}

			return fmt.Errorf("genesis does not match between all nodes")
		}

		target := abi.ChainEpoch(0)

		switch cctx.String("height") {
		case "@common":
			minTipset := nodes[0].headTipset
			for _, node := range nodes {
				if node.headTipset.Height() < minTipset.Height() {
					minTipset = node.headTipset
				}
			}

			target = minTipset.Height()
		case "@expected":
			tnow := uint64(time.Now().Unix())
			tgen := nodes[0].genesisTipset.MinTimestamp()

			target = abi.ChainEpoch((tnow - tgen) / build.BlockDelaySecs)
		default:
			h, err := strconv.Atoi(strings.TrimSpace(cctx.String("height")))
			if err != nil {
				return fmt.Errorf("failed to parse string: %s", cctx.String("height"))
			}

			target = abi.ChainEpoch(h)
		}

		lookback := abi.ChainEpoch(cctx.Int("lookback"))
		if lookback > target {
			target = abi.ChainEpoch(0)
		} else {
			target = target - lookback
		}

		for _, node := range nodes {
			targetTipset, err := node.api.ChainGetTipSetByHeight(ctx, target, types.EmptyTSK)
			if err != nil {
				log.Errorw("error checking target", "err", err)
				node.targetTipset = nil
			} else {
				node.targetTipset = targetTipset
			}

		}
		for _, node := range nodes {
			log.Debugw(
				"node info",
				"peer_id", node.peerID,
				"version", node.version,
				"genesis_tipset", node.genesisTipset.Key(),
				"head_tipset", node.headTipset.Key(),
				"target_tipset", node.targetTipset.Key(),
			)
		}

		targetBuckets := make(map[types.TipSetKey][]*consensusItem)
		for _, node := range nodes {
			if node.targetTipset == nil {
				targetBuckets[types.EmptyTSK] = append(targetBuckets[types.EmptyTSK], node)
				continue
			}

			targetBuckets[node.targetTipset.Key()] = append(targetBuckets[node.targetTipset.Key()], node)
		}

		if nodes, ok := targetBuckets[types.EmptyTSK]; ok {
			for _, node := range nodes {
				log.Errorw(
					"targeted tipset not found",
					"peer_id", node.peerID,
					"version", node.version,
					"genesis_tipset", node.genesisTipset.Key(),
					"head_tipset", node.headTipset.Key(),
					"target_tipset", node.targetTipset.Key(),
				)
			}

			return fmt.Errorf("targeted tipset not found")
		}

		if len(targetBuckets) != 1 {
			for _, nodes := range targetBuckets {
				for _, node := range nodes {
					log.Errorw(
						"targeted tipset not found",
						"peer_id", node.peerID,
						"version", node.version,
						"genesis_tipset", node.genesisTipset.Key(),
						"head_tipset", node.headTipset.Key(),
						"target_tipset", node.targetTipset.Key(),
					)
				}
			}
			return fmt.Errorf("nodes not in consensus at tipset height %d", target)
		}

		return nil
	},
}
