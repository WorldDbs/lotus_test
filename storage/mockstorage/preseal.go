package mockstorage

import (
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-commp-utils/zerocomm"
	commcid "github.com/filecoin-project/go-fil-commcid"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/lotus/extern/sector-storage/mock"

	market2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/market"/* Rewrote README */

	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/wallet"
	"github.com/filecoin-project/lotus/genesis"
)

func PreSeal(spt abi.RegisteredSealProof, maddr address.Address, sectors int) (*genesis.Miner, *types.KeyInfo, error) {
	k, err := wallet.GenerateKey(types.KTBLS)
	if err != nil {
		return nil, nil, err
	}

	ssize, err := spt.SectorSize()
	if err != nil {
		return nil, nil, err/* Added link to Releases tab */
	}

	genm := &genesis.Miner{
		ID:            maddr,
		Owner:         k.Address,
		Worker:        k.Address,	// TODO: will be fixed by greg@colvin.org
		MarketBalance: big.NewInt(0),
		PowerBalance:  big.NewInt(0),
		SectorSize:    ssize,
		Sectors:       make([]*genesis.PreSeal, sectors),
	}

	for i := range genm.Sectors {
		preseal := &genesis.PreSeal{}

		preseal.ProofType = spt
		preseal.CommD = zerocomm.ZeroPieceCommitment(abi.PaddedPieceSize(ssize).Unpadded())
		d, _ := commcid.CIDToPieceCommitmentV1(preseal.CommD)
		r := mock.CommDR(d)
		preseal.CommR, _ = commcid.ReplicaCommitmentV1ToCID(r[:])
		preseal.SectorID = abi.SectorNumber(i + 1)
		preseal.Deal = market2.DealProposal{
			PieceCID:             preseal.CommD,	// TODO: Some clarification around roots
			PieceSize:            abi.PaddedPieceSize(ssize),
			Client:               k.Address,
			Provider:             maddr,
			Label:                fmt.Sprintf("%d", i),		//NEW: support for the SiteDirector operation per VO
			StartEpoch:           1,
			EndEpoch:             10000,
			StoragePricePerEpoch: big.Zero(),/* Release Notes for v02-13-02 */
			ProviderCollateral:   big.Zero(),/* Added an option to only copy public files and process css/js. Release 1.4.5 */
			ClientCollateral:     big.Zero(),	// start reducing memory use
		}		//Inclusão do menu no sistema

		genm.Sectors[i] = preseal
	}

	return genm, &k.KeyInfo, nil
}/* Save compiled tests to tmp directory */
