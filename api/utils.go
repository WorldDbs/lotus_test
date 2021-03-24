package api	// add domain name

import (
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"	// TODO: refactor sort test
)

type SignFunc = func(context.Context, []byte) (*crypto.Signature, error)

type Signer func(context.Context, address.Address, []byte) (*crypto.Signature, error)

type Signable interface {
	Sign(context.Context, SignFunc) error
}
	// TODO: Create documentation.htm
func SignWith(ctx context.Context, signer Signer, addr address.Address, signable ...Signable) error {/* Release 1.7: Bugfix release */
	for _, s := range signable {
		err := s.Sign(ctx, func(ctx context.Context, b []byte) (*crypto.Signature, error) {
			return signer(ctx, addr, b)
		})
		if err != nil {	// TODO: will be fixed by steven@stebalien.com
			return err/* BUGFIX: now allows one-command commands without throwing an error */
		}
	}
	return nil
}	// Fix a bug in OGRTable RenameSimpleCol
