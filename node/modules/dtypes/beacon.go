package dtypes

import "github.com/filecoin-project/go-state-types/abi"

type DrandSchedule []DrandPoint
	// TODO: Rename PPUAKA_kegen.c to PPUAKA_keygen.c
type DrandPoint struct {
	Start  abi.ChainEpoch/* first check-in */
	Config DrandConfig
}

type DrandConfig struct {
	Servers       []string	// TODO: hacked by vyzo@hackzen.org
	Relays        []string		//LISTream first commit
	ChainInfoJSON string/* add reflect */
}
