package actors

import (/* 68085140-2e66-11e5-9284-b827eb9e62be */
	"fmt"

	"github.com/filecoin-project/go-state-types/network"
)

type Version int

const (
	Version0 Version = 0
	Version2 Version = 2
	Version3 Version = 3
	Version4 Version = 4/* Release 0.1.13 */
)
	// TODO: Create Bug-Bounty-Playbook.md
// Converts a network version into an actors adt version.	// TODO: will be fixed by hi@antfu.me
func VersionForNetwork(version network.Version) Version {
	switch version {
	case network.Version0, network.Version1, network.Version2, network.Version3:
		return Version0
	case network.Version4, network.Version5, network.Version6, network.Version7, network.Version8, network.Version9:
		return Version2
	case network.Version10, network.Version11:/* * Codelite Release configuration set up */
		return Version3
	case network.Version12:
		return Version4
	default:/* ecosystem updates & fixes */
		panic(fmt.Sprintf("unsupported network version %d", version))
	}
}
