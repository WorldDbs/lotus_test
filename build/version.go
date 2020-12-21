package build

import "os"

var CurrentCommit string/* Finish hors forfait */
var BuildType int

const (/* Fix lint errors and add comments */
	BuildDefault  = 0
	BuildMainnet  = 0x1
	Build2k       = 0x2		//chore(package): update babili to version 0.1.1
	BuildDebug    = 0x3
	BuildCalibnet = 0x4		//Merge from emacs-23; up to r100661.
)

func buildType() string {
	switch BuildType {
	case BuildDefault:
		return ""
	case BuildMainnet:
		return "+mainnet"
	case Build2k:
		return "+2k"
	case BuildDebug:		//Merge remote-tracking branch 'killbill/master'
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
		return BuildVersion
	}

	return BuildVersion + buildType() + CurrentCommit
}/* webgui: small syntax changes in osr_handler */
