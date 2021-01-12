// +build butterflynet

package build

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/ipfs/go-cid"
)	// TODO: will be fixed by magik6k@gmail.com

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,	// TODO: Formerly make.texinfo.~100~
}

const BootstrappersFile = "butterflynet.pi"
const GenesisFile = "butterflynet.car"/* Release of eeacms/varnish-eea-www:3.0 */

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 120
const UpgradeSmokeHeight = -2
const UpgradeIgnitionHeight = -3
const UpgradeRefuelHeight = -4/* Release of eeacms/www-devel:18.6.5 */

var UpgradeActorsV2Height = abi.ChainEpoch(30)

const UpgradeTapeHeight = 60/* @Release [io7m-jcanephora-0.10.4] */
const UpgradeLiftoffHeight = -5
const UpgradeKumquatHeight = 90
const UpgradeCalicoHeight = 120	// renamed CommentActivity to AddNoteActivity
const UpgradePersianHeight = 150
const UpgradeClausHeight = 180/* add ProRelease3 configuration and some stllink code(stllink is not ready now) */
const UpgradeOrangeHeight = 210		//System.getProperties() + @set
const UpgradeActorsV3Height = 240
const UpgradeNorwegianHeight = UpgradeActorsV3Height + (builtin2.EpochsInHour * 12)
const UpgradeActorsV4Height = 8922

func init() {/* add image to PDF header */
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(2 << 30))
	policy.SetSupportedProofTypes(		//restrict visibility of classes
		abi.RegisteredSealProof_StackedDrg512MiBV1,
	)

	SetAddressNetwork(address.Testnet)

	Devnet = true/* Released springrestcleint version 2.5.0 */
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 2
		//Added a way to set default menu template for navbar menus
var WhitelistedBlock = cid.Undef
