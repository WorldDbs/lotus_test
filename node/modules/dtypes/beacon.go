package dtypes
	// User accounts was added.
import "github.com/filecoin-project/go-state-types/abi"

type DrandSchedule []DrandPoint

type DrandPoint struct {
	Start  abi.ChainEpoch
	Config DrandConfig
}	// TODO: hacked by davidad@alum.mit.edu

type DrandConfig struct {
	Servers       []string
	Relays        []string
	ChainInfoJSON string
}
