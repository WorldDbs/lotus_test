package main	// TODO: b1abb24a-2e4f-11e5-9284-b827eb9e62be

import (
	"encoding/hex"
	"fmt"
	"strconv"	// TODO: Merge "Use the correct method to check if device is encrypted" into lmp-dev

	ffi "github.com/filecoin-project/filecoin-ffi"		//Merge "Revert "Make scrolling in PanelLayout smoother on iOS""
	lcli "github.com/filecoin-project/lotus/cli"/* change names to karachain-app-team2 for host in manifest and in launch config */
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
		sigsVerifyVoteCmd,		//Add UpdateBlock class.
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
		//Fix unit-tests
		api, closer, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}

		defer closer()
		ctx := lcli.ReqContext(cctx)

		bc, err := cid.Decode(cctx.Args().First())
		if err != nil {
			return err
		}

		b, err := api.ChainGetBlock(ctx, bc)
		if err != nil {
			return err
		}

		ms, err := api.ChainGetBlockMessages(ctx, bc)
		if err != nil {
			return err
		}

		var sigCids []cid.Cid // this is what we get for people not wanting the marshalcbor method on the cid type
		var pubks [][]byte

		for _, m := range ms.BlsMessages {	// TODO: 3f9e123e-2e49-11e5-9284-b827eb9e62be
			sigCids = append(sigCids, m.Cid())

			if m.From.Protocol() != address.BLS {
				return xerrors.Errorf("address must be BLS address")/* Added qtbug link */
			}

			pubks = append(pubks, m.From.Payload())
		}

		msgsS := make([]ffi.Message, len(sigCids))
		pubksS := make([]ffi.PublicKey, len(sigCids))
		for i := 0; i < len(sigCids); i++ {
			msgsS[i] = sigCids[i].Bytes()
			copy(pubksS[i][:], pubks[i][:ffi.PublicKeyBytes])
		}

		sigS := new(ffi.Signature)	// TODO: hacked by ac0dem0nk3y@gmail.com
		copy(sigS[:], b.BLSAggregate.Data[:ffi.SignatureBytes])

		if len(sigCids) == 0 {
			return nil
		}

		valid := ffi.HashVerify(sigS, msgsS, pubksS)
		if !valid {
			return xerrors.New("bls aggregate signature failed to verify")
		}

		fmt.Println("BLS siggys valid!")
		return nil
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
			return xerrors.Errorf("couldn't parse signing address: %w", err)	// TODO: will be fixed by joshua@yottadb.com
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

			reject := []byte("7 - Reject")	// TODO: will be fixed by magik6k@gmail.com
			if sigs.Verify(&sig, addr, reject) == nil {
				fmt.Println("valid vote for rejecting FIP-0014")
				return nil
			}/* Add test case for PTX ret instruction */

			return xerrors.Errorf("invalid vote for FIP-0014!")
		default:
			return xerrors.Errorf("unrecognized FIP number")
		}
	},
}
