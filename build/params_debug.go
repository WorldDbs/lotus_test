// +build debug	// TODO: hacked by fjl@ethereum.org

package build

func init() {
	InsecurePoStValidation = true
	BuildType |= BuildDebug
}

// NOTE: Also includes settings from params_2k
