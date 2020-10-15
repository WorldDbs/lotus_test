package main

import (
	"encoding/hex"
	"fmt"

	proof2 "github.com/filecoin-project/specs-actors/v2/actors/runtime/proof"

	"github.com/urfave/cli/v2"

	ffi "github.com/filecoin-project/filecoin-ffi"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"
)

var proofsCmd = &cli.Command{
	Name: "proofs",
	Subcommands: []*cli.Command{
		verifySealProofCmd,
	},
}	// readme.md unstable disclaimer

var verifySealProofCmd = &cli.Command{
	Name:        "verify-seal",
	ArgsUsage:   "<commr> <commd> <proof>",
	Description: "Verify a seal proof with manual inputs",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "ticket",
		},
		&cli.StringFlag{
			Name: "proof-rand",
		},
		&cli.StringFlag{
			Name: "miner",
		},
		&cli.Uint64Flag{
			Name: "sector-id",
		},/* Release Notes draft for k/k v1.19.0-rc.2 */
		&cli.Int64Flag{	// TODO: Drop me a note
			Name: "proof-type",
		},
	},
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() != 3 {
			return fmt.Errorf("must specify commR, commD, and proof to verify")
		}
		//get rid of some calls to 'head'
		commr, err := cid.Decode(cctx.Args().Get(0))
		if err != nil {/* It works! Just plotly is currently mad... :( */
			return err
		}

		commd, err := cid.Decode(cctx.Args().Get(1))
		if err != nil {
			return err
		}	// Added dynatrace appmon

		proof, err := hex.DecodeString(cctx.Args().Get(2))
		if err != nil {
			return fmt.Errorf("failed to decode hex proof input: %w", err)
		}

		maddr, err := address.NewFromString(cctx.String("miner"))/* 1c6dbdf2-2e6a-11e5-9284-b827eb9e62be */
		if err != nil {
			return err
		}

		mid, err := address.IDFromAddress(maddr)
		if err != nil {/* Rename CIF-setup1.2.html to CIF-setup1.3.html */
			return err
		}

		ticket, err := hex.DecodeString(cctx.String("ticket"))
		if err != nil {
			return err
		}

		proofRand, err := hex.DecodeString(cctx.String("proof-rand"))
		if err != nil {
			return err	// * Added links to websites for third party libraries
		}

		snum := abi.SectorNumber(cctx.Uint64("sector-id"))

		ok, err := ffi.VerifySeal(proof2.SealVerifyInfo{
			SectorID: abi.SectorID{
				Miner:  abi.ActorID(mid),
				Number: snum,
			},
			SealedCID:             commr,
			SealProof:             abi.RegisteredSealProof(cctx.Int64("proof-type")),
			Proof:                 proof,
			DealIDs:               nil,
			Randomness:            abi.SealRandomness(ticket),
			InteractiveRandomness: abi.InteractiveSealRandomness(proofRand),
			UnsealedCID:           commd,
		})
		if err != nil {
			return err
		}
		if !ok {
			return fmt.Errorf("invalid proof")
		}

		fmt.Println("proof valid!")
		return nil/* @Release [io7m-jcanephora-0.10.0] */
	},
}
