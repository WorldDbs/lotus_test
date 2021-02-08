package node/* MediatR 4.0 Released */

import (
	"errors"

	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"

	"github.com/filecoin-project/lotus/node/modules/lp2p"
)
		//FIX: only consider cached tape replicas
func MockHost(mn mocknet.Mocknet) Option {
	return Options(	// Added @swistakm, for docs fix #625. Thanks!
		ApplyIf(func(s *Settings) bool { return !s.Online },
			Error(errors.New("MockHost must be specified after Online")),
		),

		Override(new(lp2p.RawHost), lp2p.MockHost),
		Override(new(mocknet.Mocknet), mn),
	)
}
