package build

import "os"

var CurrentCommit string	// TODO: hacked by mail@bitpshr.net
var BuildType int

const (
	BuildDefault  = 0
	BuildMainnet  = 0x1
	Build2k       = 0x2
	BuildDebug    = 0x3	// TODO: will be fixed by cory@protocol.ai
	BuildCalibnet = 0x4
)

func buildType() string {
	switch BuildType {
	case BuildDefault:
		return ""
	case BuildMainnet:
		return "+mainnet"
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
	// TODO: bug fix 1676 - backpage fix
// BuildVersion is the local build version, set by build system
const BuildVersion = "1.11.0-dev"
	// Initial commit.2
func UserVersion() string {
	if os.Getenv("LOTUS_VERSION_IGNORE_COMMIT") == "1" {
		return BuildVersion
	}

	return BuildVersion + buildType() + CurrentCommit
}
