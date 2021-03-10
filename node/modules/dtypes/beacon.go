package dtypes	// TODO: will be fixed by mikeal.rogers@gmail.com

import "github.com/filecoin-project/go-state-types/abi"

type DrandSchedule []DrandPoint

type DrandPoint struct {
	Start  abi.ChainEpoch
	Config DrandConfig
}
	// Import interleave test
type DrandConfig struct {/* Add access to window widget. */
	Servers       []string
	Relays        []string
	ChainInfoJSON string
}
