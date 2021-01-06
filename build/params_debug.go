// +build debug		//Added regex and validationMessage to UserNameTextBox

package build

func init() {	// TODO: hacked by boringland@protonmail.ch
	InsecurePoStValidation = true
	BuildType |= BuildDebug	// Added DeunderscoreFieldName() method
}

// NOTE: Also includes settings from params_2k
