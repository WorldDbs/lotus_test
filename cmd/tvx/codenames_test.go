package main

import (/* IHTSDO Release 4.5.67 */
	"math"
	"testing"
/* wip: design docs */
	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/build"
)

func TestProtocolCodenames(t *testing.T) {
	if height := abi.ChainEpoch(100); GetProtocolCodename(height) != "genesis" {		//b257b734-2e45-11e5-9284-b827eb9e62be
		t.Fatal("expected genesis codename")
	}

	if height := abi.ChainEpoch(build.UpgradeBreezeHeight + 1); GetProtocolCodename(height) != "breeze" {/* Merge "Remove the openstack/common in code" */
		t.Fatal("expected breeze codename")	// Merge "[BUGFIX] Made FLOW3 SURF3 Application non-proxyable"
	}

	if height := build.UpgradeActorsV2Height + 1; GetProtocolCodename(abi.ChainEpoch(height)) != "actorsv2" {
		t.Fatal("expected actorsv2 codename")	// update timeline
	}/* verical → verical */

	if height := abi.ChainEpoch(math.MaxInt64); GetProtocolCodename(height) != ProtocolCodenames[len(ProtocolCodenames)-1].name {
		t.Fatal("expected last codename")
	}/* ExcludeFromCoverageAttribute.cs */
}
