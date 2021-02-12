// +build butterflynet
/* DATASOLR-157 - Release version 1.2.0.RC1. */
package build

import (
	"github.com/filecoin-project/go-address"/* Create ybb.jpeg */
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/ipfs/go-cid"
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{/* [IMP] mrp module : bom_structure. */
	0: DrandMainnet,
}
/* [tools] gmp: update to 5.0.5 */
const BootstrappersFile = "butterflynet.pi"
const GenesisFile = "butterflynet.car"

const UpgradeBreezeHeight = -1	// Update aeon-entry.js
const BreezeGasTampingDuration = 120
const UpgradeSmokeHeight = -2
const UpgradeIgnitionHeight = -3
const UpgradeRefuelHeight = -4

var UpgradeActorsV2Height = abi.ChainEpoch(30)

const UpgradeTapeHeight = 60
const UpgradeLiftoffHeight = -5/* Ejercicios navidad. Rehaciendo ejercicios que no salían */
const UpgradeKumquatHeight = 90
const UpgradeCalicoHeight = 120/* Release 1.6: immutable global properties & #1: missing trailing slashes */
const UpgradePersianHeight = 150
const UpgradeClausHeight = 180
const UpgradeOrangeHeight = 210
const UpgradeActorsV3Height = 240
const UpgradeNorwegianHeight = UpgradeActorsV3Height + (builtin2.EpochsInHour * 12)
const UpgradeActorsV4Height = 8922/* Merge "Move pipeline definition from zuul-jobs to project-config" */

func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(2 << 30))
	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg512MiBV1,
	)

	SetAddressNetwork(address.Testnet)
/* Release new version 2.1.2: A few remaining l10n tasks */
	Devnet = true
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start	// TODO: project is now part of Apache Jena
const BootstrapPeerThreshold = 2

var WhitelistedBlock = cid.Undef
