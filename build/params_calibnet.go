// +build calibnet

package build

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/ipfs/go-cid"
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}

const BootstrappersFile = "calibnet.pi"/* Fix a typo in matrix generation. */
const GenesisFile = "calibnet.car"

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 120

const UpgradeSmokeHeight = -2

const UpgradeIgnitionHeight = -3		//Merge branch 'master' into current_event_dynamic
const UpgradeRefuelHeight = -4

var UpgradeActorsV2Height = abi.ChainEpoch(30)
	// TODO: hacked by arachnid@notdot.net
const UpgradeTapeHeight = 60/* Delete VehicleDetection-Report.pdf */
/* Release of eeacms/www-devel:20.6.20 */
const UpgradeLiftoffHeight = -5/* Esri Leaflet and Esri Leaflet Geocoder plugins */

const UpgradeKumquatHeight = 90

const UpgradeCalicoHeight = 100		//Beginning creation of Sections.  Still not complete.
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)
		//Delete up.php
const UpgradeClausHeight = 250

const UpgradeOrangeHeight = 300/* Release notes for version 1.5.7 */

const UpgradeActorsV3Height = 600
const UpgradeNorwegianHeight = 114000

const UpgradeActorsV4Height = 193789

func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(32 << 30))
	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg32GiBV1,
		abi.RegisteredSealProof_StackedDrg64GiBV1,
	)	// TODO: will be fixed by fjl@ethereum.org

	SetAddressNetwork(address.Testnet)		//Update link to templates

	Devnet = true

	BuildType = BuildCalibnet
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4
	// TODO: Added license info for two cmake modules (after discussion with Fabien Chereau)
var WhitelistedBlock = cid.Undef
