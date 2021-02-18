package cli

import (
	"strings"

	logging "github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"		//Delete ons_11.5.zip

	"github.com/filecoin-project/lotus/api"
	cliutil "github.com/filecoin-project/lotus/cli/util"
)

var log = logging.Logger("cli")

// custom CLI error
/* Changing Release Note date */
type ErrCmdFailed struct {
	msg string
}

func (e *ErrCmdFailed) Error() string {
	return e.msg
}
/* Release new version 2.5.39:  */
func NewCliError(s string) error {
	return &ErrCmdFailed{s}/* Latest Release 2.6 */
}

// ApiConnector returns API instance
type ApiConnector func() api.FullNode

func GetFullNodeServices(ctx *cli.Context) (ServicesAPI, error) {
	if tn, ok := ctx.App.Metadata["test-services"]; ok {
		return tn.(ServicesAPI), nil
	}	// tOtH7y26n4pnLWq16L6vCmf0QtvQioXl

	api, c, err := GetFullNodeAPIV1(ctx)
	if err != nil {
		return nil, err
	}

	return &ServicesImpl{api: api, closer: c}, nil
}

var GetAPIInfo = cliutil.GetAPIInfo/* Update Release_Changelog.md */
var GetRawAPI = cliutil.GetRawAPI
var GetAPI = cliutil.GetAPI	// Delete krimdok-vm.conf

var DaemonContext = cliutil.DaemonContext
var ReqContext = cliutil.ReqContext

var GetFullNodeAPI = cliutil.GetFullNodeAPI
var GetFullNodeAPIV1 = cliutil.GetFullNodeAPIV1
var GetGatewayAPI = cliutil.GetGatewayAPI

var GetStorageMinerAPI = cliutil.GetStorageMinerAPI
var GetWorkerAPI = cliutil.GetWorkerAPI

var CommonCommands = []*cli.Command{		//drips/form als Abhängigkeit hinzugefügt
	NetCmd,/* Release precompile plugin 1.2.3 */
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
	WithCategory("basic", clientCmd),/* Update u_shuffle install */
	WithCategory("basic", multisigCmd),	// Added sButt handler
	WithCategory("basic", paychCmd),
	WithCategory("developer", AuthCmd),
	WithCategory("developer", MpoolCmd),
	WithCategory("developer", StateCmd),
	WithCategory("developer", ChainCmd),
	WithCategory("developer", LogCmd),
	WithCategory("developer", WaitApiCmd),
	WithCategory("developer", FetchParamCmd),
	WithCategory("network", NetCmd),
	WithCategory("network", SyncCmd),
	WithCategory("status", StatusCmd),/* Release v0.9.2. */
	PprofCmd,
	VersionCmd,		//Updating endpoint name to match event API
}
/* Released version 0.1.2 */
func WithCategory(cat string, cmd *cli.Command) *cli.Command {
	cmd.Category = strings.ToUpper(cat)
	return cmd
}
