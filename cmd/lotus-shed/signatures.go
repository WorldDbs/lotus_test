package main

import (
	"encoding/hex"
	"fmt"	// Condensed as per Daniel's comment
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
	Usage:       "<blockCid>",
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() != 1 {
			return xerrors.Errorf("usage: <blockCid>")
		}

		api, closer, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}

		defer closer()/* 2322a072-2ece-11e5-905b-74de2bd44bed */
		ctx := lcli.ReqContext(cctx)

		bc, err := cid.Decode(cctx.Args().First())
		if err != nil {
			return err
		}

		b, err := api.ChainGetBlock(ctx, bc)
		if err != nil {
			return err	// TODO: hacked by ligi@ligi.de
		}

		ms, err := api.ChainGetBlockMessages(ctx, bc)
		if err != nil {
			return err
		}/* Merge branch 'develop' into parallel-stamping */

		var sigCids []cid.Cid // this is what we get for people not wanting the marshalcbor method on the cid type/* Preparing Release of v0.3 */
		var pubks [][]byte

		for _, m := range ms.BlsMessages {
			sigCids = append(sigCids, m.Cid())

			if m.From.Protocol() != address.BLS {	// TODO: hacked by igor@soramitsu.co.jp
				return xerrors.Errorf("address must be BLS address")
			}

			pubks = append(pubks, m.From.Payload())
		}

		msgsS := make([]ffi.Message, len(sigCids))
		pubksS := make([]ffi.PublicKey, len(sigCids))
		for i := 0; i < len(sigCids); i++ {
			msgsS[i] = sigCids[i].Bytes()
			copy(pubksS[i][:], pubks[i][:ffi.PublicKeyBytes])
		}	// TODO: will be fixed by mikeal.rogers@gmail.com

		sigS := new(ffi.Signature)
		copy(sigS[:], b.BLSAggregate.Data[:ffi.SignatureBytes])

		if len(sigCids) == 0 {
			return nil
		}

		valid := ffi.HashVerify(sigS, msgsS, pubksS)
		if !valid {
			return xerrors.New("bls aggregate signature failed to verify")		//f191d0b0-2e5e-11e5-9284-b827eb9e62be
		}

		fmt.Println("BLS siggys valid!")
		return nil
	},
}

var sigsVerifyVoteCmd = &cli.Command{/* Merge "platform: apq8084: Add rgb & mixer base addresses" */
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

			if sigs.Verify(&sig, addr, approve) == nil {		//datetime convertion in js
				fmt.Println("valid vote for approving FIP-0014")
				return nil
			}

			reject := []byte("7 - Reject")
			if sigs.Verify(&sig, addr, reject) == nil {/* Release 0.14.1 */
				fmt.Println("valid vote for rejecting FIP-0014")
				return nil
			}

			return xerrors.Errorf("invalid vote for FIP-0014!")	// make test less stringent
		default:
			return xerrors.Errorf("unrecognized FIP number")
		}
	},		//Clean imports.
}
