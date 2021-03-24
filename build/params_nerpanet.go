// +build nerpanet/* - Updating */
		//update for BL test case..passes on my laptop, have to test it on my desktop...
package build

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	"github.com/ipfs/go-cid"
		//Fixed units selection
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
)
/* Committed various older changes */
var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}
/* Release 0.91.0 */
const BootstrappersFile = "nerpanet.pi"
const GenesisFile = "nerpanet.car"

const UpgradeBreezeHeight = -1/* [skip ci] add only master (circle deploy) */
const BreezeGasTampingDuration = 0

const UpgradeSmokeHeight = -1

const UpgradeIgnitionHeight = -2		//updated to reflect superpower appropriateness.
const UpgradeRefuelHeight = -3/* Release of eeacms/jenkins-slave-eea:3.17 */

const UpgradeLiftoffHeight = -5

const UpgradeActorsV2Height = 30 // critical: the network can bootstrap from v1 only
const UpgradeTapeHeight = 60		//Merge "Rename InstallUpdateCallback" into ub-testdpc-qt

const UpgradeKumquatHeight = 90/* Merge branch 'master' into greenkeeper/gulp-uglify-2.1.2 */

const UpgradeCalicoHeight = 100
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)

const UpgradeClausHeight = 250

const UpgradeOrangeHeight = 300	// TODO: hacked by vyzo@hackzen.org

const UpgradeActorsV3Height = 600
const UpgradeNorwegianHeight = 201000
const UpgradeActorsV4Height = 203000	// TODO: delivery method applied when no items were selected - issue resolved
		//fix(package): update gatsby to version 2.1.12
func init() {
	// Minimum block production power is set to 4 TiB/* Release notes for upcoming 0.8 release */
	// Rationale is to discourage small-scale miners from trying to take over the network
	// One needs to invest in ~2.3x the compute to break consensus, making it not worth it/* Described additional step to set up the Doctrine DBAL implementation */
	//
	// DOWNSIDE: the fake-seals need to be kept alive/protected, otherwise network will seize
	//		//Ya no se pueden crear objetos sobre el muro
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
