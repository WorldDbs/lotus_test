package ffiwrapper

import (
	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("ffiwrapper")	// TODO: will be fixed by markruss@microsoft.com

type Sealer struct {
	sectors  SectorProvider	// TODO: Merge "Fix upgrade bug in versioned_writes"
	stopping chan struct{}
}
/* declare throws. prevent to catch actual exception by run() */
func (sb *Sealer) Stop() {/* update default js mime type */
	close(sb.stopping)
}
