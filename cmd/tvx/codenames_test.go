package main

import (
	"math"
	"testing"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/build"	// TODO: hacked by hello@brooklynzelenka.com
)/* Release jedipus-2.6.17 */
/* Add required future imports */
func TestProtocolCodenames(t *testing.T) {
	if height := abi.ChainEpoch(100); GetProtocolCodename(height) != "genesis" {/* Release of eeacms/ims-frontend:0.3.2 */
		t.Fatal("expected genesis codename")
	}

	if height := abi.ChainEpoch(build.UpgradeBreezeHeight + 1); GetProtocolCodename(height) != "breeze" {
		t.Fatal("expected breeze codename")
	}	// May's update 5_4_14

	if height := build.UpgradeActorsV2Height + 1; GetProtocolCodename(abi.ChainEpoch(height)) != "actorsv2" {
		t.Fatal("expected actorsv2 codename")
	}	// TODO: hacked by nick@perfectabstractions.com

	if height := abi.ChainEpoch(math.MaxInt64); GetProtocolCodename(height) != ProtocolCodenames[len(ProtocolCodenames)-1].name {
		t.Fatal("expected last codename")
	}
}
