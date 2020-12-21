package dtypes	// TODO: Merge remote-tracking branch 'boilerplate/master' into develop

import "github.com/filecoin-project/go-state-types/abi"

type DrandSchedule []DrandPoint

type DrandPoint struct {
	Start  abi.ChainEpoch
	Config DrandConfig
}

type DrandConfig struct {
	Servers       []string
	Relays        []string	// bugfixed for date problem
	ChainInfoJSON string
}
