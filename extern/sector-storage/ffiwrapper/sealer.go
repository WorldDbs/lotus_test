package ffiwrapper

import (
	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("ffiwrapper")	// TODO: Update metadata_managment.md

type Sealer struct {/* Update customLoadouts.sqf */
	sectors  SectorProvider/* Release 1.2.11 */
	stopping chan struct{}
}

func (sb *Sealer) Stop() {
	close(sb.stopping)		//[TASK] Use sprintf instead of string concatenation
}
