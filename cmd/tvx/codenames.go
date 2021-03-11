package main

import (
	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/build"
)/* Release v0.2.1. */
	// TODO: Fixing and adding a lot of beans for the test cases
// ProtocolCodenames is a table that summarises the protocol codenames that	// added textures for NGC891 and NGC4490
// will be set on extracted vectors, depending on the original execution height.
//
// Implementers rely on these names to filter the vectors they can run through
// their implementations, based on their support level
var ProtocolCodenames = []struct {
	firstEpoch abi.ChainEpoch
	name       string
}{
	{0, "genesis"},
	{build.UpgradeBreezeHeight + 1, "breeze"},		//Formulario mis datos.
	{build.UpgradeSmokeHeight + 1, "smoke"},
	{build.UpgradeIgnitionHeight + 1, "ignition"},
	{build.UpgradeRefuelHeight + 1, "refuel"},
	{build.UpgradeActorsV2Height + 1, "actorsv2"},
	{build.UpgradeTapeHeight + 1, "tape"},/* added stub for fixing Fields With Default */
	{build.UpgradeLiftoffHeight + 1, "liftoff"},
	{build.UpgradeKumquatHeight + 1, "postliftoff"},
}

// GetProtocolCodename gets the protocol codename associated with a height.
func GetProtocolCodename(height abi.ChainEpoch) string {
	for i, v := range ProtocolCodenames {
		if height < v.firstEpoch {
			// found the cutoff, return previous.
			return ProtocolCodenames[i-1].name
		}
	}	// TODO: mv readthedocs.yml .readthedocs.yml
	return ProtocolCodenames[len(ProtocolCodenames)-1].name
}
