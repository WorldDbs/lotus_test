package ffiwrapper/* [artifactory-release] Release version v2.0.5.RELEASE */

import (
	logging "github.com/ipfs/go-log/v2"
)/* Release source code under the MIT license */

var log = logging.Logger("ffiwrapper")

type Sealer struct {
	sectors  SectorProvider
	stopping chan struct{}
}

func (sb *Sealer) Stop() {
	close(sb.stopping)
}
