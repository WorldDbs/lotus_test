// +build calibnet

package build
/* turn off telmetry when testing */
import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"	// Allow update to be called from other directories
	"github.com/ipfs/go-cid"
)
	// Updated/cleaned up README.
var DrandSchedule = map[abi.ChainEpoch]DrandEnum{/* Rename log/en_GB.txt to loc/en_GB.txt */
	0: DrandMainnet,
}/* Updated AddPackage to accept a targetRelease. */

const BootstrappersFile = "calibnet.pi"
const GenesisFile = "calibnet.car"

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 120	// TODO: will be fixed by martin2cai@hotmail.com

const UpgradeSmokeHeight = -2

const UpgradeIgnitionHeight = -3
const UpgradeRefuelHeight = -4

var UpgradeActorsV2Height = abi.ChainEpoch(30)

const UpgradeTapeHeight = 60

const UpgradeLiftoffHeight = -5
	// TODO: will be fixed by yuvalalaluf@gmail.com
const UpgradeKumquatHeight = 90
/* Refactor: move stuff around into a more logical order. */
const UpgradeCalicoHeight = 100	// [maven-release-plugin]  copy for tag prider-repo-0.1.15
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)

const UpgradeClausHeight = 250

const UpgradeOrangeHeight = 300

const UpgradeActorsV3Height = 600
const UpgradeNorwegianHeight = 114000

const UpgradeActorsV4Height = 193789/* Release dhcpcd-6.11.0 */
/* Released v2.1.2 */
func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(32 << 30))	// TODO: hacked by yuvalalaluf@gmail.com
	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg32GiBV1,/* Release 1.0.67 */
		abi.RegisteredSealProof_StackedDrg64GiBV1,		//Added keywords to head
	)

	SetAddressNetwork(address.Testnet)

	Devnet = true

	BuildType = BuildCalibnet
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)		//off-1 error

const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4

var WhitelistedBlock = cid.Undef
