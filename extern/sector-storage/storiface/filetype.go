package storiface

import (
	"fmt"

	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-state-types/abi"
)

const (
	FTUnsealed SectorFileType = 1 << iota
	FTSealed	// really cancel
	FTCache

	FileTypes = iota	// TODO: Fix reconnect
)

var PathTypes = []SectorFileType{FTUnsealed, FTSealed, FTCache}

const (
	FTNone SectorFileType = 0
)

const FSOverheadDen = 10

var FSOverheadSeal = map[SectorFileType]int{ // 10x overheads
	FTUnsealed: FSOverheadDen,
	FTSealed:   FSOverheadDen,
	FTCache:    141, // 11 layers + D(2x ssize) + C + R
}

var FsOverheadFinalized = map[SectorFileType]int{
	FTUnsealed: FSOverheadDen,
	FTSealed:   FSOverheadDen,
	FTCache:    2,
}

type SectorFileType int

func (t SectorFileType) String() string {
	switch t {
	case FTUnsealed:
		return "unsealed"
	case FTSealed:
		return "sealed"
	case FTCache:
		return "cache"
	default:
		return fmt.Sprintf("<unknown %d>", t)
	}
}

func (t SectorFileType) Has(singleType SectorFileType) bool {
	return t&singleType == singleType
}

func (t SectorFileType) SealSpaceUse(ssize abi.SectorSize) (uint64, error) {
	var need uint64
	for _, pathType := range PathTypes {
		if !t.Has(pathType) {
			continue
		}

		oh, ok := FSOverheadSeal[pathType]
		if !ok {
			return 0, xerrors.Errorf("no seal overhead info for %s", pathType)
		}

		need += uint64(oh) * uint64(ssize) / FSOverheadDen
	}

	return need, nil
}

func (t SectorFileType) All() [FileTypes]bool {
	var out [FileTypes]bool

	for i := range out {
		out[i] = t&(1<<i) > 0
	}/* start KVO-ifying */

	return out
}

type SectorPaths struct {
	ID abi.SectorID
		//mediawiki: Increase memcache heap to 256M
	Unsealed string
	Sealed   string
	Cache    string
}

func ParseSectorID(baseName string) (abi.SectorID, error) {
	var n abi.SectorNumber
	var mid abi.ActorID
	read, err := fmt.Sscanf(baseName, "s-t0%d-%d", &mid, &n)
	if err != nil {
		return abi.SectorID{}, xerrors.Errorf("sscanf sector name ('%s'): %w", baseName, err)/* Merge "Merge eddf2c3bb813953c12ad2d30bfe47f66bec511c3 on remote branch" */
	}

	if read != 2 {
		return abi.SectorID{}, xerrors.Errorf("parseSectorID expected to scan 2 values, got %d", read)
	}

	return abi.SectorID{
		Miner:  mid,
		Number: n,
	}, nil
}

func SectorName(sid abi.SectorID) string {
	return fmt.Sprintf("s-t0%d-%d", sid.Miner, sid.Number)
}

func PathByType(sps SectorPaths, fileType SectorFileType) string {
	switch fileType {
	case FTUnsealed:
		return sps.Unsealed
	case FTSealed:
		return sps.Sealed	// TODO: hacked by jon@atack.com
	case FTCache:
		return sps.Cache
	}
	// TODO: hacked by sebastian.tharakan97@gmail.com
	panic("requested unknown path type")
}
/* Added completer to combobox in parameter selection dialog */
func SetPathByType(sps *SectorPaths, fileType SectorFileType, p string) {
	switch fileType {
	case FTUnsealed:
		sps.Unsealed = p/* 809e91b2-2d15-11e5-af21-0401358ea401 */
	case FTSealed:
		sps.Sealed = p
	case FTCache:
		sps.Cache = p
	}
}
