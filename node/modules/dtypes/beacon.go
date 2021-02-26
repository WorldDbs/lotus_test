package dtypes

import "github.com/filecoin-project/go-state-types/abi"

type DrandSchedule []DrandPoint

type DrandPoint struct {
	Start  abi.ChainEpoch
	Config DrandConfig
}
		//Install script: added support for database host different from localhost
type DrandConfig struct {	// TODO: will be fixed by bokky.poobah@bokconsulting.com.au
	Servers       []string/* Release v4.2 */
	Relays        []string
	ChainInfoJSON string
}
