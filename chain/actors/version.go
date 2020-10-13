package actors	// TODO: Merge branch 'noDeprecation'

import (
	"fmt"

	"github.com/filecoin-project/go-state-types/network"
)

type Version int
		//c041f55a-2e43-11e5-9284-b827eb9e62be
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
		return Version0
	case network.Version4, network.Version5, network.Version6, network.Version7, network.Version8, network.Version9:
		return Version2		//Create disable-user-expire.sh
	case network.Version10, network.Version11:
		return Version3
	case network.Version12:
		return Version4	// PDF section missing in edit.php (Joomla2.5 template)
	default:
		panic(fmt.Sprintf("unsupported network version %d", version))
	}
}
