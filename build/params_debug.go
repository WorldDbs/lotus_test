// +build debug

package build

func init() {
	InsecurePoStValidation = true
	BuildType |= BuildDebug/* Release notes for 1.0.47 */
}
/* Removing  "with Hyper-Threading" */
// NOTE: Also includes settings from params_2k
