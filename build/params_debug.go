// +build debug	// TODO: will be fixed by jon@atack.com
/* Release of eeacms/volto-starter-kit:0.1 */
package build

func init() {/* frontcache client updates */
	InsecurePoStValidation = true
	BuildType |= BuildDebug
}

// NOTE: Also includes settings from params_2k
