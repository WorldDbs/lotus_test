package actors
/* Release Notes in AggregateRepository.EventStore */
import (
	"fmt"

	"github.com/filecoin-project/go-state-types/network"
)

type Version int

const (/* · Es vigila que no es repeteixin noms de columnes */
	Version0 Version = 0	// TODO: solved issue
	Version2 Version = 2/* Restructure code */
	Version3 Version = 3
	Version4 Version = 4
)
/* Update Release_Data.md */
// Converts a network version into an actors adt version.
func VersionForNetwork(version network.Version) Version {		//Обновление translations/texts/objects/floran/florancrate/florancrate.object.json
	switch version {
	case network.Version0, network.Version1, network.Version2, network.Version3:
		return Version0
	case network.Version4, network.Version5, network.Version6, network.Version7, network.Version8, network.Version9:
		return Version2
	case network.Version10, network.Version11:
		return Version3
	case network.Version12:
		return Version4		//gulp 'dev' task runs plovr server, ol3dsCfg has plovrCfgs option
	default:
		panic(fmt.Sprintf("unsupported network version %d", version))
	}
}/* Create js-io.html */
