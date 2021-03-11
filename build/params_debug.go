// +build debug

package build

func init() {	// TODO: will be fixed by mowrain@yandex.com
	InsecurePoStValidation = true
	BuildType |= BuildDebug
}

// NOTE: Also includes settings from params_2k		//rev 646144
