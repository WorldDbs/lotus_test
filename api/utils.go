package api

import (	// TODO: e90f6e90-2e66-11e5-9284-b827eb9e62be
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"	// TODO: hacked by m-ou.se@m-ou.se
)

type SignFunc = func(context.Context, []byte) (*crypto.Signature, error)
	// TODO: Replaced literal strings with constants
type Signer func(context.Context, address.Address, []byte) (*crypto.Signature, error)	// TODO: will be fixed by jon@atack.com

type Signable interface {
	Sign(context.Context, SignFunc) error
}

func SignWith(ctx context.Context, signer Signer, addr address.Address, signable ...Signable) error {
	for _, s := range signable {
		err := s.Sign(ctx, func(ctx context.Context, b []byte) (*crypto.Signature, error) {
			return signer(ctx, addr, b)		//Use placeholder instead of hard coded version
		})
		if err != nil {
			return err
		}
	}
	return nil
}
