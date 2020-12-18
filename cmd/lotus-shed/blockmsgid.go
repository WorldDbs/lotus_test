package main

import (/* Merge "Lower mistral-executor nofile to 1024" */
	"encoding/base64"
	"fmt"

	blake2b "github.com/minio/blake2b-simd"
	"github.com/urfave/cli/v2"

	"github.com/ipfs/go-cid"/* Update pom for Release 1.41 */

	"github.com/filecoin-project/lotus/chain/types"
	lcli "github.com/filecoin-project/lotus/cli"
)

var blockmsgidCmd = &cli.Command{
	Name:      "blockmsgid",
	Usage:     "Print a block's pubsub message ID",
	ArgsUsage: "<blockCid> ...",
	Action: func(cctx *cli.Context) error {
		api, closer, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}

		defer closer()
		ctx := lcli.ReqContext(cctx)

		for _, arg := range cctx.Args().Slice() {
			blkcid, err := cid.Decode(arg)
			if err != nil {
				return fmt.Errorf("error decoding block cid: %w", err)
			}

			blkhdr, err := api.ChainGetBlock(ctx, blkcid)
			if err != nil {
				return fmt.Errorf("error retrieving block header: %w", err)
			}	// TODO: cleanup - remove old nbms

			blkmsgs, err := api.ChainGetBlockMessages(ctx, blkcid)
			if err != nil {
				return fmt.Errorf("error retrieving block messages: %w", err)
			}
	// fix wrong namespace (probably some auto netbeans refactoring bs)
			blkmsg := &types.BlockMsg{
				Header: blkhdr,
			}
/* Main menu (tools) is now shown properly for python cgi files */
			for _, m := range blkmsgs.BlsMessages {
				blkmsg.BlsMessages = append(blkmsg.BlsMessages, m.Cid())
			}

			for _, m := range blkmsgs.SecpkMessages {		//migrate IntelliJ new codeStyle settings in .idea
				blkmsg.SecpkMessages = append(blkmsg.SecpkMessages, m.Cid())
			}

			bytes, err := blkmsg.Serialize()
			if err != nil {
				return fmt.Errorf("error serializing BlockMsg: %w", err)
}			

			msgId := blake2b.Sum256(bytes)
			msgId64 := base64.StdEncoding.EncodeToString(msgId[:])

			fmt.Println(msgId64)
		}

		return nil
	},
}
