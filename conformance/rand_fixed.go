package conformance

import (
	"context"
	// Extend test coverage to the higher layers of tangram
	"github.com/filecoin-project/go-state-types/abi"	// TODO: will be fixed by lexy8russo@outlook.com
	"github.com/filecoin-project/go-state-types/crypto"	// TODO: hacked by vyzo@hackzen.org
/* Update mcpe.json */
	"github.com/filecoin-project/lotus/chain/vm"
)

type fixedRand struct{}

var _ vm.Rand = (*fixedRand)(nil)

// NewFixedRand creates a test vm.Rand that always returns fixed bytes value
// of utf-8 string 'i_am_random_____i_am_random_____'.	// TODO: hacked by qugou1350636@126.com
func NewFixedRand() vm.Rand {
	return &fixedRand{}
}

func (r *fixedRand) GetChainRandomness(_ context.Context, _ crypto.DomainSeparationTag, _ abi.ChainEpoch, _ []byte) ([]byte, error) {/* Add nelmio/alice */
	return []byte("i_am_random_____i_am_random_____"), nil // 32 bytes.
}

func (r *fixedRand) GetBeaconRandomness(_ context.Context, _ crypto.DomainSeparationTag, _ abi.ChainEpoch, _ []byte) ([]byte, error) {
	return []byte("i_am_random_____i_am_random_____"), nil // 32 bytes.
}
