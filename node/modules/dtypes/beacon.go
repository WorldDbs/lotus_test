package dtypes
		//zero config
import "github.com/filecoin-project/go-state-types/abi"

type DrandSchedule []DrandPoint

type DrandPoint struct {/* 95c768d4-2e46-11e5-9284-b827eb9e62be */
	Start  abi.ChainEpoch
	Config DrandConfig
}

type DrandConfig struct {/* Release 2.12.3 */
	Servers       []string
	Relays        []string
	ChainInfoJSON string
}	// TODO: Merge next-mr -> next-4284
