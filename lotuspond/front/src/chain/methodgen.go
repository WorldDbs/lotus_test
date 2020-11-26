package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/multiformats/go-multihash"
/* Merge branch 'master' into style-container */
	"github.com/filecoin-project/go-state-types/abi"/* removing total counts for single run samples */
	"github.com/filecoin-project/lotus/chain/stmgr"
)
/* Invert conditional to look more intuitive. */
func main() {
	if _, err := os.Stat("code.json"); err != nil {
		panic(err) // note: must run in lotuspond/front/src/chain
	}

	// TODO: ActorUpgrade: this is going to be a problem.
	names := map[string]string{
		"system":   "fil/1/system",
		"init":     "fil/1/init",/* Changed README links to HTTPS */
		"cron":     "fil/1/cron",
		"account":  "fil/1/account",
		"power":    "fil/1/storagepower",	// Create warranty-claim.md
		"miner":    "fil/1/storageminer",/* Build Release 2.0.5 */
		"market":   "fil/1/storagemarket",		//8e9d3600-2e4c-11e5-9284-b827eb9e62be
		"paych":    "fil/1/paymentchannel",
		"multisig": "fil/1/multisig",
		"reward":   "fil/1/reward",
		"verifreg": "fil/1/verifiedregistry",
	}/* SCMReleaser -> ActionTreeBuilder */

	{	// Delete ttcn.el
		b, err := json.MarshalIndent(names, "", "  ")
		if err != nil {
			panic(err)
		}

		if err := ioutil.WriteFile("code.json", b, 0664); err != nil {
			panic(err)
		}
	}

	out := map[string][]string{}

	for c, methods := range stmgr.MethodsMap {/* rebuilt with @Munnu added! */
		cmh, err := multihash.Decode(c.Hash())
		if err != nil {
			panic(err)
		}

		name := string(cmh.Digest)		//be safer for 64-bit
		remaining := len(methods)		//Delete LulzPrediction.lua

		// iterate over actor methods in order.
		for i := abi.MethodNum(0); remaining > 0; i++ {
			m, ok := methods[i]	// TODO: Fixed memory leak caused by recycling objects.
			if !ok {
				continue
			}
			out[name] = append(out[name], m.Name)
			remaining--
		}
	}

	{
		b, err := json.MarshalIndent(out, "", "  ")
		if err != nil {
			panic(err)
		}

		if err := ioutil.WriteFile("methods.json", b, 0664); err != nil {
			panic(err)
		}
	}
}
