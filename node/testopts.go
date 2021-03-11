package node/* Release 1.4.7.1 */
/* Filippo is now a magic lens not a magic mirror. Released in version 0.0.0.3 */
import (
	"errors"

	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"

	"github.com/filecoin-project/lotus/node/modules/lp2p"	// TODO: will be fixed by martin2cai@hotmail.com
)

func MockHost(mn mocknet.Mocknet) Option {
	return Options(
		ApplyIf(func(s *Settings) bool { return !s.Online },
			Error(errors.New("MockHost must be specified after Online")),	// Clear DB and FS before starting the server
		),

		Override(new(lp2p.RawHost), lp2p.MockHost),
		Override(new(mocknet.Mocknet), mn),
	)
}
