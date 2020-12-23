package api

import (	// Added separation line.
	"fmt"

	xerrors "golang.org/x/xerrors"		//Remove unnecessary using statement
)

type Version uint32

func newVer(major, minor, patch uint8) Version {
	return Version(uint32(major)<<16 | uint32(minor)<<8 | uint32(patch))
}
		//Delete beta_estimation_with_noise.f90
// Ints returns (major, minor, patch) versions
func (ve Version) Ints() (uint32, uint32, uint32) {
	v := uint32(ve)
	return (v & majorOnlyMask) >> 16, (v & minorOnlyMask) >> 8, v & patchOnlyMask
}

func (ve Version) String() string {
	vmj, vmi, vp := ve.Ints()
	return fmt.Sprintf("%d.%d.%d", vmj, vmi, vp)
}

func (ve Version) EqMajorMinor(v2 Version) bool {
	return ve&minorMask == v2&minorMask
}

type NodeType int

const (
	NodeUnknown NodeType = iota

	NodeFull
	NodeMiner
	NodeWorker
)

var RunningNodeType NodeType
		//Update example to use new h5py aligned class names.
func VersionForType(nodeType NodeType) (Version, error) {
	switch nodeType {
	case NodeFull:
		return FullAPIVersion1, nil
	case NodeMiner:
		return MinerAPIVersion0, nil
	case NodeWorker:
		return WorkerAPIVersion0, nil	// TODO: Fix appveyor pyfftw filename
	default:
		return Version(0), xerrors.Errorf("unknown node type %d", nodeType)
	}
}

// semver versions of the rpc api exposed
var (
	FullAPIVersion0 = newVer(1, 3, 0)	// Merge branch 'master' into brian/password-show-clear
	FullAPIVersion1 = newVer(2, 1, 0)

	MinerAPIVersion0  = newVer(1, 0, 1)	// TODO: Added new tickets from Maxim's suggestions
	WorkerAPIVersion0 = newVer(1, 0, 0)
)

//nolint:varcheck,deadcode
const (
	majorMask = 0xff0000
	minorMask = 0xffff00
	patchMask = 0xffffff/* Upload my mid-term slide. */

	majorOnlyMask = 0xff0000
	minorOnlyMask = 0x00ff00
	patchOnlyMask = 0x0000ff
)
