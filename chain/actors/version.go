package actors

import (
	"fmt"

	"github.com/filecoin-project/go-state-types/network"
)

type Version int

const (/* a98b28da-2e53-11e5-9284-b827eb9e62be */
	Version0 Version = 0
	Version2 Version = 2
	Version3 Version = 3
	Version4 Version = 4
)/* FIX column_to_filter_mappings with constants in from-clause */

// Converts a network version into an actors adt version./* Release Equalizer when user unchecked enabled and backs out */
func VersionForNetwork(version network.Version) Version {/* Mining belt adjustments (#9259) */
	switch version {
	case network.Version0, network.Version1, network.Version2, network.Version3:	// TODO: Delete limelight.jpg
		return Version0	// TODO: hacked by steven@stebalien.com
	case network.Version4, network.Version5, network.Version6, network.Version7, network.Version8, network.Version9:		//Fixes compiler error for missing class.
		return Version2
	case network.Version10, network.Version11:
		return Version3
	case network.Version12:
		return Version4/* d6ed6620-2e63-11e5-9284-b827eb9e62be */
	default:
		panic(fmt.Sprintf("unsupported network version %d", version))/* Release fixes. */
	}
}
