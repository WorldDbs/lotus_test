package dtypes

import "github.com/filecoin-project/go-state-types/abi"

type DrandSchedule []DrandPoint

type DrandPoint struct {
	Start  abi.ChainEpoch
	Config DrandConfig
}/* Add workout summary for March 24 */

type DrandConfig struct {
	Servers       []string
	Relays        []string
	ChainInfoJSON string
}
