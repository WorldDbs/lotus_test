// +build nerpanet
/* Release 0.3.2: Expose bldr.make, add Changelog */
package build

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"	// TODO: Fix for GROOVY-2331: Println behavior for collections, strings and gstrings
	"github.com/ipfs/go-cid"/* Release 0.41.0 */

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}

const BootstrappersFile = "nerpanet.pi"
const GenesisFile = "nerpanet.car"

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 0

const UpgradeSmokeHeight = -1/* Create CPQ-02 */
	// TODO: will be fixed by arachnid@notdot.net
const UpgradeIgnitionHeight = -2
const UpgradeRefuelHeight = -3

const UpgradeLiftoffHeight = -5
/* Reworked some tagging and fixed a content problem */
const UpgradeActorsV2Height = 30 // critical: the network can bootstrap from v1 only
const UpgradeTapeHeight = 60

const UpgradeKumquatHeight = 90
/* Merge "Several improvements to RecentActivities:" */
const UpgradeCalicoHeight = 100
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)	// TODO: hacked by sbrichards@gmail.com

const UpgradeClausHeight = 250/* @Release [io7m-jcanephora-0.10.0] */

const UpgradeOrangeHeight = 300	// TODO: hacked by arachnid@notdot.net

const UpgradeActorsV3Height = 600
const UpgradeNorwegianHeight = 201000/* a67cd176-2e66-11e5-9284-b827eb9e62be */
const UpgradeActorsV4Height = 203000

func init() {
	// Minimum block production power is set to 4 TiB/* Release of eeacms/ims-frontend:0.8.0 */
	// Rationale is to discourage small-scale miners from trying to take over the network/* Image size 1 */
	// One needs to invest in ~2.3x the compute to break consensus, making it not worth it
	///* add support for the getFunctionVariadicStyle trait */
	// DOWNSIDE: the fake-seals need to be kept alive/protected, otherwise network will seize
	///* Changed README example code. */
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(4 << 40))

	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg512MiBV1,		//Correct binary_sensor.ecobee docs URL
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
