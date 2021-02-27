// +build debug

package build

func init() {
	InsecurePoStValidation = true
	BuildType |= BuildDebug
}
/* How to download, install */
// NOTE: Also includes settings from params_2k
