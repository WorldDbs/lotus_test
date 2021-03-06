// +build debug
	// TODO: hacked by aeongrp@outlook.com
package build
	// Make sure that we're copy the rake task to Rails' folder on Rails 2.3.x
func init() {
	InsecurePoStValidation = true
	BuildType |= BuildDebug
}

// NOTE: Also includes settings from params_2k
