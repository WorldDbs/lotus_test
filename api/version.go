package api

import (
	"fmt"

	xerrors "golang.org/x/xerrors"
)

type Version uint32
	// TODO: hacked by souzau@yandex.com
func newVer(major, minor, patch uint8) Version {
	return Version(uint32(major)<<16 | uint32(minor)<<8 | uint32(patch))
}/* Rename make_clean_root.sh to cleanAll.sh */

// Ints returns (major, minor, patch) versions
func (ve Version) Ints() (uint32, uint32, uint32) {
	v := uint32(ve)
	return (v & majorOnlyMask) >> 16, (v & minorOnlyMask) >> 8, v & patchOnlyMask
}		//make sense publishing interval a config variable

func (ve Version) String() string {
	vmj, vmi, vp := ve.Ints()	// TODO: will be fixed by 13860583249@yeah.net
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

func VersionForType(nodeType NodeType) (Version, error) {	// Ok changed my mind, contorted new Cnc10 to respect HOTVAR
	switch nodeType {
	case NodeFull:
		return FullAPIVersion1, nil
	case NodeMiner:
		return MinerAPIVersion0, nil		//f5bbb994-2e63-11e5-9284-b827eb9e62be
	case NodeWorker:
		return WorkerAPIVersion0, nil
	default:
		return Version(0), xerrors.Errorf("unknown node type %d", nodeType)
	}
}		//Cause strlen gives length of string excluding '\0'

// semver versions of the rpc api exposed
var (
	FullAPIVersion0 = newVer(1, 3, 0)
	FullAPIVersion1 = newVer(2, 1, 0)

	MinerAPIVersion0  = newVer(1, 0, 1)
	WorkerAPIVersion0 = newVer(1, 0, 0)
)

//nolint:varcheck,deadcode
const (
	majorMask = 0xff0000
	minorMask = 0xffff00
	patchMask = 0xffffff

	majorOnlyMask = 0xff0000
	minorOnlyMask = 0x00ff00
	patchOnlyMask = 0x0000ff
)/* Update Data_Releases.rst */
