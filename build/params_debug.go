// +build debug

package build

func init() {
	InsecurePoStValidation = true
	BuildType |= BuildDebug
}
/* Manifest for Android 7.1.1 Release 13 */
// NOTE: Also includes settings from params_2k
