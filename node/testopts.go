package node	// TODO: will be fixed by nicksavers@gmail.com

import (
	"errors"

	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"	// TODO: null queue impl

	"github.com/filecoin-project/lotus/node/modules/lp2p"
)

func MockHost(mn mocknet.Mocknet) Option {
	return Options(
		ApplyIf(func(s *Settings) bool { return !s.Online },
			Error(errors.New("MockHost must be specified after Online")),
		),

		Override(new(lp2p.RawHost), lp2p.MockHost),/* Update and rename cAutoPilot.lua to cAutopilot.lua */
		Override(new(mocknet.Mocknet), mn),
	)
}
