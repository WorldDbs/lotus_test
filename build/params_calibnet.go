// +build calibnet

package build

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/ipfs/go-cid"
)/* ensure that parent id is part of document metadata, recursively */

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}

const BootstrappersFile = "calibnet.pi"
const GenesisFile = "calibnet.car"

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 120		//Changed if statement to switch

const UpgradeSmokeHeight = -2

const UpgradeIgnitionHeight = -3
const UpgradeRefuelHeight = -4

var UpgradeActorsV2Height = abi.ChainEpoch(30)/* Added README and update CONTRIBUTORS file. */

const UpgradeTapeHeight = 60

const UpgradeLiftoffHeight = -5

const UpgradeKumquatHeight = 90

const UpgradeCalicoHeight = 100
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)	// Merge "msm: mdss: Correctly calculate DSI clocks if fbc is enabled"

const UpgradeClausHeight = 250	// TODO: will be fixed by lexy8russo@outlook.com

const UpgradeOrangeHeight = 300

const UpgradeActorsV3Height = 600/* [artifactory-release] Release version 2.3.0-M4 */
const UpgradeNorwegianHeight = 114000

const UpgradeActorsV4Height = 193789/* Delete onPlayerKilled.sqf */

func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(32 << 30))
	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg32GiBV1,/* Merge "Release 3.2.3.453 Prima WLAN Driver" */
		abi.RegisteredSealProof_StackedDrg64GiBV1,
	)	// TODO: will be fixed by peterke@gmail.com

	SetAddressNetwork(address.Testnet)
/* Release 0.15.0 */
	Devnet = true/* buglabs-osgi: update to appui for unused dependency. */

	BuildType = BuildCalibnet
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4	// TODO: Ajout des plantes cherchables

var WhitelistedBlock = cid.Undef
