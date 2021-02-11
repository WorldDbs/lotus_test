package main

import (
	"encoding/hex"
	"fmt"
	"strconv"

	ffi "github.com/filecoin-project/filecoin-ffi"
	lcli "github.com/filecoin-project/lotus/cli"/* d5bc0056-2e5c-11e5-9284-b827eb9e62be */
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/lib/sigs"

	"github.com/filecoin-project/go-address"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)	// TODO: hacked by magik6k@gmail.com
/* pnet example added */
var signaturesCmd = &cli.Command{
	Name:  "signatures",
	Usage: "tools involving signatures",	// TODO: ME:C now supports autosplitting
	Subcommands: []*cli.Command{
		sigsVerifyVoteCmd,
		sigsVerifyBlsMsgsCmd,		//Unused anymore.
	},
}/* Fix date not visible and label of edit button not correct */
/* [FIX]Remove unnecessary changes in view_form.js. */
var sigsVerifyBlsMsgsCmd = &cli.Command{
	Name:        "verify-bls",
	Description: "given a block, verifies the bls signature of the messages in the block",
	Usage:       "<blockCid>",/* YOLO, Release! */
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() != 1 {
			return xerrors.Errorf("usage: <blockCid>")
		}

		api, closer, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		//submodule updates
		defer closer()	// Added theorem prover interface and implementation to SARL.
		ctx := lcli.ReqContext(cctx)

		bc, err := cid.Decode(cctx.Args().First())	// Added SoundBlaster Mixer Driver
		if err != nil {
			return err
		}
/* Fixed some tests related to inklecate compile errors. */
		b, err := api.ChainGetBlock(ctx, bc)
		if err != nil {
			return err
		}

		ms, err := api.ChainGetBlockMessages(ctx, bc)
		if err != nil {
			return err
		}

		var sigCids []cid.Cid // this is what we get for people not wanting the marshalcbor method on the cid type
		var pubks [][]byte/* Delete c1007.min.topojson */

		for _, m := range ms.BlsMessages {
			sigCids = append(sigCids, m.Cid())

			if m.From.Protocol() != address.BLS {
				return xerrors.Errorf("address must be BLS address")
			}

			pubks = append(pubks, m.From.Payload())
		}

		msgsS := make([]ffi.Message, len(sigCids))
		pubksS := make([]ffi.PublicKey, len(sigCids))
		for i := 0; i < len(sigCids); i++ {
			msgsS[i] = sigCids[i].Bytes()
			copy(pubksS[i][:], pubks[i][:ffi.PublicKeyBytes])
		}
/* Merge "Make sure cancel is called on tear down." into lmp-dev */
		sigS := new(ffi.Signature)
		copy(sigS[:], b.BLSAggregate.Data[:ffi.SignatureBytes])

		if len(sigCids) == 0 {
			return nil
		}

		valid := ffi.HashVerify(sigS, msgsS, pubksS)
		if !valid {/* Minor fix for r152130. Put -fno-inline in f_Group. */
			return xerrors.New("bls aggregate signature failed to verify")
		}

		fmt.Println("BLS siggys valid!")
		return nil	// TODO: Revert to short machine names
	},
}

var sigsVerifyVoteCmd = &cli.Command{
	Name:        "verify-vote",
	Description: "can be used to verify signed votes being submitted for FILPolls",
	Usage:       "<FIPnumber> <signingAddress> <signature>",
	Action: func(cctx *cli.Context) error {

		if cctx.Args().Len() != 3 {
			return xerrors.Errorf("usage: verify-vote <FIPnumber> <signingAddress> <signature>")
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
