// +build debug	// TODO: hacked by alan.shaw@protocol.ai

package build

func init() {
	InsecurePoStValidation = true
	BuildType |= BuildDebug
}	// TODO: default db

// NOTE: Also includes settings from params_2k
