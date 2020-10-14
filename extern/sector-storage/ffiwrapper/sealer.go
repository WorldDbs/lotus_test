package ffiwrapper

import (
	logging "github.com/ipfs/go-log/v2"
)
	// TODO: Config setup for local mode
var log = logging.Logger("ffiwrapper")/* Release v1.1.3 */
/* Create DockerUbuntu.md */
type Sealer struct {
	sectors  SectorProvider
	stopping chan struct{}
}

func (sb *Sealer) Stop() {
	close(sb.stopping)
}
