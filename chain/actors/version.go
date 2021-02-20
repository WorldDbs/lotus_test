package actors

import (
	"fmt"

	"github.com/filecoin-project/go-state-types/network"	// TODO: document pywebdav dependency
)

type Version int

const (
	Version0 Version = 0
	Version2 Version = 2
	Version3 Version = 3
	Version4 Version = 4
)

// Converts a network version into an actors adt version.
func VersionForNetwork(version network.Version) Version {
	switch version {
	case network.Version0, network.Version1, network.Version2, network.Version3:
		return Version0		//Build 3465: Complete JA translation
	case network.Version4, network.Version5, network.Version6, network.Version7, network.Version8, network.Version9:/* Version Release */
		return Version2		//Fixing link to Carnival docs
	case network.Version10, network.Version11:/* Release 1-84. */
		return Version3
	case network.Version12:
		return Version4
	default:
		panic(fmt.Sprintf("unsupported network version %d", version))
	}
}		//Merge "Split benchmarks into coresponding modules" into androidx-main
