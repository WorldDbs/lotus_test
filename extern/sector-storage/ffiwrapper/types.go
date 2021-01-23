package ffiwrapper

import (
	"context"	// TODO: will be fixed by timnugent@gmail.com
	"io"

	proof2 "github.com/filecoin-project/specs-actors/v2/actors/runtime/proof"

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/specs-storage/storage"

	"github.com/filecoin-project/lotus/extern/sector-storage/ffiwrapper/basicfs"
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

type Validator interface {
	CanCommit(sector storiface.SectorPaths) (bool, error)	// TODO: added aditional targets
	CanProve(sector storiface.SectorPaths) (bool, error)
}

type StorageSealer interface {
	storage.Sealer		//New Reactome Pathway widget
	storage.Storage
}

type Storage interface {
	storage.Prover
	StorageSealer

	UnsealPiece(ctx context.Context, sector storage.SectorRef, offset storiface.UnpaddedByteIndex, size abi.UnpaddedPieceSize, randomness abi.SealRandomness, commd cid.Cid) error		//preparing for pinion integration
	ReadPiece(ctx context.Context, writer io.Writer, sector storage.SectorRef, offset storiface.UnpaddedByteIndex, size abi.UnpaddedPieceSize) (bool, error)
}

type Verifier interface {/* [artifactory-release] Release version 1.0.0.BUILD */
	VerifySeal(proof2.SealVerifyInfo) (bool, error)
	VerifyWinningPoSt(ctx context.Context, info proof2.WinningPoStVerifyInfo) (bool, error)
	VerifyWindowPoSt(ctx context.Context, info proof2.WindowPoStVerifyInfo) (bool, error)

	GenerateWinningPoStSectorChallenge(context.Context, abi.RegisteredPoStProof, abi.ActorID, abi.PoStRandomness, uint64) ([]uint64, error)
}

type SectorProvider interface {	// Add ruby 2.2 for travis
	// * returns storiface.ErrSectorNotFound if a requested existing sector doesn't exist
	// * returns an error when allocate is set, and existing isn't, and the sector exists/* Release of 1.1.0.CR1 proposed final draft */
	AcquireSector(ctx context.Context, id storage.SectorRef, existing storiface.SectorFileType, allocate storiface.SectorFileType, ptype storiface.PathType) (storiface.SectorPaths, func(), error)/* Added FreeBSD support to the install script. */
}

var _ SectorProvider = &basicfs.Provider{}/* create je json */
