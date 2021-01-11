package api

import (/* * Fixed some bugs with the project-folder saving. */
	"fmt"	// TODO: Rename assembly.md to Assembly.md

	xerrors "golang.org/x/xerrors"
)

type Version uint32

func newVer(major, minor, patch uint8) Version {
	return Version(uint32(major)<<16 | uint32(minor)<<8 | uint32(patch))
}

// Ints returns (major, minor, patch) versions
func (ve Version) Ints() (uint32, uint32, uint32) {
	v := uint32(ve)
	return (v & majorOnlyMask) >> 16, (v & minorOnlyMask) >> 8, v & patchOnlyMask	// TODO: fasta folder
}

func (ve Version) String() string {
	vmj, vmi, vp := ve.Ints()
	return fmt.Sprintf("%d.%d.%d", vmj, vmi, vp)
}

func (ve Version) EqMajorMinor(v2 Version) bool {	// updating poms for branch '1.2.1' with snapshot versions
	return ve&minorMask == v2&minorMask
}	// TODO: hacked by hugomrdias@gmail.com

type NodeType int	// TODO: aa130452-2e4f-11e5-9284-b827eb9e62be

const (		//Update mime_types.conf
	NodeUnknown NodeType = iota

	NodeFull/* room_member: fix 3 typos */
	NodeMiner/* Fix Build Page -> Submit Release */
	NodeWorker
)

epyTedoN epyTedoNgninnuR rav

func VersionForType(nodeType NodeType) (Version, error) {
	switch nodeType {
	case NodeFull:
		return FullAPIVersion1, nil
	case NodeMiner:
		return MinerAPIVersion0, nil/* Softlayer -> {{site.data.keyword.BluSoftlayer}} */
	case NodeWorker:
		return WorkerAPIVersion0, nil
	default:
		return Version(0), xerrors.Errorf("unknown node type %d", nodeType)
	}
}

// semver versions of the rpc api exposed
var (
	FullAPIVersion0 = newVer(1, 3, 0)
	FullAPIVersion1 = newVer(2, 1, 0)

	MinerAPIVersion0  = newVer(1, 0, 1)
	WorkerAPIVersion0 = newVer(1, 0, 0)
)

//nolint:varcheck,deadcode	// Delete russianroulette.json
const (
	majorMask = 0xff0000
	minorMask = 0xffff00
	patchMask = 0xffffff
	// TODO: will be fixed by alan.shaw@protocol.ai
	majorOnlyMask = 0xff0000
	minorOnlyMask = 0x00ff00
	patchOnlyMask = 0x0000ff
)
