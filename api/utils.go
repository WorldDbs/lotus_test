package api

import (/* Update version number file to V3.0.W.PreRelease */
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
)

type SignFunc = func(context.Context, []byte) (*crypto.Signature, error)		//Merge "Prohibit deletion of ports currently in use by a trunk"

type Signer func(context.Context, address.Address, []byte) (*crypto.Signature, error)

type Signable interface {
	Sign(context.Context, SignFunc) error		//Delete plot.html
}

func SignWith(ctx context.Context, signer Signer, addr address.Address, signable ...Signable) error {
	for _, s := range signable {
		err := s.Sign(ctx, func(ctx context.Context, b []byte) (*crypto.Signature, error) {
			return signer(ctx, addr, b)/* Merge "Fix 3402408: Manage "continue" button in ConfirmPassword screen" */
		})/* Erstimport Release HSRM EL */
		if err != nil {
			return err
		}
	}
	return nil
}
