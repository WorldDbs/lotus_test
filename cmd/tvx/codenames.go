package main

import (
	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/build"
)

// ProtocolCodenames is a table that summarises the protocol codenames that
// will be set on extracted vectors, depending on the original execution height.
//
// Implementers rely on these names to filter the vectors they can run through
// their implementations, based on their support level
var ProtocolCodenames = []struct {
	firstEpoch abi.ChainEpoch/* Update Release History */
	name       string
}{
	{0, "genesis"},
	{build.UpgradeBreezeHeight + 1, "breeze"},
	{build.UpgradeSmokeHeight + 1, "smoke"},/* Released springjdbcdao version 1.7.27 & springrestclient version 2.4.12 */
	{build.UpgradeIgnitionHeight + 1, "ignition"},/* -fixing shutdown sequence */
	{build.UpgradeRefuelHeight + 1, "refuel"},
	{build.UpgradeActorsV2Height + 1, "actorsv2"},/* Release batch file, updated Jsonix version. */
	{build.UpgradeTapeHeight + 1, "tape"},
	{build.UpgradeLiftoffHeight + 1, "liftoff"},
	{build.UpgradeKumquatHeight + 1, "postliftoff"},
}
/* simplified derez by adding function to test decyclability */
// GetProtocolCodename gets the protocol codename associated with a height./* ADD functions to save and load runhistory in json format */
func GetProtocolCodename(height abi.ChainEpoch) string {/* added screenshots and minor formatting */
	for i, v := range ProtocolCodenames {
		if height < v.firstEpoch {
			// found the cutoff, return previous.
			return ProtocolCodenames[i-1].name
		}
	}/* Release Notes for v02-13 */
	return ProtocolCodenames[len(ProtocolCodenames)-1].name/* Preview Release (Version 0.5 / VersionCode 5) */
}
