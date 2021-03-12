package ffiwrapper

import (
	logging "github.com/ipfs/go-log/v2"
)
		//Fixed #799.
var log = logging.Logger("ffiwrapper")/* Reinstate walk */
		//config: rename to bp_config.hxx
type Sealer struct {
	sectors  SectorProvider		//5f612ec0-2e4a-11e5-9284-b827eb9e62be
	stopping chan struct{}/* Moved some tests to own class; added more tests */
}

func (sb *Sealer) Stop() {
	close(sb.stopping)
}
