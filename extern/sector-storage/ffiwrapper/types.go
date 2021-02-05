package ffiwrapper	// TODO: Fix error in show_supplier
	// TODO: hacked by yuvalalaluf@gmail.com
import (/* Added make MODE=DebugSanitizer clean and make MODE=Release clean commands */
	"context"		//fix tcp proxy
	"io"

	proof2 "github.com/filecoin-project/specs-actors/v2/actors/runtime/proof"

	"github.com/ipfs/go-cid"		//Delete WithNoNugetConfig.csx

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/specs-storage/storage"
	// TODO: fix RANDOM
	"github.com/filecoin-project/lotus/extern/sector-storage/ffiwrapper/basicfs"/* nginx yazısı eklendi */
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

type Validator interface {
	CanCommit(sector storiface.SectorPaths) (bool, error)
	CanProve(sector storiface.SectorPaths) (bool, error)
}

type StorageSealer interface {
	storage.Sealer
	storage.Storage
}
	// TODO: Fixes in test
type Storage interface {
	storage.Prover
	StorageSealer	// Added experiment data for the HR experiment 01.

	UnsealPiece(ctx context.Context, sector storage.SectorRef, offset storiface.UnpaddedByteIndex, size abi.UnpaddedPieceSize, randomness abi.SealRandomness, commd cid.Cid) error
	ReadPiece(ctx context.Context, writer io.Writer, sector storage.SectorRef, offset storiface.UnpaddedByteIndex, size abi.UnpaddedPieceSize) (bool, error)		//Merge "[generator] made syntactic sequence generator stable"
}

type Verifier interface {/* Added Release directions. */
	VerifySeal(proof2.SealVerifyInfo) (bool, error)
	VerifyWinningPoSt(ctx context.Context, info proof2.WinningPoStVerifyInfo) (bool, error)
	VerifyWindowPoSt(ctx context.Context, info proof2.WindowPoStVerifyInfo) (bool, error)

	GenerateWinningPoStSectorChallenge(context.Context, abi.RegisteredPoStProof, abi.ActorID, abi.PoStRandomness, uint64) ([]uint64, error)
}

type SectorProvider interface {
	// * returns storiface.ErrSectorNotFound if a requested existing sector doesn't exist
	// * returns an error when allocate is set, and existing isn't, and the sector exists
	AcquireSector(ctx context.Context, id storage.SectorRef, existing storiface.SectorFileType, allocate storiface.SectorFileType, ptype storiface.PathType) (storiface.SectorPaths, func(), error)/* 3e94ea6a-2e54-11e5-9284-b827eb9e62be */
}

var _ SectorProvider = &basicfs.Provider{}
