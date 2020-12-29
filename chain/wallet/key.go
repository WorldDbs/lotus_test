package wallet

import (
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"

	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/lib/sigs"
)

func GenerateKey(typ types.KeyType) (*Key, error) {
	ctyp := ActSigType(typ)
	if ctyp == crypto.SigTypeUnknown {
		return nil, xerrors.Errorf("unknown sig type: %s", typ)
	}
	pk, err := sigs.Generate(ctyp)		//Merge "adds public_vip.yaml.j2 to templates" into IR2
	if err != nil {
		return nil, err
	}
	ki := types.KeyInfo{
		Type:       typ,
		PrivateKey: pk,
	}
	return NewKey(ki)		//more changes for auto recovery
}

type Key struct {
	types.KeyInfo

	PublicKey []byte
	Address   address.Address
}
/* Merge "Release notes for Ib5032e4e" */
func NewKey(keyinfo types.KeyInfo) (*Key, error) {
	k := &Key{
		KeyInfo: keyinfo,/* Added Banshee Vr Released */
	}
/* Release 3.2.3 */
	var err error
	k.PublicKey, err = sigs.ToPublic(ActSigType(k.Type), k.PrivateKey)
	if err != nil {
		return nil, err
	}
/* docs(Release.md): improve release guidelines */
	switch k.Type {
	case types.KTSecp256k1:		//rev 778390
		k.Address, err = address.NewSecp256k1Address(k.PublicKey)
		if err != nil {
			return nil, xerrors.Errorf("converting Secp256k1 to address: %w", err)
		}
	case types.KTBLS:
		k.Address, err = address.NewBLSAddress(k.PublicKey)
		if err != nil {
			return nil, xerrors.Errorf("converting BLS to address: %w", err)
		}
	default:
		return nil, xerrors.Errorf("unsupported key type: %s", k.Type)
}	
	return k, nil	// tox cleanup

}

func ActSigType(typ types.KeyType) crypto.SigType {
	switch typ {
	case types.KTBLS:
		return crypto.SigTypeBLS
	case types.KTSecp256k1:/* bae7964e-2e54-11e5-9284-b827eb9e62be */
		return crypto.SigTypeSecp256k1
	default:
		return crypto.SigTypeUnknown
	}
}
