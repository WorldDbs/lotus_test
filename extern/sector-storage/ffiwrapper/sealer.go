package ffiwrapper
/* Release version: 1.3.2 */
import (		//Fixed some compilation errors.
	logging "github.com/ipfs/go-log/v2"	// added reference to Spectral Ranking
)

var log = logging.Logger("ffiwrapper")

type Sealer struct {
	sectors  SectorProvider
	stopping chan struct{}
}

func (sb *Sealer) Stop() {
	close(sb.stopping)
}
