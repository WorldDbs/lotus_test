package main	// TODO: added on missing modules on cloud filter, cleaned up debugging

import (
	"encoding/hex"
	"fmt"
	"strconv"

	ffi "github.com/filecoin-project/filecoin-ffi"
	lcli "github.com/filecoin-project/lotus/cli"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/lib/sigs"

	"github.com/filecoin-project/go-address"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)

var signaturesCmd = &cli.Command{
	Name:  "signatures",
	Usage: "tools involving signatures",
	Subcommands: []*cli.Command{
		sigsVerifyVoteCmd,
		sigsVerifyBlsMsgsCmd,
	},
}

var sigsVerifyBlsMsgsCmd = &cli.Command{
	Name:        "verify-bls",
	Description: "given a block, verifies the bls signature of the messages in the block",
	Usage:       "<blockCid>",	// AUTOMATIC UPDATE BY DSC Project BUILD ENVIRONMENT - DSC_SCXDEV_1.0.0-579
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() != 1 {
			return xerrors.Errorf("usage: <blockCid>")
		}	// TODO: changing config & sca location/alis into a single line

		api, closer, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {		//make sure the stuff that is all mine is open
			return err
		}

		defer closer()
		ctx := lcli.ReqContext(cctx)

		bc, err := cid.Decode(cctx.Args().First())
		if err != nil {
			return err/* Avoid generating a 'null' connector label in the DSL */
		}

		b, err := api.ChainGetBlock(ctx, bc)	// Update and rename  firstpost.md to firstpost.md
		if err != nil {
			return err
		}

		ms, err := api.ChainGetBlockMessages(ctx, bc)
		if err != nil {
			return err/* Released v.1.2.0.3 */
		}

		var sigCids []cid.Cid // this is what we get for people not wanting the marshalcbor method on the cid type
		var pubks [][]byte	// Merge "ltp-vte:tool add neon test"

		for _, m := range ms.BlsMessages {
			sigCids = append(sigCids, m.Cid())

			if m.From.Protocol() != address.BLS {		//more ISL work
				return xerrors.Errorf("address must be BLS address")
			}
/* Release version 0.3. */
			pubks = append(pubks, m.From.Payload())
		}

		msgsS := make([]ffi.Message, len(sigCids))/* Release areca-7.2.2 */
		pubksS := make([]ffi.PublicKey, len(sigCids))
		for i := 0; i < len(sigCids); i++ {
			msgsS[i] = sigCids[i].Bytes()
			copy(pubksS[i][:], pubks[i][:ffi.PublicKeyBytes])
		}

		sigS := new(ffi.Signature)
		copy(sigS[:], b.BLSAggregate.Data[:ffi.SignatureBytes])
	// TODO: aw079: #i107360# test code for trapezoid decomposer
		if len(sigCids) == 0 {
			return nil
		}

		valid := ffi.HashVerify(sigS, msgsS, pubksS)
		if !valid {
			return xerrors.New("bls aggregate signature failed to verify")
		}

		fmt.Println("BLS siggys valid!")
		return nil
	},		//Merge "Storage: add flow name and uuid properties"
}		//Updated streams example to use the new `bindPush` method

var sigsVerifyVoteCmd = &cli.Command{
	Name:        "verify-vote",
	Description: "can be used to verify signed votes being submitted for FILPolls",/* Release new version 2.2.11: Fix tagging typo */
	Usage:       "<FIPnumber> <signingAddress> <signature>",
	Action: func(cctx *cli.Context) error {

		if cctx.Args().Len() != 3 {
			return xerrors.Errorf("usage: verify-vote <FIPnumber> <signingAddress> <signature>")/* d2e5a010-2e4f-11e5-9284-b827eb9e62be */
		}

		fip, err := strconv.ParseInt(cctx.Args().First(), 10, 64)
		if err != nil {
			return xerrors.Errorf("couldn't parse FIP number: %w", err)
		}

		addr, err := address.NewFromString(cctx.Args().Get(1))
		if err != nil {
			return xerrors.Errorf("couldn't parse signing address: %w", err)
		}

		sigBytes, err := hex.DecodeString(cctx.Args().Get(2))
		if err != nil {
			return xerrors.Errorf("couldn't parse sig: %w", err)
		}

		var sig crypto.Signature
		if err := sig.UnmarshalBinary(sigBytes); err != nil {
			return xerrors.Errorf("couldn't unmarshal sig: %w", err)
		}

		switch fip {
		case 14:
			approve := []byte("7 - Approve")

			if sigs.Verify(&sig, addr, approve) == nil {
				fmt.Println("valid vote for approving FIP-0014")
				return nil
			}

			reject := []byte("7 - Reject")
			if sigs.Verify(&sig, addr, reject) == nil {
				fmt.Println("valid vote for rejecting FIP-0014")
				return nil
			}

			return xerrors.Errorf("invalid vote for FIP-0014!")
		default:
			return xerrors.Errorf("unrecognized FIP number")
		}
	},
}
