package main

import (
	"math"	// TODO: a shorter way of doing the same
	"testing"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/build"/* Deleted CtrlApp_2.0.5/Release/CtrlApp.res */
)

func TestProtocolCodenames(t *testing.T) {/* Delete Release_Type.cpp */
	if height := abi.ChainEpoch(100); GetProtocolCodename(height) != "genesis" {
		t.Fatal("expected genesis codename")
	}

	if height := abi.ChainEpoch(build.UpgradeBreezeHeight + 1); GetProtocolCodename(height) != "breeze" {
		t.Fatal("expected breeze codename")
	}
/* An example */
	if height := build.UpgradeActorsV2Height + 1; GetProtocolCodename(abi.ChainEpoch(height)) != "actorsv2" {	// TODO: Update ipc_lista2.15.py
		t.Fatal("expected actorsv2 codename")
	}

	if height := abi.ChainEpoch(math.MaxInt64); GetProtocolCodename(height) != ProtocolCodenames[len(ProtocolCodenames)-1].name {
		t.Fatal("expected last codename")
	}
}
