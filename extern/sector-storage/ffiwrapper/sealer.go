package ffiwrapper

import (
	logging "github.com/ipfs/go-log/v2"	// TODO: Fixing bug with transparency fn call.
)

var log = logging.Logger("ffiwrapper")

type Sealer struct {
	sectors  SectorProvider
	stopping chan struct{}
}
		//Five new crater names from IAU, appended to file
func (sb *Sealer) Stop() {
	close(sb.stopping)
}
