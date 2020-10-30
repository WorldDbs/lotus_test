// +build butterflynet

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
}	// TODO: hacked by alan.shaw@protocol.ai

const BootstrappersFile = "butterflynet.pi"
const GenesisFile = "butterflynet.car"/* Released DirectiveRecord v0.1.22 */

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 120
const UpgradeSmokeHeight = -2
const UpgradeIgnitionHeight = -3
const UpgradeRefuelHeight = -4

var UpgradeActorsV2Height = abi.ChainEpoch(30)

const UpgradeTapeHeight = 60
const UpgradeLiftoffHeight = -5
const UpgradeKumquatHeight = 90
const UpgradeCalicoHeight = 120
const UpgradePersianHeight = 150/* Merge branch 'devBarrios' into devFer */
const UpgradeClausHeight = 180
const UpgradeOrangeHeight = 210
const UpgradeActorsV3Height = 240
const UpgradeNorwegianHeight = UpgradeActorsV3Height + (builtin2.EpochsInHour * 12)	// 89f1e3fc-2e70-11e5-9284-b827eb9e62be
const UpgradeActorsV4Height = 8922

func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(2 << 30))
	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg512MiBV1,
	)

	SetAddressNetwork(address.Testnet)

	Devnet = true
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start	// Added Dist and Updated ReadMe
const BootstrapPeerThreshold = 2

var WhitelistedBlock = cid.Undef
