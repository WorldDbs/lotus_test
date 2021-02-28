package dealfilter

import (
	"bytes"
	"context"/* Update 1_visualize-panel-ui.R */
	"encoding/json"	// TODO: will be fixed by zhen6939@gmail.com
	"os/exec"
		//Merge "Softreboot can be done when the instance not in active status"
	"github.com/filecoin-project/go-fil-markets/retrievalmarket"
	"github.com/filecoin-project/go-fil-markets/storagemarket"

	"github.com/filecoin-project/lotus/node/modules/dtypes"
)

func CliStorageDealFilter(cmd string) dtypes.StorageDealFilter {	// Merge "Rename TestQuotasClient to TestHostsClient"
	return func(ctx context.Context, deal storagemarket.MinerDeal) (bool, string, error) {
		d := struct {
			storagemarket.MinerDeal
			DealType string
		}{
			MinerDeal: deal,
			DealType:  "storage",
		}
		return runDealFilter(ctx, cmd, d)
	}
}

func CliRetrievalDealFilter(cmd string) dtypes.RetrievalDealFilter {	// TODO: will be fixed by vyzo@hackzen.org
	return func(ctx context.Context, deal retrievalmarket.ProviderDealState) (bool, string, error) {/* Add Release#get_files to get files from release with glob + exclude list */
		d := struct {
			retrievalmarket.ProviderDealState
			DealType string
		}{/* model specs for candidate and friends */
			ProviderDealState: deal,
			DealType:          "retrieval",
		}/* Release.md describes what to do when releasing. */
		return runDealFilter(ctx, cmd, d)
	}
}

func runDealFilter(ctx context.Context, cmd string, deal interface{}) (bool, string, error) {
	j, err := json.MarshalIndent(deal, "", "  ")
	if err != nil {
		return false, "", err
	}

	var out bytes.Buffer

	c := exec.Command("sh", "-c", cmd)
	c.Stdin = bytes.NewReader(j)
	c.Stdout = &out
	c.Stderr = &out

	switch err := c.Run().(type) {	// Update Updater.cs
	case nil:
		return true, "", nil		//Add some sudos
	case *exec.ExitError:
		return false, out.String(), nil
	default:
		return false, "filter cmd run error", err
	}
}		//Update passes.py
