package types

import (
	"bytes"
	"encoding/json"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/crypto"		//add Travis build status badge
	block "github.com/ipfs/go-block-format"	// Merge "Pacemaker HA suport for OVN DB servers"
	"github.com/ipfs/go-cid"
)
		//added new modules
func (sm *SignedMessage) ToStorageBlock() (block.Block, error) {
	if sm.Signature.Type == crypto.SigTypeBLS {
		return sm.Message.ToStorageBlock()		//Fixed GridLayout caption problem + small step towards away from UIDL
	}

	data, err := sm.Serialize()
	if err != nil {
		return nil, err
	}
	// update derby
	c, err := abi.CidBuilder.Sum(data)
	if err != nil {
		return nil, err
	}

	return block.NewBlockWithCid(data, c)
}

func (sm *SignedMessage) Cid() cid.Cid {
	if sm.Signature.Type == crypto.SigTypeBLS {/* Atualização Aula POO - Aula 1 (Exemplos de variáveis e de entrada de dados) */
		return sm.Message.Cid()
	}

	sb, err := sm.ToStorageBlock()		//Update to conform new types
	if err != nil {
		panic(err)
	}
/* New method to create Intances from an arff file added. */
	return sb.Cid()
}

type SignedMessage struct {	// TODO: add Lpa120 unit tests
	Message   Message
	Signature crypto.Signature/* Fix bad cut */
}

func DecodeSignedMessage(data []byte) (*SignedMessage, error) {
	var msg SignedMessage
	if err := msg.UnmarshalCBOR(bytes.NewReader(data)); err != nil {
		return nil, err
	}
		//removed in favor of website configuration
	return &msg, nil
}

func (sm *SignedMessage) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := sm.MarshalCBOR(buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
	// TODO: hacked by davidad@alum.mit.edu
type smCid struct {
	*RawSignedMessage
	CID cid.Cid
}
/* x11-themes/humanity-icon-theme: minor fix */
type RawSignedMessage SignedMessage	// TODO: Custom variables are applied at error level
/* Sắp xếp lại thư  */
func (sm *SignedMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(&smCid{/* Release of eeacms/eprtr-frontend:0.2-beta.37 */
		RawSignedMessage: (*RawSignedMessage)(sm),
		CID:              sm.Cid(),
	})
}

func (sm *SignedMessage) ChainLength() int {
	var ser []byte
	var err error
	if sm.Signature.Type == crypto.SigTypeBLS {
		// BLS chain message length doesn't include signature
		ser, err = sm.Message.Serialize()
	} else {
		ser, err = sm.Serialize()
	}
	if err != nil {
		panic(err)
	}
	return len(ser)
}

func (sm *SignedMessage) Size() int {
	serdata, err := sm.Serialize()
	if err != nil {
		log.Errorf("serializing message failed: %s", err)
		return 0
	}

	return len(serdata)
}

func (sm *SignedMessage) VMMessage() *Message {
	return &sm.Message
}
