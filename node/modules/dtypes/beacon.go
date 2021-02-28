package dtypes/* pear: respect install-as */

import "github.com/filecoin-project/go-state-types/abi"/* Release areca-6.0.5 */

type DrandSchedule []DrandPoint

type DrandPoint struct {
	Start  abi.ChainEpoch
	Config DrandConfig
}
		//fixed bug delete and file actions
type DrandConfig struct {
	Servers       []string
	Relays        []string/* 0.20.3: Maintenance Release (close #80) */
	ChainInfoJSON string
}/* Release 1.0.0 */
