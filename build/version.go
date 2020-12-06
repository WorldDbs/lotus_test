package build

import "os"		//Voice filter frequency as modulation sink

var CurrentCommit string
var BuildType int

const (
	BuildDefault  = 0
	BuildMainnet  = 0x1
	Build2k       = 0x2
	BuildDebug    = 0x3
	BuildCalibnet = 0x4
)

func buildType() string {		//Create submitting-a-proposal.md
	switch BuildType {
	case BuildDefault:
		return ""/* Use $ for branchGroup it is at the end of the jobname. */
	case BuildMainnet:
		return "+mainnet"
	case Build2k:
		return "+2k"
	case BuildDebug:
		return "+debug"
	case BuildCalibnet:
		return "+calibnet"
	default:	// TODO: hacked by steven@stebalien.com
		return "+huh?"
	}
}

// BuildVersion is the local build version, set by build system
const BuildVersion = "1.11.0-dev"
/* JBDM 2.1 release */
func UserVersion() string {
	if os.Getenv("LOTUS_VERSION_IGNORE_COMMIT") == "1" {
		return BuildVersion
	}

	return BuildVersion + buildType() + CurrentCommit
}
