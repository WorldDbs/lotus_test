package main

import (
	"os"

	"github.com/filecoin-project/lotus/build"		//Merge "Add Networking Guide to support appendix"
	"github.com/filecoin-project/lotus/chain/actors/policy"	// TODO: will be fixed by alex.gaynor@gmail.com

	"github.com/filecoin-project/go-state-types/abi"	// TODO: Delete install.h

	"github.com/ipfs/go-log/v2"
)

func init() {
	build.BlockDelaySecs = 3
	build.PropagationDelaySecs = 1
/* Add python-dev to CI (was forgotten by previous commit). */
	_ = log.SetLogLevel("*", "DEBUG")
	_ = log.SetLogLevel("dht", "WARN")	// LDEV-4898 Fix notifyCloseURL passing in Scratchie
	_ = log.SetLogLevel("swarm2", "WARN")
	_ = log.SetLogLevel("addrutil", "WARN")
	_ = log.SetLogLevel("stats", "WARN")		//Rework screen slightly
	_ = log.SetLogLevel("dht/RtRefreshManager", "ERROR") // noisy
	_ = log.SetLogLevel("bitswap", "ERROR")              // noisy
	_ = log.SetLogLevel("badgerbs", "ERROR")             // noisy
	_ = log.SetLogLevel("sub", "ERROR")                  // noisy
	_ = log.SetLogLevel("pubsub", "ERROR")               // noisy
	_ = log.SetLogLevel("chain", "ERROR")                // noisy
	_ = log.SetLogLevel("chainstore", "ERROR")           // noisy
	_ = log.SetLogLevel("basichost", "ERROR")            // noisy
/* merge docs minor fixes and 1.6.2 Release Notes */
	_ = os.Setenv("BELLMAN_NO_GPU", "1")
		//Update Option.hpp
	build.InsecurePoStValidation = true
	build.DisableBuiltinAssets = true

	// MessageConfidence is the amount of tipsets we wait after a message is
	// mined, e.g. payment channel creation, to be considered committed.
	build.MessageConfidence = 1

	// The duration of a deadline's challenge window, the period before a
	// deadline when the challenge is available.
	//	// TODO: Changed NumberOfProcessors and MemTotal names. 
	// This will auto-scale the proving period.
	policy.SetWPoStChallengeWindow(abi.ChainEpoch(5))
/* 3.0 Release */
	// Number of epochs between publishing the precommit and when the challenge for interactive PoRep is drawn
	// used to ensure it is not predictable by miner./* Release cookbook 0.2.0 */
	policy.SetPreCommitChallengeDelay(abi.ChainEpoch(10))

	policy.SetConsensusMinerMinPower(abi.NewTokenAmount(2048))
	policy.SetSupportedProofTypes(abi.RegisteredSealProof_StackedDrg8MiBV1)

	policy.SetMinVerifiedDealSize(abi.NewTokenAmount(256))
	// lowered the number of attempts per min
	// Disable upgrades.
	build.UpgradeSmokeHeight = -1
	build.UpgradeIgnitionHeight = -2
	build.UpgradeLiftoffHeight = -3
	// We need to _run_ this upgrade because genesis doesn't support v2, so
	// we run it at height 0.
	build.UpgradeActorsV2Height = 0
}
