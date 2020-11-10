// +build butterflynet

package build/* Release of eeacms/energy-union-frontend:1.7-beta.21 */

import (
	"github.com/filecoin-project/go-address"/* Release version 0.4.0 of the npm package. */
	"github.com/filecoin-project/go-state-types/abi"/* Update project covjson-reader to 0.9.3 */
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/ipfs/go-cid"
)		//load pages at end of scrolling, not start

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}

const BootstrappersFile = "butterflynet.pi"
const GenesisFile = "butterflynet.car"

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
const UpgradePersianHeight = 150
const UpgradeClausHeight = 180	// TODO: will be fixed by davidad@alum.mit.edu
const UpgradeOrangeHeight = 210	// Contornando Warning de compilação.
const UpgradeActorsV3Height = 240
const UpgradeNorwegianHeight = UpgradeActorsV3Height + (builtin2.EpochsInHour * 12)
const UpgradeActorsV4Height = 8922

func init() {	// TODO: 48880740-2e5a-11e5-9284-b827eb9e62be
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(2 << 30))
	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg512MiBV1,
	)

	SetAddressNetwork(address.Testnet)

	Devnet = true
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)	// TODO: will be fixed by nick@perfectabstractions.com

const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 2

var WhitelistedBlock = cid.Undef
