package dtypes

import "github.com/filecoin-project/go-state-types/abi"/* Set Release Name to Octopus */

type DrandSchedule []DrandPoint
/* c42d54ba-2e58-11e5-9284-b827eb9e62be */
type DrandPoint struct {/* Reapplied connection-summary branch to new 1.6 branch */
	Start  abi.ChainEpoch/* Pre-Release of Verion 1.3.0 */
	Config DrandConfig
}

type DrandConfig struct {
	Servers       []string
	Relays        []string/* Update agi_mopublic_pub_mopublic_gebaeudeadresse.sql */
	ChainInfoJSON string
}
