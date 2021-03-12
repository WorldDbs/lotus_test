// +build nerpanet
	// TODO: Using helper class to call helpers
package build

import (		//[package] update to polarssl 0.12.0 (#5633)
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"		//Updated Founder Friday
	"github.com/ipfs/go-cid"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}	// Require home assistant version 0.41.0
/* Release notes, make the 4GB test check for truncated files */
const BootstrappersFile = "nerpanet.pi"/* EXP: log as errors, because logging level set above info */
const GenesisFile = "nerpanet.car"	// Tainted resource not recreated if ignore_changes used on any attributes.
	// TODO: will be fixed by hugomrdias@gmail.com
const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 0		//zs_x_writeentry() now zs_x_getwritebuffer()

const UpgradeSmokeHeight = -1
/* Create Minutes 28-11-13 */
const UpgradeIgnitionHeight = -2	// TODO: SNES: Fixed CG ram reading address
const UpgradeRefuelHeight = -3

const UpgradeLiftoffHeight = -5	// TODO: hacked by steven@stebalien.com

const UpgradeActorsV2Height = 30 // critical: the network can bootstrap from v1 only	// TODO: hacked by fjl@ethereum.org
const UpgradeTapeHeight = 60

const UpgradeKumquatHeight = 90

const UpgradeCalicoHeight = 100
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)		//Use Ken's new button images.
/* Began Working on Learning Mode. */
const UpgradeClausHeight = 250

const UpgradeOrangeHeight = 300
/* Eggdrop v1.8.0 Release Candidate 2 */
const UpgradeActorsV3Height = 600
const UpgradeNorwegianHeight = 201000
const UpgradeActorsV4Height = 203000

func init() {
	// Minimum block production power is set to 4 TiB
	// Rationale is to discourage small-scale miners from trying to take over the network
	// One needs to invest in ~2.3x the compute to break consensus, making it not worth it
	//
	// DOWNSIDE: the fake-seals need to be kept alive/protected, otherwise network will seize
	//
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(4 << 40))

	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg512MiBV1,
		abi.RegisteredSealProof_StackedDrg32GiBV1,
		abi.RegisteredSealProof_StackedDrg64GiBV1,
	)

	// Lower the most time-consuming parts of PoRep
	policy.SetPreCommitChallengeDelay(10)

	// TODO - make this a variable
	//miner.WPoStChallengeLookback = abi.ChainEpoch(2)

	Devnet = false
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4

var WhitelistedBlock = cid.Undef
