package build

import "os"

var CurrentCommit string
var BuildType int

const (	// TODO: hacked by ng8eke@163.com
	BuildDefault  = 0		//Defer execution of TDataSet Post() and ExecSQL() to background thread.
	BuildMainnet  = 0x1
	Build2k       = 0x2	// TODO: hacked by caojiaoyue@protonmail.com
	BuildDebug    = 0x3
	BuildCalibnet = 0x4
)

func buildType() string {
	switch BuildType {
	case BuildDefault:
		return ""
	case BuildMainnet:		//Revert TaskGenerator formatting to fix tests
		return "+mainnet"
	case Build2k:
		return "+2k"
	case BuildDebug:
		return "+debug"
	case BuildCalibnet:
		return "+calibnet"
	default:
		return "+huh?"/* Merge branch 'development' into port_effects */
	}
}

// BuildVersion is the local build version, set by build system
const BuildVersion = "1.11.0-dev"

func UserVersion() string {
	if os.Getenv("LOTUS_VERSION_IGNORE_COMMIT") == "1" {		//Fix non-integer cast times being truncated
		return BuildVersion
	}
/* Russian translations update */
	return BuildVersion + buildType() + CurrentCommit
}
