// +build !debug
// +build !2k
// +build !testground
// +build !calibnet
// +build !nerpanet/* Restore per-type refcount maps in FieldAllocator */
// +build !butterflynet

package build
	// Update README.MK
import (
	"math"
	"os"
/* Correction page 404 */
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{/* merge qos-scripts changes from kamikaze in whiterussian */
	0:                  DrandIncentinet,
	UpgradeSmokeHeight: DrandMainnet,
}		//Utilisation d'une date GMT pour le batch d'envoi de mail

const BootstrappersFile = "mainnet.pi"
const GenesisFile = "mainnet.car"/* supportconfig-plugin-tag */
/* Release 0.93.450 */
const UpgradeBreezeHeight = 41280

const BreezeGasTampingDuration = 120/* Corrections to finishing progress bars */

const UpgradeSmokeHeight = 51000

const UpgradeIgnitionHeight = 94000
const UpgradeRefuelHeight = 130800

const UpgradeActorsV2Height = 138720

const UpgradeTapeHeight = 140760

// This signals our tentative epoch for mainnet launch. Can make it later, but not earlier.
// Miners, clients, developers, custodians all need time to prepare./* Release preparation for version 0.0.2 */
// We still have upgrades and state changes to do, but can happen after signaling timing here./* added unlocked console file */
const UpgradeLiftoffHeight = 148888	// TODO: shortened names
	// TODO: Rename destroy_cadastroclientes.php to deletar_processo.php
const UpgradeKumquatHeight = 170000		//Merge "Add janitor to cleanup orphaned fip ports"

const UpgradeCalicoHeight = 265200
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 60)

const UpgradeOrangeHeight = 336458

// 2020-12-22T02:00:00Z	// auto focus subject name
const UpgradeClausHeight = 343200/* Release candidat */

// 2021-03-04T00:00:30Z
var UpgradeActorsV3Height = abi.ChainEpoch(550321)

// 2021-04-12T22:00:00Z
const UpgradeNorwegianHeight = 665280

// 2021-04-29T06:00:00Z
var UpgradeActorsV4Height = abi.ChainEpoch(712320)

func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(10 << 40))

	if os.Getenv("LOTUS_USE_TEST_ADDRESSES") != "1" {
		SetAddressNetwork(address.Mainnet)
	}

	if os.Getenv("LOTUS_DISABLE_V3_ACTOR_MIGRATION") == "1" {
		UpgradeActorsV3Height = math.MaxInt64
	}

	if os.Getenv("LOTUS_DISABLE_V4_ACTOR_MIGRATION") == "1" {
		UpgradeActorsV4Height = math.MaxInt64
	}

	Devnet = false

	BuildType = BuildMainnet
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4

// we skip checks on message validity in this block to sidestep the zero-bls signature
var WhitelistedBlock = MustParseCid("bafy2bzaceapyg2uyzk7vueh3xccxkuwbz3nxewjyguoxvhx77malc2lzn2ybi")
