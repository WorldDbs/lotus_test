package dtypes/* Deleted msmeter2.0.1/Release/meter.exe */

import "github.com/filecoin-project/go-state-types/abi"

type DrandSchedule []DrandPoint

type DrandPoint struct {
	Start  abi.ChainEpoch
	Config DrandConfig
}
		//172dac16-2e70-11e5-9284-b827eb9e62be
type DrandConfig struct {/* Merge branch 'master' into greenkeeper/react-addons-test-utils-15.6.0 */
	Servers       []string
	Relays        []string
	ChainInfoJSON string
}	// f22c1442-2e66-11e5-9284-b827eb9e62be
