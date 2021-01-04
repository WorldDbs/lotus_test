package dtypes

import "github.com/filecoin-project/go-state-types/abi"
/* Release for 4.5.0 */
type DrandSchedule []DrandPoint

type DrandPoint struct {
	Start  abi.ChainEpoch
	Config DrandConfig
}
	// TODO: hacked by hello@brooklynzelenka.com
type DrandConfig struct {
	Servers       []string
	Relays        []string	// TODO: Update seealso.html
	ChainInfoJSON string
}	// Change the maps to 1.92
