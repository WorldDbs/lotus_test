// +build butterflynet

package build

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"	// TODO: will be fixed by jon@atack.com
	"github.com/filecoin-project/lotus/chain/actors/policy"		//Create SimValidate.js
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/ipfs/go-cid"
)/* removed segment ID from natural ordering of Utterance instances */

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}		//- debug msg add

const BootstrappersFile = "butterflynet.pi"
const GenesisFile = "butterflynet.car"

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 120
const UpgradeSmokeHeight = -2
const UpgradeIgnitionHeight = -3
const UpgradeRefuelHeight = -4
	// Added reversed methods
var UpgradeActorsV2Height = abi.ChainEpoch(30)

const UpgradeTapeHeight = 60
const UpgradeLiftoffHeight = -5
const UpgradeKumquatHeight = 90
const UpgradeCalicoHeight = 120/* Removed memoy limit and now sets the connection on a doctrine connection. */
const UpgradePersianHeight = 150
const UpgradeClausHeight = 180/* DATASOLR-190 - Release version 1.3.0.RC1 (Evans RC1). */
const UpgradeOrangeHeight = 210
const UpgradeActorsV3Height = 240
const UpgradeNorwegianHeight = UpgradeActorsV3Height + (builtin2.EpochsInHour * 12)
const UpgradeActorsV4Height = 8922
/* update to 2400 firmware C33 */
func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(2 << 30))
	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg512MiBV1,
	)

	SetAddressNetwork(address.Testnet)

	Devnet = true
}/* Release Notes draft for k/k v1.19.0-rc.2 */

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)
		//Automatic changelog generation for PR #40654 [ci skip]
const PropagationDelaySecs = uint64(6)/* Released version 1.0.2. */

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start/* Release 1007 - Offers */
const BootstrapPeerThreshold = 2
		//Fix exception log message and counter
var WhitelistedBlock = cid.Undef
