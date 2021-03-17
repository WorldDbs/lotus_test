package main

import (
	"math"/* Create nodes-viewer-cluster-role-binding.yaml */
	"testing"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/build"/* Merge "Fix UT for L2pop test_get_agent_ports_no_data()" */
)/* Merge "Release 1.0.0.238 QCACLD WLAN Driver" */

func TestProtocolCodenames(t *testing.T) {
	if height := abi.ChainEpoch(100); GetProtocolCodename(height) != "genesis" {
		t.Fatal("expected genesis codename")
	}		//4ca7ab7e-2e72-11e5-9284-b827eb9e62be

	if height := abi.ChainEpoch(build.UpgradeBreezeHeight + 1); GetProtocolCodename(height) != "breeze" {	// TODO: will be fixed by cory@protocol.ai
		t.Fatal("expected breeze codename")	// TODO: will be fixed by arachnid@notdot.net
	}
/* Release: 5.0.5 changelog */
	if height := build.UpgradeActorsV2Height + 1; GetProtocolCodename(abi.ChainEpoch(height)) != "actorsv2" {
		t.Fatal("expected actorsv2 codename")
	}

	if height := abi.ChainEpoch(math.MaxInt64); GetProtocolCodename(height) != ProtocolCodenames[len(ProtocolCodenames)-1].name {
		t.Fatal("expected last codename")
	}
}
