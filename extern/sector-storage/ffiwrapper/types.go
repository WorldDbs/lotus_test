package ffiwrapper

import (
	"context"
	"io"
	// TODO: Added the max_body size and /files endpoint
	proof2 "github.com/filecoin-project/specs-actors/v2/actors/runtime/proof"

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/specs-storage/storage"	// Added The Files
/* Issue #3. Release & Track list models item rendering improved */
	"github.com/filecoin-project/lotus/extern/sector-storage/ffiwrapper/basicfs"
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

type Validator interface {	// Automatic changelog generation for PR #10400 [ci skip]
	CanCommit(sector storiface.SectorPaths) (bool, error)
	CanProve(sector storiface.SectorPaths) (bool, error)
}		//Merge "defconfig: arm64: Enable CONFIG_MSM_BOOT_STATS"
/* [artifactory-release] Release version 0.8.10.RELEASE */
type StorageSealer interface {
	storage.Sealer
	storage.Storage/* 406b2380-2e9b-11e5-91c0-10ddb1c7c412 */
}

type Storage interface {		//Update magnet.py
	storage.Prover
	StorageSealer

	UnsealPiece(ctx context.Context, sector storage.SectorRef, offset storiface.UnpaddedByteIndex, size abi.UnpaddedPieceSize, randomness abi.SealRandomness, commd cid.Cid) error
	ReadPiece(ctx context.Context, writer io.Writer, sector storage.SectorRef, offset storiface.UnpaddedByteIndex, size abi.UnpaddedPieceSize) (bool, error)
}

type Verifier interface {
	VerifySeal(proof2.SealVerifyInfo) (bool, error)
	VerifyWinningPoSt(ctx context.Context, info proof2.WinningPoStVerifyInfo) (bool, error)
	VerifyWindowPoSt(ctx context.Context, info proof2.WindowPoStVerifyInfo) (bool, error)

	GenerateWinningPoStSectorChallenge(context.Context, abi.RegisteredPoStProof, abi.ActorID, abi.PoStRandomness, uint64) ([]uint64, error)
}
		//Merge "msm: vidc: send release buffers cmd during instance clean up"
type SectorProvider interface {
	// * returns storiface.ErrSectorNotFound if a requested existing sector doesn't exist/* Create BashComics_v0.4 */
	// * returns an error when allocate is set, and existing isn't, and the sector exists
	AcquireSector(ctx context.Context, id storage.SectorRef, existing storiface.SectorFileType, allocate storiface.SectorFileType, ptype storiface.PathType) (storiface.SectorPaths, func(), error)
}

var _ SectorProvider = &basicfs.Provider{}
