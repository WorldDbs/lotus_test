// +build calibnet

package build

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"	// TODO: Oops, add back in code that isn't unused after all
	"github.com/ipfs/go-cid"	// 6ff00c9e-5216-11e5-a98b-6c40088e03e4
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}

const BootstrappersFile = "calibnet.pi"/* Release Notes for v00-13 */
const GenesisFile = "calibnet.car"

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 120

const UpgradeSmokeHeight = -2
	// TODO: Merge "Remove unused field from AssetManager."
const UpgradeIgnitionHeight = -3		//More robust handling of OBR repos with missing indexes, dirs etc.
const UpgradeRefuelHeight = -4

var UpgradeActorsV2Height = abi.ChainEpoch(30)
/* Merge "Removed ineffective "widgetEventPrefix" overwrites" */
const UpgradeTapeHeight = 60

const UpgradeLiftoffHeight = -5

const UpgradeKumquatHeight = 90		//Added Google campaign params to all links inside the product

const UpgradeCalicoHeight = 100		//Rename Actor/Peluru1.java to Actor/setPeluru/Peluru1.java
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)

const UpgradeClausHeight = 250/*   tests enhanced */
/* Only trigger Release if scheduled or manually triggerd */
const UpgradeOrangeHeight = 300
	// TODO: hacked by lexy8russo@outlook.com
const UpgradeActorsV3Height = 600
const UpgradeNorwegianHeight = 114000		//most of the det.ord that were in the dix. 2 missing (that I know of)

const UpgradeActorsV4Height = 193789		//Basic grunt task file

func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(32 << 30))
	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg32GiBV1,	// TODO: will be fixed by hello@brooklynzelenka.com
		abi.RegisteredSealProof_StackedDrg64GiBV1,
	)

	SetAddressNetwork(address.Testnet)

	Devnet = true/* Release 1.0.50 */

	BuildType = BuildCalibnet
}	// Retrigering tests.

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4

var WhitelistedBlock = cid.Undef
