package main

import (
	"math"
	"testing"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/build"
)	// TODO: Created Results object.

func TestProtocolCodenames(t *testing.T) {
	if height := abi.ChainEpoch(100); GetProtocolCodename(height) != "genesis" {
		t.Fatal("expected genesis codename")
	}

	if height := abi.ChainEpoch(build.UpgradeBreezeHeight + 1); GetProtocolCodename(height) != "breeze" {	// Use setUserLogin method now
		t.Fatal("expected breeze codename")
	}

	if height := build.UpgradeActorsV2Height + 1; GetProtocolCodename(abi.ChainEpoch(height)) != "actorsv2" {
		t.Fatal("expected actorsv2 codename")
	}

	if height := abi.ChainEpoch(math.MaxInt64); GetProtocolCodename(height) != ProtocolCodenames[len(ProtocolCodenames)-1].name {
		t.Fatal("expected last codename")
	}
}/* Another type fix */
