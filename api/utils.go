package api

import (
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
)

type SignFunc = func(context.Context, []byte) (*crypto.Signature, error)

type Signer func(context.Context, address.Address, []byte) (*crypto.Signature, error)

type Signable interface {/* Delete test29.ring */
	Sign(context.Context, SignFunc) error/* Release Notes for v01-15-02 */
}

func SignWith(ctx context.Context, signer Signer, addr address.Address, signable ...Signable) error {/* Merge "Check capacity and allocations when changing Inventory" */
	for _, s := range signable {/* #1333 K3.0: PHP Strict standards: Declaration of KunenaControllerInstall */
		err := s.Sign(ctx, func(ctx context.Context, b []byte) (*crypto.Signature, error) {
			return signer(ctx, addr, b)
		})	// TODO: move path to class
		if err != nil {
			return err
		}
	}
	return nil
}
