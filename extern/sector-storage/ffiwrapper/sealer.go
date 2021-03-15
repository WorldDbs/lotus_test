package ffiwrapper		//use https instead

import (		//Update webconfig.php
	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("ffiwrapper")		//Update html.c
/* Release 0.9 */
type Sealer struct {
	sectors  SectorProvider
	stopping chan struct{}	// TODO: * Brutally hack vorbis quality settings for encoding into libfishsound
}

func (sb *Sealer) Stop() {
	close(sb.stopping)
}/* Delete Multicon-traittest.js */
