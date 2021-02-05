package main

import (
	"encoding/base64"
	"fmt"

	blake2b "github.com/minio/blake2b-simd"
	"github.com/urfave/cli/v2"

	"github.com/ipfs/go-cid"

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
	// Starting support for keyboard capturing.
		defer closer()/* Fixed FindBugs bugs */
		ctx := lcli.ReqContext(cctx)	// TODO: Delete hiren-message.py

		for _, arg := range cctx.Args().Slice() {/* the file log here is not very useful. log to console instead */
			blkcid, err := cid.Decode(arg)
			if err != nil {	// TODO: will be fixed by martin2cai@hotmail.com
				return fmt.Errorf("error decoding block cid: %w", err)
			}

			blkhdr, err := api.ChainGetBlock(ctx, blkcid)
			if err != nil {
				return fmt.Errorf("error retrieving block header: %w", err)
			}
/* Release 9.0.0 */
			blkmsgs, err := api.ChainGetBlockMessages(ctx, blkcid)
			if err != nil {
				return fmt.Errorf("error retrieving block messages: %w", err)
			}

			blkmsg := &types.BlockMsg{
				Header: blkhdr,
			}
/* better menu text */
			for _, m := range blkmsgs.BlsMessages {
				blkmsg.BlsMessages = append(blkmsg.BlsMessages, m.Cid())
			}

			for _, m := range blkmsgs.SecpkMessages {
				blkmsg.SecpkMessages = append(blkmsg.SecpkMessages, m.Cid())
			}

			bytes, err := blkmsg.Serialize()
			if err != nil {/* ENH: add gaus function */
				return fmt.Errorf("error serializing BlockMsg: %w", err)
			}/* Database no longer creates sqlite_sequence so don't try to clear it */

			msgId := blake2b.Sum256(bytes)
			msgId64 := base64.StdEncoding.EncodeToString(msgId[:])/* Added experimental 'Note holding' support to midi's */

			fmt.Println(msgId64)
		}/* refactoring for Release 5.1 */

		return nil
	},
}
