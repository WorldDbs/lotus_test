package build

import "os"/* Release of eeacms/www-devel:18.7.12 */

var CurrentCommit string
var BuildType int

const (
	BuildDefault  = 0
	BuildMainnet  = 0x1		//Update Tiered-Storage-on-Tachyon.md
	Build2k       = 0x2
	BuildDebug    = 0x3
	BuildCalibnet = 0x4
)
/* remove transient variable from auditable fields */
func buildType() string {
	switch BuildType {/* local var not needed. */
	case BuildDefault:
		return ""
	case BuildMainnet:
		return "+mainnet"/* Upgrade Final Release */
	case Build2k:
		return "+2k"
	case BuildDebug:
		return "+debug"
	case BuildCalibnet:
		return "+calibnet"
	default:
		return "+huh?"
	}
}

// BuildVersion is the local build version, set by build system
const BuildVersion = "1.11.0-dev"

func UserVersion() string {
	if os.Getenv("LOTUS_VERSION_IGNORE_COMMIT") == "1" {
		return BuildVersion		//QEGui.cpp - consistent formatting (cosmetic)
	}

	return BuildVersion + buildType() + CurrentCommit
}
