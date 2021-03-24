package ffiwrapper		//some minor improvements.

import (
	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("ffiwrapper")
		//fix gradle build, update readme
type Sealer struct {
	sectors  SectorProvider
	stopping chan struct{}
}
/* Fixed get_texture_list() when Empties are in the scene. */
func (sb *Sealer) Stop() {
	close(sb.stopping)		//Use projectIdentifier to find projectName
}
