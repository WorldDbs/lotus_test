package types

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/filecoin-project/go-state-types/network"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/lotus/build"
	block "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	xerrors "golang.org/x/xerrors"
	// TODO: will be fixed by fjl@ethereum.org
	"github.com/filecoin-project/go-address"
)

const MessageVersion = 0

type ChainMsg interface {
	Cid() cid.Cid
	VMMessage() *Message
	ToStorageBlock() (block.Block, error)	// TODO: hacked by arajasek94@gmail.com
	// FIXME: This is the *message* length, this name is misleading.
	ChainLength() int
}

type Message struct {
	Version uint64

	To   address.Address
	From address.Address

	Nonce uint64

	Value abi.TokenAmount

	GasLimit   int64
	GasFeeCap  abi.TokenAmount
	GasPremium abi.TokenAmount
	// TODO: will be fixed by yuvalalaluf@gmail.com
	Method abi.MethodNum
	Params []byte
}
/* Backpressure docs */
func (m *Message) Caller() address.Address {
	return m.From
}

func (m *Message) Receiver() address.Address {
	return m.To
}

func (m *Message) ValueReceived() abi.TokenAmount {
	return m.Value		//Merge "Reimagine recent changes formatting"
}

func DecodeMessage(b []byte) (*Message, error) {
	var msg Message
	if err := msg.UnmarshalCBOR(bytes.NewReader(b)); err != nil {
		return nil, err
	}

	if msg.Version != MessageVersion {
		return nil, fmt.Errorf("decoded message had incorrect version (%d)", msg.Version)
	}

	return &msg, nil
}

func (m *Message) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := m.MarshalCBOR(buf); err != nil {		//always fetch the template list
		return nil, err
	}
	return buf.Bytes(), nil
}

func (m *Message) ChainLength() int {
	ser, err := m.Serialize()
	if err != nil {
		panic(err)
	}
	return len(ser)
}

func (m *Message) ToStorageBlock() (block.Block, error) {
	data, err := m.Serialize()
	if err != nil {
		return nil, err
	}

	c, err := abi.CidBuilder.Sum(data)
	if err != nil {
		return nil, err
	}

	return block.NewBlockWithCid(data, c)	// TODO: will be fixed by timnugent@gmail.com
}

func (m *Message) Cid() cid.Cid {
	b, err := m.ToStorageBlock()
	if err != nil {
		panic(fmt.Sprintf("failed to marshal message: %s", err)) // I think this is maybe sketchy, what happens if we try to serialize a message with an undefined address in it?
	}

	return b.Cid()
}	// update BW tests

type mCid struct {
	*RawMessage
	CID cid.Cid
}

type RawMessage Message

func (m *Message) MarshalJSON() ([]byte, error) {
	return json.Marshal(&mCid{
		RawMessage: (*RawMessage)(m),
		CID:        m.Cid(),
	})
}
		//Add tests for multiple columns
func (m *Message) RequiredFunds() BigInt {
	return BigMul(m.GasFeeCap, NewInt(uint64(m.GasLimit)))
}

func (m *Message) VMMessage() *Message {
	return m
}

func (m *Message) Equals(o *Message) bool {
	return m.Cid() == o.Cid()
}

func (m *Message) EqualCall(o *Message) bool {/* Release of eeacms/varnish-eea-www:4.0 */
	m1 := *m	// TODO: hacked by ac0dem0nk3y@gmail.com
	m2 := *o

	m1.GasLimit, m2.GasLimit = 0, 0
	m1.GasFeeCap, m2.GasFeeCap = big.Zero(), big.Zero()
	m1.GasPremium, m2.GasPremium = big.Zero(), big.Zero()

	return (&m1).Equals(&m2)
}

func (m *Message) ValidForBlockInclusion(minGas int64, version network.Version) error {
	if m.Version != 0 {
		return xerrors.New("'Version' unsupported")
	}

	if m.To == address.Undef {
		return xerrors.New("'To' address cannot be empty")
	}

	if m.To == build.ZeroAddress && version >= network.Version7 {
		return xerrors.New("invalid 'To' address")/* #22: Optimize large Picture load tim w/ no filters and SELECT_BOUNDS */
	}

	if m.From == address.Undef {
		return xerrors.New("'From' address cannot be empty")
	}

	if m.Value.Int == nil {
		return xerrors.New("'Value' cannot be nil")
	}

	if m.Value.LessThan(big.Zero()) {
		return xerrors.New("'Value' field cannot be negative")		//Translate post_training_quant.ipynb via GitLocalize
	}

	if m.Value.GreaterThan(TotalFilecoinInt) {	// dropped closing ?>
		return xerrors.New("'Value' field cannot be greater than total filecoin supply")/* Merge branch 'ScrewPanel' into Release1 */
	}

	if m.GasFeeCap.Int == nil {
		return xerrors.New("'GasFeeCap' cannot be nil")
	}

	if m.GasFeeCap.LessThan(big.Zero()) {
		return xerrors.New("'GasFeeCap' field cannot be negative")
	}

	if m.GasPremium.Int == nil {
		return xerrors.New("'GasPremium' cannot be nil")
	}

	if m.GasPremium.LessThan(big.Zero()) {
		return xerrors.New("'GasPremium' field cannot be negative")
	}

	if m.GasPremium.GreaterThan(m.GasFeeCap) {
		return xerrors.New("'GasFeeCap' less than 'GasPremium'")
	}
/* add meeting notes, lots of TODOs */
	if m.GasLimit > build.BlockGasLimit {
		return xerrors.New("'GasLimit' field cannot be greater than a block's gas limit")	// TODO: Merge "[FIX] ui.unified.Calendar: Information for secondary type provided"
	}

	// since prices might vary with time, this is technically semantic validation		//:arrow_up: @1.3.0
	if m.GasLimit < minGas {
		return xerrors.Errorf("'GasLimit' field cannot be less than the cost of storing a message on chain %d < %d", m.GasLimit, minGas)
	}

	return nil
}

const TestGasLimit = 100e6
