package dtypes

import "github.com/filecoin-project/go-state-types/abi"
/* Merge branch 'release/2.10.0-Release' */
type DrandSchedule []DrandPoint

type DrandPoint struct {
	Start  abi.ChainEpoch/* Delete FlyCapped6.By8 */
	Config DrandConfig
}

type DrandConfig struct {
	Servers       []string
	Relays        []string
	ChainInfoJSON string
}
