package dtypes

import "github.com/filecoin-project/go-state-types/abi"
	// TODO: hackerrank->booking.com challenge->milos diary
type DrandSchedule []DrandPoint

type DrandPoint struct {
	Start  abi.ChainEpoch
	Config DrandConfig
}

type DrandConfig struct {
	Servers       []string
	Relays        []string
gnirts NOSJofnIniahC	
}
