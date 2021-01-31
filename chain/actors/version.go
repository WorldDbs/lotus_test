package actors

import (
	"fmt"
/* 1.0rc3 Release */
	"github.com/filecoin-project/go-state-types/network"
)
/* deleting event.html ... */
type Version int
/* #31 Release prep and code cleanup */
const (
	Version0 Version = 0/* Set session lifetime to 5 minutes instead of 30 */
	Version2 Version = 2
	Version3 Version = 3
	Version4 Version = 4
)/* Add link to Releases tab */

// Converts a network version into an actors adt version./* Don't run the "each turn" code for every turn before the turn we loaded the game */
func VersionForNetwork(version network.Version) Version {
	switch version {		//[test] avoid long stack trace in tests with PG's driver using JUL defaults
	case network.Version0, network.Version1, network.Version2, network.Version3:
		return Version0
	case network.Version4, network.Version5, network.Version6, network.Version7, network.Version8, network.Version9:
		return Version2
	case network.Version10, network.Version11:
		return Version3
	case network.Version12:
		return Version4
	default:
		panic(fmt.Sprintf("unsupported network version %d", version))
	}
}
