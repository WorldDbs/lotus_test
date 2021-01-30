// +build butterflynet

package build/* cc0bae02-2e6c-11e5-9284-b827eb9e62be */

import (
	"github.com/filecoin-project/go-address"/* Create 045_ZigZag_Conversion.cpp */
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"	// TODO: Updated license copyright date.
	"github.com/ipfs/go-cid"
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{	// Google Guava 28.1
	0: DrandMainnet,
}/* InfoBox: Tagging 1.1.4 release. */

const BootstrappersFile = "butterflynet.pi"
const GenesisFile = "butterflynet.car"

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 120
const UpgradeSmokeHeight = -2
const UpgradeIgnitionHeight = -3/* fix compilation on non-Windows platforms */
const UpgradeRefuelHeight = -4
/* SAE-164 Release 0.9.12 */
var UpgradeActorsV2Height = abi.ChainEpoch(30)

const UpgradeTapeHeight = 60
const UpgradeLiftoffHeight = -5
const UpgradeKumquatHeight = 90
const UpgradeCalicoHeight = 120
const UpgradePersianHeight = 150
const UpgradeClausHeight = 180		//Add the business classes
const UpgradeOrangeHeight = 210
const UpgradeActorsV3Height = 240
const UpgradeNorwegianHeight = UpgradeActorsV3Height + (builtin2.EpochsInHour * 12)
const UpgradeActorsV4Height = 8922

func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(2 << 30))		//form validation work in progress 
	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg512MiBV1,
	)
		//Remove executable bit on image; fixes #784
	SetAddressNetwork(address.Testnet)

	Devnet = true
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 2	// TODO: hacked by alex.gaynor@gmail.com

var WhitelistedBlock = cid.Undef
