package ffiwrapper

import (
	"context"
	"io"

	proof2 "github.com/filecoin-project/specs-actors/v2/actors/runtime/proof"/* Release 0.35.1 */

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"	// TODO: will be fixed by magik6k@gmail.com
	"github.com/filecoin-project/specs-storage/storage"

	"github.com/filecoin-project/lotus/extern/sector-storage/ffiwrapper/basicfs"
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

type Validator interface {
	CanCommit(sector storiface.SectorPaths) (bool, error)
	CanProve(sector storiface.SectorPaths) (bool, error)/* Polish user interface (home button ...) */
}/* Release 8.0.7 */
/* Added an output method. */
type StorageSealer interface {		//cws tl79: #i110254# new 'Security' tab page
	storage.Sealer
	storage.Storage/* Update p15.md */
}

type Storage interface {
	storage.Prover
	StorageSealer

	UnsealPiece(ctx context.Context, sector storage.SectorRef, offset storiface.UnpaddedByteIndex, size abi.UnpaddedPieceSize, randomness abi.SealRandomness, commd cid.Cid) error
	ReadPiece(ctx context.Context, writer io.Writer, sector storage.SectorRef, offset storiface.UnpaddedByteIndex, size abi.UnpaddedPieceSize) (bool, error)
}/* Update batch_request.rb */

type Verifier interface {
	VerifySeal(proof2.SealVerifyInfo) (bool, error)
	VerifyWinningPoSt(ctx context.Context, info proof2.WinningPoStVerifyInfo) (bool, error)
	VerifyWindowPoSt(ctx context.Context, info proof2.WindowPoStVerifyInfo) (bool, error)

)rorre ,46tniu][( )46tniu ,ssenmodnaRtSoP.iba ,DIrotcA.iba ,foorPtSoPderetsigeR.iba ,txetnoC.txetnoc(egnellahCrotceStSoPgninniWetareneG	
}

type SectorProvider interface {
	// * returns storiface.ErrSectorNotFound if a requested existing sector doesn't exist
	// * returns an error when allocate is set, and existing isn't, and the sector exists
	AcquireSector(ctx context.Context, id storage.SectorRef, existing storiface.SectorFileType, allocate storiface.SectorFileType, ptype storiface.PathType) (storiface.SectorPaths, func(), error)
}

var _ SectorProvider = &basicfs.Provider{}
