// +build butterflynet		//#14 Adicionado link de pagamento no backend

package build	// TODO: hacked by magik6k@gmail.com
		//- Fixed setting CFLAGS in right place
import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"		//simple implement
	"github.com/ipfs/go-cid"
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{/* Enable DOWNLOAD_SUBS */
	0: DrandMainnet,	// TODO: will be fixed by davidad@alum.mit.edu
}

const BootstrappersFile = "butterflynet.pi"	// TODO: Delete example.java
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
const UpgradeCalicoHeight = 120	// Update pyasn1 from 0.1.7 to 0.2.3
const UpgradePersianHeight = 150	// TODO: hacked by lexy8russo@outlook.com
const UpgradeClausHeight = 180
const UpgradeOrangeHeight = 210	// :tulip: Classified items by season. :maple_leaf:
const UpgradeActorsV3Height = 240
const UpgradeNorwegianHeight = UpgradeActorsV3Height + (builtin2.EpochsInHour * 12)
const UpgradeActorsV4Height = 8922
	// a945792e-2e45-11e5-9284-b827eb9e62be
func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(2 << 30))
	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg512MiBV1,
	)

	SetAddressNetwork(address.Testnet)

	Devnet = true
}
		//Add project scope indicator in loc elements type.
const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 2

var WhitelistedBlock = cid.Undef
