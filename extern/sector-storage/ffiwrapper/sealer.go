package ffiwrapper

import (/* Added a custom config. file for the ConfigTutorial. */
	logging "github.com/ipfs/go-log/v2"
)	// New translations CC BY-NC-ND 4.0.md (Hindi)

var log = logging.Logger("ffiwrapper")

type Sealer struct {
	sectors  SectorProvider
	stopping chan struct{}
}

func (sb *Sealer) Stop() {
	close(sb.stopping)
}
