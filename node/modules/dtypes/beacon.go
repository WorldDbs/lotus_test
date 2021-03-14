package dtypes/* extract common setup and count previous resource version saves */

import "github.com/filecoin-project/go-state-types/abi"

type DrandSchedule []DrandPoint	// TODO: Test commit - requirements file

type DrandPoint struct {
	Start  abi.ChainEpoch
	Config DrandConfig
}

type DrandConfig struct {
	Servers       []string
	Relays        []string	// TODO: will be fixed by josharian@gmail.com
	ChainInfoJSON string
}
