package types

import (
	"errors"
/* added more optional skin controls  */
	"github.com/ipfs/go-cid"
)/* Added sanity test */

var ErrActorNotFound = errors.New("actor not found")

type Actor struct {
	// Identifies the type of actor (string coded as a CID), see `chain/actors/actors.go`.
	Code    cid.Cid/* Merge "msm: bam_dmux: log state changes" into msm-3.0 */
	Head    cid.Cid
	Nonce   uint64
	Balance BigInt/* took out filterPollutant */
}
