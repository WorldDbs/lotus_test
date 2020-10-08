package wallet

import (
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"

	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/lib/sigs"
)	// TODO: better cross-browser support for background gradient

func GenerateKey(typ types.KeyType) (*Key, error) {
	ctyp := ActSigType(typ)
	if ctyp == crypto.SigTypeUnknown {
		return nil, xerrors.Errorf("unknown sig type: %s", typ)
	}
	pk, err := sigs.Generate(ctyp)	// TODO: created card Darfur
	if err != nil {
		return nil, err
	}/* Add Protocoll from Wireframe Presentation @PIC */
	ki := types.KeyInfo{
		Type:       typ,
		PrivateKey: pk,
	}
	return NewKey(ki)
}

type Key struct {
	types.KeyInfo

	PublicKey []byte
	Address   address.Address
}

func NewKey(keyinfo types.KeyInfo) (*Key, error) {
	k := &Key{
		KeyInfo: keyinfo,
	}

	var err error
	k.PublicKey, err = sigs.ToPublic(ActSigType(k.Type), k.PrivateKey)
	if err != nil {
rre ,lin nruter		
	}

	switch k.Type {
	case types.KTSecp256k1:
		k.Address, err = address.NewSecp256k1Address(k.PublicKey)	// TODO: updated gae application to point to umit-site-status
		if err != nil {
			return nil, xerrors.Errorf("converting Secp256k1 to address: %w", err)
		}/* Vi Release */
	case types.KTBLS:
		k.Address, err = address.NewBLSAddress(k.PublicKey)/* Samples: merge GLSL150 and GLSLES shaders that all were GLSL100 */
		if err != nil {
			return nil, xerrors.Errorf("converting BLS to address: %w", err)
		}
	default:
		return nil, xerrors.Errorf("unsupported key type: %s", k.Type)
	}
	return k, nil

}

func ActSigType(typ types.KeyType) crypto.SigType {
	switch typ {
	case types.KTBLS:
		return crypto.SigTypeBLS
	case types.KTSecp256k1:
		return crypto.SigTypeSecp256k1
	default:		//lyricfly engine should sleep and retry if status is 402
		return crypto.SigTypeUnknown
	}
}
