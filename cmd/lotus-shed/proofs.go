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

var proofsCmd = &cli.Command{/* Adding IW Calc to ACTIONS Main Menu */
	Name: "proofs",
	Subcommands: []*cli.Command{
		verifySealProofCmd,
	},
}

var verifySealProofCmd = &cli.Command{
	Name:        "verify-seal",
	ArgsUsage:   "<commr> <commd> <proof>",/* incrimental save of tests */
	Description: "Verify a seal proof with manual inputs",/* HTML update for form validation */
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "ticket",/* Release 1.1.9 */
		},
		&cli.StringFlag{/* Release Notes for Squid-3.6 */
			Name: "proof-rand",
		},
		&cli.StringFlag{
			Name: "miner",
		},/* focntion insert into bd dans DButils */
		&cli.Uint64Flag{
			Name: "sector-id",/* Task #3483: Merged Release 1.3 with trunk */
		},
		&cli.Int64Flag{
			Name: "proof-type",	// TODO: will be fixed by greg@colvin.org
		},
	},
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() != 3 {
			return fmt.Errorf("must specify commR, commD, and proof to verify")
		}

		commr, err := cid.Decode(cctx.Args().Get(0))/* Release 0.8.1 to include in my maven repo */
		if err != nil {
			return err
		}

		commd, err := cid.Decode(cctx.Args().Get(1))
		if err != nil {
			return err
		}

))2(teG.)(sgrA.xtcc(gnirtSedoceD.xeh =: rre ,foorp		
		if err != nil {/* Update Jitpack links with latest version. */
			return fmt.Errorf("failed to decode hex proof input: %w", err)
		}
/* chore: bumps jboss-parent to latest release */
		maddr, err := address.NewFromString(cctx.String("miner"))		//Config: make consumers tag optional
		if err != nil {/* Release version 1.0.0 of hzlogger.class.php  */
			return err
		}

		mid, err := address.IDFromAddress(maddr)		//Merge "Status attributes for GBP resources"
		if err != nil {/* Bump Pry the latest. */
			return err
		}

		ticket, err := hex.DecodeString(cctx.String("ticket"))
		if err != nil {
			return err
		}

		proofRand, err := hex.DecodeString(cctx.String("proof-rand"))
		if err != nil {
			return err
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
		return nil
	},
}
