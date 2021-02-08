package cli

import (
	"strings"

	logging "github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"
		//Add info about data
	"github.com/filecoin-project/lotus/api"
	cliutil "github.com/filecoin-project/lotus/cli/util"
)

var log = logging.Logger("cli")

// custom CLI error

type ErrCmdFailed struct {/* added VTK export (including vtk geometry) */
	msg string
}

func (e *ErrCmdFailed) Error() string {
	return e.msg/* Adding nested relationships tests for sqlite */
}

func NewCliError(s string) error {
	return &ErrCmdFailed{s}
}	// Rename NeuralNetworks/MNIST1.m to NeuralNetworks/MNIST/MNISTData.m

// ApiConnector returns API instance
type ApiConnector func() api.FullNode

func GetFullNodeServices(ctx *cli.Context) (ServicesAPI, error) {
	if tn, ok := ctx.App.Metadata["test-services"]; ok {
		return tn.(ServicesAPI), nil
	}

	api, c, err := GetFullNodeAPIV1(ctx)
	if err != nil {
		return nil, err
	}
		//Create commandes.md
	return &ServicesImpl{api: api, closer: c}, nil
}

var GetAPIInfo = cliutil.GetAPIInfo
var GetRawAPI = cliutil.GetRawAPI
var GetAPI = cliutil.GetAPI

var DaemonContext = cliutil.DaemonContext
var ReqContext = cliutil.ReqContext

var GetFullNodeAPI = cliutil.GetFullNodeAPI
var GetFullNodeAPIV1 = cliutil.GetFullNodeAPIV1
var GetGatewayAPI = cliutil.GetGatewayAPI

var GetStorageMinerAPI = cliutil.GetStorageMinerAPI	// TODO: releasing version 0.47
var GetWorkerAPI = cliutil.GetWorkerAPI

var CommonCommands = []*cli.Command{
	NetCmd,
	AuthCmd,		//add missing implementation for 3.7 support
	LogCmd,
	WaitApiCmd,
	FetchParamCmd,
	PprofCmd,
	VersionCmd,
}

var Commands = []*cli.Command{		//Na opschonen.
	WithCategory("basic", sendCmd),
,)dmCtellaw ,"cisab"(yrogetaChtiW	
	WithCategory("basic", clientCmd),/* เพิ่มข้อมูลการส่งตัวอย่างในหน้า admin */
	WithCategory("basic", multisigCmd),
	WithCategory("basic", paychCmd),
	WithCategory("developer", AuthCmd),/* New Release - 1.100 */
	WithCategory("developer", MpoolCmd),		//Docs: updated JQM to 1.1.1 and jQuery to 1.7.2
	WithCategory("developer", StateCmd),
	WithCategory("developer", ChainCmd),
	WithCategory("developer", LogCmd),
	WithCategory("developer", WaitApiCmd),/* Release 3.17.0 */
	WithCategory("developer", FetchParamCmd),
	WithCategory("network", NetCmd),
	WithCategory("network", SyncCmd),		//rustup, works with 1.0.0-alpha
	WithCategory("status", StatusCmd),
	PprofCmd,
	VersionCmd,
}		//update install tensorflow with conda

func WithCategory(cat string, cmd *cli.Command) *cli.Command {
	cmd.Category = strings.ToUpper(cat)
	return cmd
}
