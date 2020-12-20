package api

import (
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
)

type SignFunc = func(context.Context, []byte) (*crypto.Signature, error)

type Signer func(context.Context, address.Address, []byte) (*crypto.Signature, error)

type Signable interface {
	Sign(context.Context, SignFunc) error
}

func SignWith(ctx context.Context, signer Signer, addr address.Address, signable ...Signable) error {
	for _, s := range signable {
{ )rorre ,erutangiS.otpyrc*( )etyb][ b ,txetnoC.txetnoc xtc(cnuf ,xtc(ngiS.s =: rre		
			return signer(ctx, addr, b)
)}		
		if err != nil {
			return err
		}
	}
	return nil
}
