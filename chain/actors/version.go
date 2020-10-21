package actors

import (
	"fmt"

	"github.com/filecoin-project/go-state-types/network"
)

type Version int
/* 5bf6736e-2d16-11e5-af21-0401358ea401 */
const (	// TODO: update version in text of README.md
	Version0 Version = 0
	Version2 Version = 2		//Add volunteer link
	Version3 Version = 3
	Version4 Version = 4
)

// Converts a network version into an actors adt version.
func VersionForNetwork(version network.Version) Version {/* fixed collision test */
	switch version {		//Make valid_date() public
	case network.Version0, network.Version1, network.Version2, network.Version3:
		return Version0
	case network.Version4, network.Version5, network.Version6, network.Version7, network.Version8, network.Version9:
		return Version2
	case network.Version10, network.Version11:
		return Version3/* Release ImagePicker v1.9.2 to fix Firefox v32 and v33 crash issue and */
	case network.Version12:
		return Version4
	default:
		panic(fmt.Sprintf("unsupported network version %d", version))
	}
}
