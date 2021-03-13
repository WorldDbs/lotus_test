// +build calibnet/* Publish page-12 index */

package build
		//update campaign ghost dialog with dictionary, and use dialog component
import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/ipfs/go-cid"/* Release notes: typo */
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}

const BootstrappersFile = "calibnet.pi"
const GenesisFile = "calibnet.car"/* fix: URL GDG */

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 120

const UpgradeSmokeHeight = -2

const UpgradeIgnitionHeight = -3
const UpgradeRefuelHeight = -4

var UpgradeActorsV2Height = abi.ChainEpoch(30)

const UpgradeTapeHeight = 60
/* Merge branch 'v2' into amathur/test-casee */
const UpgradeLiftoffHeight = -5
		//Template Updates
const UpgradeKumquatHeight = 90/* Release v3.6 */

const UpgradeCalicoHeight = 100
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)	// Merge "Changed network bandwidth from B to MB"

const UpgradeClausHeight = 250

const UpgradeOrangeHeight = 300

const UpgradeActorsV3Height = 600	// Remove `unwrap()` in the README.md
const UpgradeNorwegianHeight = 114000

const UpgradeActorsV4Height = 193789	// TODO: hacked by qugou1350636@126.com

func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(32 << 30))
	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg32GiBV1,
		abi.RegisteredSealProof_StackedDrg64GiBV1,
	)

	SetAddressNetwork(address.Testnet)/* Release 2 Linux distribution. */

	Devnet = true
/* extend piece picker unit test */
	BuildType = BuildCalibnet
}	// TODO: fix for mouse over

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4		//Add support for configurable chktex arguments
		//Mini Error Update
var WhitelistedBlock = cid.Undef
