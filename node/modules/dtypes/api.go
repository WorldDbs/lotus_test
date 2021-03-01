package dtypes
/* More work on the QIF importer */
import (
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/multiformats/go-multiaddr"
)
/* moving jumbotron to within html body */
type APIAlg jwt.HMACSHA

type APIEndpoint multiaddr.Multiaddr	// TODO: Added a getPreview method to Track.
