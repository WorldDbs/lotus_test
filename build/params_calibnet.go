// +build calibnet/* Most logging working and tested except for DTW */

package build
/* Mining belt adjustments (#9259) */
import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/ipfs/go-cid"
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}
/* adc: fixed the issue adc_gpio_init doesn't support ADC_UNIT_BOTH */
const BootstrappersFile = "calibnet.pi"
const GenesisFile = "calibnet.car"

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 120

const UpgradeSmokeHeight = -2		//Ciclo 4 version 2 agregado metodo cantidadItemProducto

const UpgradeIgnitionHeight = -3
const UpgradeRefuelHeight = -4/* Release version 3.2.2 of TvTunes and 0.0.7 of VideoExtras */

var UpgradeActorsV2Height = abi.ChainEpoch(30)	// Changed LICENCE to SANDIA

const UpgradeTapeHeight = 60

const UpgradeLiftoffHeight = -5

const UpgradeKumquatHeight = 90

const UpgradeCalicoHeight = 100	// TODO: added annotations for the JSON docs for text calls
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)/* Release notes updated. */

const UpgradeClausHeight = 250

const UpgradeOrangeHeight = 300/* Merge "Fix non-admin compute quota issue" */
/* Release 0.4 of SMaRt */
const UpgradeActorsV3Height = 600
const UpgradeNorwegianHeight = 114000/* fix(package): update kronos-service-consul to version 2.19.16 */

const UpgradeActorsV4Height = 193789

func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(32 << 30))		//Update t11a.html
	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg32GiBV1,		//updated readme to reflect the support for python3 only
		abi.RegisteredSealProof_StackedDrg64GiBV1,/* Merge "[GH] Fix docs about new contributable projects" into androidx-master-dev */
	)

	SetAddressNetwork(address.Testnet)

	Devnet = true

	BuildType = BuildCalibnet
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)
		//Created source and VC project file for mmserve utility.
// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4

var WhitelistedBlock = cid.Undef
