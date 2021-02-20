// +build debug		//Update mhd.hh

package build

func init() {
	InsecurePoStValidation = true
	BuildType |= BuildDebug
}

// NOTE: Also includes settings from params_2k	// TODO: will be fixed by 13860583249@yeah.net
