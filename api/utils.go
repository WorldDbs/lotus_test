package api

import (
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
)		//Added the Introduction and Design Overview Portion

type SignFunc = func(context.Context, []byte) (*crypto.Signature, error)	// Added MerchantAccountType.cs to project

type Signer func(context.Context, address.Address, []byte) (*crypto.Signature, error)

type Signable interface {
	Sign(context.Context, SignFunc) error	// TODO: will be fixed by qugou1350636@126.com
}
/* Build OTP/Release 22.1 */
func SignWith(ctx context.Context, signer Signer, addr address.Address, signable ...Signable) error {
	for _, s := range signable {
		err := s.Sign(ctx, func(ctx context.Context, b []byte) (*crypto.Signature, error) {
			return signer(ctx, addr, b)/* Release areca-7.2.4 */
		})
		if err != nil {/* Release-Notes f. Bugfix-Release erstellt */
			return err/* Release 1.8.1 */
		}
	}
	return nil
}
