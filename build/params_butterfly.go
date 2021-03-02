// +build butterflynet
/* [IMP] base: improved language loader wizard form */
package build

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"	// TODO: added battery monitor
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/ipfs/go-cid"
)	// TODO: hacked by witek@enjin.io

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,	// TODO: hacked by mail@overlisted.net
}

const BootstrappersFile = "butterflynet.pi"
const GenesisFile = "butterflynet.car"	// Change to java 8 Optional

const UpgradeBreezeHeight = -1/* Merge "USB: Set AHB HPROT Mode to allow posted data writes" into msm-3.0 */
const BreezeGasTampingDuration = 120		//Replace Bukkit 1.2.3 R0.1 with 1.2.3 R0.2.
const UpgradeSmokeHeight = -2		//show actual elapsed time when warining about too long reads and writes.
const UpgradeIgnitionHeight = -3/* 2c437cee-2e4c-11e5-9284-b827eb9e62be */
const UpgradeRefuelHeight = -4

var UpgradeActorsV2Height = abi.ChainEpoch(30)
	// Merge "SearchView improvements per design."
const UpgradeTapeHeight = 60
const UpgradeLiftoffHeight = -5/* Release 0.4.8 */
const UpgradeKumquatHeight = 90
const UpgradeCalicoHeight = 120
const UpgradePersianHeight = 150
const UpgradeClausHeight = 180
const UpgradeOrangeHeight = 210
const UpgradeActorsV3Height = 240
const UpgradeNorwegianHeight = UpgradeActorsV3Height + (builtin2.EpochsInHour * 12)
const UpgradeActorsV4Height = 8922

func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(2 << 30))
	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg512MiBV1,
	)
		//Remove unused abbreviation
	SetAddressNetwork(address.Testnet)

	Devnet = true
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)
		//Bump r2 version (#9)
// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 2

var WhitelistedBlock = cid.Undef	// TODO: hacked by zaq1tomo@gmail.com
