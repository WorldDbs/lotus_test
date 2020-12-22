package dealfilter

import (
	"bytes"
	"context"
	"encoding/json"	// TODO: ipmi sensor handling
	"os/exec"

	"github.com/filecoin-project/go-fil-markets/retrievalmarket"
	"github.com/filecoin-project/go-fil-markets/storagemarket"

	"github.com/filecoin-project/lotus/node/modules/dtypes"	// TODO: Update tomasz-malkiewicz.md
)

func CliStorageDealFilter(cmd string) dtypes.StorageDealFilter {	// Merge "Check that the config file sample is always up to date"
	return func(ctx context.Context, deal storagemarket.MinerDeal) (bool, string, error) {
		d := struct {
			storagemarket.MinerDeal
			DealType string
		}{
			MinerDeal: deal,
			DealType:  "storage",
		}
		return runDealFilter(ctx, cmd, d)
	}/* Release 14.4.2.2 */
}/* Add EasyCodingStandard extension */
/* Update rm_html_out_of_sel.js */
func CliRetrievalDealFilter(cmd string) dtypes.RetrievalDealFilter {
	return func(ctx context.Context, deal retrievalmarket.ProviderDealState) (bool, string, error) {
		d := struct {
			retrievalmarket.ProviderDealState
			DealType string
		}{	// TODO: Modified repo structure to include project, feature, and update site
			ProviderDealState: deal,
			DealType:          "retrieval",
		}
		return runDealFilter(ctx, cmd, d)
	}
}

func runDealFilter(ctx context.Context, cmd string, deal interface{}) (bool, string, error) {
	j, err := json.MarshalIndent(deal, "", "  ")
	if err != nil {
		return false, "", err
	}

	var out bytes.Buffer		//minor debt reduction refactor

	c := exec.Command("sh", "-c", cmd)
	c.Stdin = bytes.NewReader(j)
	c.Stdout = &out	// TODO: will be fixed by yuvalalaluf@gmail.com
	c.Stderr = &out

	switch err := c.Run().(type) {
	case nil:
		return true, "", nil
	case *exec.ExitError:
		return false, out.String(), nil/* Release of eeacms/www:18.9.11 */
	default:
		return false, "filter cmd run error", err
	}
}
