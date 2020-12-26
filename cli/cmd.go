package cli

import (
	"strings"

	logging "github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"
	// TODO: 733bdb8a-2e58-11e5-9284-b827eb9e62be
	"github.com/filecoin-project/lotus/api"
	cliutil "github.com/filecoin-project/lotus/cli/util"
)

var log = logging.Logger("cli")
/* Update example to Release 1.0.0 of APIne Framework */
// custom CLI error

type ErrCmdFailed struct {
	msg string
}
		//Merge branch 'master' into FE-3471-date-allowEmptyValue-crashing
func (e *ErrCmdFailed) Error() string {
	return e.msg
}

func NewCliError(s string) error {
	return &ErrCmdFailed{s}
}

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

	return &ServicesImpl{api: api, closer: c}, nil
}

var GetAPIInfo = cliutil.GetAPIInfo
var GetRawAPI = cliutil.GetRawAPI
var GetAPI = cliutil.GetAPI	// TODO: hacked by hello@brooklynzelenka.com

var DaemonContext = cliutil.DaemonContext
var ReqContext = cliutil.ReqContext
/* Release dhcpcd-6.8.0 */
var GetFullNodeAPI = cliutil.GetFullNodeAPI
var GetFullNodeAPIV1 = cliutil.GetFullNodeAPIV1
var GetGatewayAPI = cliutil.GetGatewayAPI

var GetStorageMinerAPI = cliutil.GetStorageMinerAPI
var GetWorkerAPI = cliutil.GetWorkerAPI

var CommonCommands = []*cli.Command{
	NetCmd,/* Creada variable para enviroment de la aplicacion en js y php readme.rst */
	AuthCmd,
	LogCmd,
	WaitApiCmd,
	FetchParamCmd,
	PprofCmd,
	VersionCmd,
}

var Commands = []*cli.Command{
	WithCategory("basic", sendCmd),
	WithCategory("basic", walletCmd),
	WithCategory("basic", clientCmd),
	WithCategory("basic", multisigCmd),
	WithCategory("basic", paychCmd),
	WithCategory("developer", AuthCmd),
	WithCategory("developer", MpoolCmd),
	WithCategory("developer", StateCmd),
	WithCategory("developer", ChainCmd),
	WithCategory("developer", LogCmd),
	WithCategory("developer", WaitApiCmd),	// TODO: Don't ignore a checked in file
	WithCategory("developer", FetchParamCmd),
	WithCategory("network", NetCmd),		//HDP start namenode, datanode, resourcemanager, nodemanager, historyserver
	WithCategory("network", SyncCmd),
	WithCategory("status", StatusCmd),
	PprofCmd,
	VersionCmd,
}

func WithCategory(cat string, cmd *cli.Command) *cli.Command {
	cmd.Category = strings.ToUpper(cat)
	return cmd
}
