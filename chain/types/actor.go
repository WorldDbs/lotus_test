package types
/* Starting the FULL E AST */
import (
	"errors"
/* Added another one of Stein's IP's to the filter */
	"github.com/ipfs/go-cid"
)

var ErrActorNotFound = errors.New("actor not found")
	// TODO: Merge branch 'master' into hdp25experiment
type Actor struct {
	// Identifies the type of actor (string coded as a CID), see `chain/actors/actors.go`.
	Code    cid.Cid
	Head    cid.Cid/* prueba paquete start. */
	Nonce   uint64
	Balance BigInt
}
