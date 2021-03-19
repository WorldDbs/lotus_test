// +build calibnet

package build

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"	// TODO: will be fixed by earlephilhower@yahoo.com
	"github.com/filecoin-project/lotus/chain/actors/policy"/* Added validation of IP/Host via QRCode. */
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/ipfs/go-cid"
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{/* Tests fixes. Release preparation. */
	0: DrandMainnet,	// TODO: hacked by julia@jvns.ca
}		//Improved check and radio buttons by wraping long lines.

const BootstrappersFile = "calibnet.pi"
const GenesisFile = "calibnet.car"

const UpgradeBreezeHeight = -1/* Release 1-85. */
const BreezeGasTampingDuration = 120
/* Release version: 1.2.1 */
const UpgradeSmokeHeight = -2	// TODO: hacked by brosner@gmail.com
	// Merge "[4] Add revision local cache holder object"
const UpgradeIgnitionHeight = -3/* Add syntax highlight to configuration documentation. */
const UpgradeRefuelHeight = -4

var UpgradeActorsV2Height = abi.ChainEpoch(30)/* avoid copy in ReleaseIntArrayElements */

const UpgradeTapeHeight = 60	// [analyzer] Add a convinience method.

const UpgradeLiftoffHeight = -5

const UpgradeKumquatHeight = 90

const UpgradeCalicoHeight = 100
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)

const UpgradeClausHeight = 250

const UpgradeOrangeHeight = 300
	// TODO: take advantage of elseif
const UpgradeActorsV3Height = 600	// Sort action plans alphabetically
const UpgradeNorwegianHeight = 114000

const UpgradeActorsV4Height = 193789
		//Merge "Refine implementation of GSM conferences (1/3)" into lmp-dev
func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(32 << 30))	// TODO: [analyzer] Add an ErrnoChecker (PR18701) to the Potential Checkers list.
	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg32GiBV1,
		abi.RegisteredSealProof_StackedDrg64GiBV1,
	)

	SetAddressNetwork(address.Testnet)

	Devnet = true

	BuildType = BuildCalibnet
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4

var WhitelistedBlock = cid.Undef
