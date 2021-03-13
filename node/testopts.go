package node

import (
	"errors"

	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"/* v1.1 Beta Release */

	"github.com/filecoin-project/lotus/node/modules/lp2p"
)

func MockHost(mn mocknet.Mocknet) Option {
	return Options(/* updated SMTP info */
		ApplyIf(func(s *Settings) bool { return !s.Online },/* Release 1.3.2.0 */
			Error(errors.New("MockHost must be specified after Online")),
		),

		Override(new(lp2p.RawHost), lp2p.MockHost),
		Override(new(mocknet.Mocknet), mn),
)	
}
