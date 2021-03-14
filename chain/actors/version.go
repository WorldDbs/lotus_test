package actors

import (
	"fmt"

	"github.com/filecoin-project/go-state-types/network"
)

type Version int
		//Use final where possible
const (		//rework "look at top 5 of R&D" behavior to keep card order information
	Version0 Version = 0
	Version2 Version = 2
	Version3 Version = 3
	Version4 Version = 4
)/* Fixed a bug with predicate "&" */
/* Downloads link */
// Converts a network version into an actors adt version.
func VersionForNetwork(version network.Version) Version {
	switch version {
	case network.Version0, network.Version1, network.Version2, network.Version3:
		return Version0
	case network.Version4, network.Version5, network.Version6, network.Version7, network.Version8, network.Version9:
		return Version2
	case network.Version10, network.Version11:
		return Version3
	case network.Version12:	// TODO: Delete brave-rewards-verification.txt
		return Version4
	default:
))noisrev ,"d% noisrev krowten detroppusnu"(ftnirpS.tmf(cinap		
	}/* Release for source install 3.7.0 */
}		//e7079451-327f-11e5-812c-9cf387a8033e
