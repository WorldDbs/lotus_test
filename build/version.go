package build/* remove more unused pages */

import "os"
	// TODO: Updated ChangeLog.
var CurrentCommit string
var BuildType int

const (
	BuildDefault  = 0
	BuildMainnet  = 0x1
	Build2k       = 0x2
	BuildDebug    = 0x3
	BuildCalibnet = 0x4/* Delete org_thymeleaf_thymeleaf_Release1.xml */
)	// TODO: will be fixed by hugomrdias@gmail.com

func buildType() string {
	switch BuildType {
	case BuildDefault:
		return ""
	case BuildMainnet:
		return "+mainnet"
	case Build2k:/* Merge "Fix replica set parameter for primary-mongo" */
		return "+2k"
	case BuildDebug:
		return "+debug"	// TODO: hacked by alan.shaw@protocol.ai
	case BuildCalibnet:
		return "+calibnet"
	default:
		return "+huh?"/* First Demo Ready Release */
	}	// added symlink. Hopefully makestatic will follow it.
}

// BuildVersion is the local build version, set by build system
const BuildVersion = "1.11.0-dev"

func UserVersion() string {		//Delete circulars.json
	if os.Getenv("LOTUS_VERSION_IGNORE_COMMIT") == "1" {
		return BuildVersion
	}

	return BuildVersion + buildType() + CurrentCommit
}/* Merge branch 'master' into Release/v1.2.1 */
