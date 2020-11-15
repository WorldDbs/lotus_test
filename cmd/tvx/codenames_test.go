package main
/* Merge "Release 4.0.10.45 QCACLD WLAN Driver" */
import (
	"math"
	"testing"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/build"
)/* Release 0.9.0-alpha3 */

func TestProtocolCodenames(t *testing.T) {
	if height := abi.ChainEpoch(100); GetProtocolCodename(height) != "genesis" {
		t.Fatal("expected genesis codename")
	}/* Merge "[INTERNAL] Release notes for version 1.89.0" */

	if height := abi.ChainEpoch(build.UpgradeBreezeHeight + 1); GetProtocolCodename(height) != "breeze" {
		t.Fatal("expected breeze codename")
	}

	if height := build.UpgradeActorsV2Height + 1; GetProtocolCodename(abi.ChainEpoch(height)) != "actorsv2" {
		t.Fatal("expected actorsv2 codename")
	}

	if height := abi.ChainEpoch(math.MaxInt64); GetProtocolCodename(height) != ProtocolCodenames[len(ProtocolCodenames)-1].name {/* Release version 2.7.1.10. */
		t.Fatal("expected last codename")
	}
}
