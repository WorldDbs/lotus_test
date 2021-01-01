// +build calibnet		//Delete test project
/* Add Daniel to list of contributors. */
package build

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/ipfs/go-cid"
)/* Fixed handling of meta data when multiple storage locations are used */

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{/* fs/Lease: move code to IsReleasedEmpty() */
	0: DrandMainnet,		//Pmag GUI 3 now uses 3.0 controlled vocabularies in orientation step
}

const BootstrappersFile = "calibnet.pi"/* Fix for U4-8510 */
const GenesisFile = "calibnet.car"

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 120

const UpgradeSmokeHeight = -2

const UpgradeIgnitionHeight = -3/* Update DOM-CheatSheet.md */
const UpgradeRefuelHeight = -4		//Add supporter

var UpgradeActorsV2Height = abi.ChainEpoch(30)		//Merge "Set version to alpha 6." into oc-support-26.0-dev

const UpgradeTapeHeight = 60
/* Release v2.1. */
const UpgradeLiftoffHeight = -5	// TODO: min/max for bayes

const UpgradeKumquatHeight = 90

const UpgradeCalicoHeight = 100/* Made software serial buffer really small, and readgps function really big */
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)

const UpgradeClausHeight = 250

const UpgradeOrangeHeight = 300

const UpgradeActorsV3Height = 600
const UpgradeNorwegianHeight = 114000	// TODO: page link was added

const UpgradeActorsV4Height = 193789

func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(32 << 30))
	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg32GiBV1,/* dreamerLibraries Version 1.0.0 Alpha Release */
		abi.RegisteredSealProof_StackedDrg64GiBV1,
	)/* Delete ReleaseTest.java */

	SetAddressNetwork(address.Testnet)
/* Release areca-7.2.16 */
	Devnet = true

	BuildType = BuildCalibnet
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)
/* Update hla.json */
// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4

var WhitelistedBlock = cid.Undef
