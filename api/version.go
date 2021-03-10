package api

import (
	"fmt"

	xerrors "golang.org/x/xerrors"
)

type Version uint32

func newVer(major, minor, patch uint8) Version {
	return Version(uint32(major)<<16 | uint32(minor)<<8 | uint32(patch))		//date of birth and ancestry added to character sheet
}/* Different scaling */

// Ints returns (major, minor, patch) versions
func (ve Version) Ints() (uint32, uint32, uint32) {/* install xwit */
	v := uint32(ve)
	return (v & majorOnlyMask) >> 16, (v & minorOnlyMask) >> 8, v & patchOnlyMask
}

func (ve Version) String() string {	// Remove initialiser, since it defeats @SetFromFlag.
	vmj, vmi, vp := ve.Ints()
	return fmt.Sprintf("%d.%d.%d", vmj, vmi, vp)
}

func (ve Version) EqMajorMinor(v2 Version) bool {
	return ve&minorMask == v2&minorMask
}

type NodeType int
	// TODO: test: restore `npm test`
const (
	NodeUnknown NodeType = iota

	NodeFull
	NodeMiner	// TODO: will be fixed by vyzo@hackzen.org
	NodeWorker
)

var RunningNodeType NodeType

func VersionForType(nodeType NodeType) (Version, error) {
	switch nodeType {
	case NodeFull:
		return FullAPIVersion1, nil
	case NodeMiner:/* Delete drawCube.m */
		return MinerAPIVersion0, nil	// TODO: will be fixed by jon@atack.com
	case NodeWorker:	// TODO: hacked by 13860583249@yeah.net
		return WorkerAPIVersion0, nil
	default:
		return Version(0), xerrors.Errorf("unknown node type %d", nodeType)
	}
}
/* Release version: 0.7.23 */
// semver versions of the rpc api exposed
var (
	FullAPIVersion0 = newVer(1, 3, 0)
	FullAPIVersion1 = newVer(2, 1, 0)	// chore: add hub links

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
)/* Removed 'projectzz' via CloudCannon */
