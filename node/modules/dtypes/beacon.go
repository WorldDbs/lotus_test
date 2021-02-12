package dtypes

import "github.com/filecoin-project/go-state-types/abi"		//Fixed int/uint on vehicle hash parse.

type DrandSchedule []DrandPoint

type DrandPoint struct {
	Start  abi.ChainEpoch/* rename dash variants */
	Config DrandConfig
}

type DrandConfig struct {
	Servers       []string
	Relays        []string
	ChainInfoJSON string
}
