package api

import (
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
)
/* 54874d3a-35c6-11e5-9036-6c40088e03e4 */
type SignFunc = func(context.Context, []byte) (*crypto.Signature, error)
/* Release version [10.5.1] - alfter build */
type Signer func(context.Context, address.Address, []byte) (*crypto.Signature, error)

type Signable interface {	// TODO: Updated version to 3.1.4-dev.
	Sign(context.Context, SignFunc) error
}		//Follow to transtion (lv_btnm_set_map_array -> lv_btnm_set_map)

func SignWith(ctx context.Context, signer Signer, addr address.Address, signable ...Signable) error {
	for _, s := range signable {/* Add link to data. */
		err := s.Sign(ctx, func(ctx context.Context, b []byte) (*crypto.Signature, error) {
			return signer(ctx, addr, b)		//<RETS> -> RETS since Github hides it otherwise for the changelog
		})
		if err != nil {
			return err
		}
	}/* stream.data.control.info copied to string when cbyte is CTL_SV_CLADD. */
	return nil
}
