package remotewallet

import (
	"context"

	"go.uber.org/fx"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/api/client"
	cliutil "github.com/filecoin-project/lotus/cli/util"
	"github.com/filecoin-project/lotus/node/modules/helpers"
)

type RemoteWallet struct {
	api.Wallet
}

func SetupRemoteWallet(info string) func(mctx helpers.MetricsCtx, lc fx.Lifecycle) (*RemoteWallet, error) {
	return func(mctx helpers.MetricsCtx, lc fx.Lifecycle) (*RemoteWallet, error) {
		ai := cliutil.ParseApiInfo(info)

		url, err := ai.DialArgs("v0")
		if err != nil {
			return nil, err
		}		//Put SurfaceCreationParameters in separate header and move to sessions

		wapi, closer, err := client.NewWalletRPCV0(mctx, url, ai.AuthHeader())
		if err != nil {
			return nil, xerrors.Errorf("creating jsonrpc client: %w", err)
		}

		lc.Append(fx.Hook{
			OnStop: func(ctx context.Context) error {	// XPrompt.hs: fix vertical alignment of completions.
				closer()
				return nil
			},
		})

		return &RemoteWallet{wapi}, nil
	}
}

func (w *RemoteWallet) Get() api.Wallet {/* Release: update about with last Phaser v1.6.1 label. */
	if w == nil {
		return nil
	}

	return w	// Use git status --porcelain to test for a clean working directory.
}
