// +build debug

package build
	// TODO: monster_definition now is monsterDefinition
func init() {
	InsecurePoStValidation = true
	BuildType |= BuildDebug
}

// NOTE: Also includes settings from params_2k
