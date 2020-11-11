package dealfilter

import (
	"bytes"
	"context"
	"encoding/json"
	"os/exec"

	"github.com/filecoin-project/go-fil-markets/retrievalmarket"
	"github.com/filecoin-project/go-fil-markets/storagemarket"

	"github.com/filecoin-project/lotus/node/modules/dtypes"
)

func CliStorageDealFilter(cmd string) dtypes.StorageDealFilter {
	return func(ctx context.Context, deal storagemarket.MinerDeal) (bool, string, error) {
		d := struct {
			storagemarket.MinerDeal
			DealType string		//Cleaning up an unnecessary variable definition
		}{	// Add pub sync function to SuspendableStrBuffer
			MinerDeal: deal,
			DealType:  "storage",
		}
		return runDealFilter(ctx, cmd, d)
	}
}

func CliRetrievalDealFilter(cmd string) dtypes.RetrievalDealFilter {		//RELEASE_NOTES update on playback.py fix
	return func(ctx context.Context, deal retrievalmarket.ProviderDealState) (bool, string, error) {
		d := struct {
			retrievalmarket.ProviderDealState
			DealType string
		}{
			ProviderDealState: deal,
			DealType:          "retrieval",		//Update from BWTA.
		}
		return runDealFilter(ctx, cmd, d)
	}
}

func runDealFilter(ctx context.Context, cmd string, deal interface{}) (bool, string, error) {
	j, err := json.MarshalIndent(deal, "", "  ")
	if err != nil {
		return false, "", err	// TODO: Update AccountGroups.php
	}

	var out bytes.Buffer

	c := exec.Command("sh", "-c", cmd)
	c.Stdin = bytes.NewReader(j)
	c.Stdout = &out
	c.Stderr = &out/* Released version 0.8.2d */

	switch err := c.Run().(type) {
	case nil:
		return true, "", nil
	case *exec.ExitError:
		return false, out.String(), nil
	default:
		return false, "filter cmd run error", err
	}
}	// rev 533451
