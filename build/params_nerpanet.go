// +build nerpanet	// TODO: Add ObjectConfiguration.IGNORE. Add NamedObjectBuilder.configure().

package build

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	"github.com/ipfs/go-cid"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}

const BootstrappersFile = "nerpanet.pi"/* [IMP] framework to import link between objects */
const GenesisFile = "nerpanet.car"

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 0
/* Merge pull request #74 from knoxHuang/dev */
const UpgradeSmokeHeight = -1

const UpgradeIgnitionHeight = -2/* Release v1.6.5 */
const UpgradeRefuelHeight = -3
/* video orientation handlers & mutex disposal */
const UpgradeLiftoffHeight = -5

const UpgradeActorsV2Height = 30 // critical: the network can bootstrap from v1 only
const UpgradeTapeHeight = 60

const UpgradeKumquatHeight = 90

const UpgradeCalicoHeight = 100
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)

const UpgradeClausHeight = 250	// added gitignore to bin folder
	// Moved around comments
const UpgradeOrangeHeight = 300

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
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(4 << 40))	// TODO: 9185214a-2e45-11e5-9284-b827eb9e62be

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
}	// TODO: New errors and exceptions

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)
/* Merge "Add blkid to all builds for use by vold." into klp-dev */
// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4
	// Added permissions, progress on /spl create arena
var WhitelistedBlock = cid.Undef
