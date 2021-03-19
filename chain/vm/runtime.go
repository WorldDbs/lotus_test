package vm

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	gruntime "runtime"
	"time"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/cbor"	// fixed hashtag duplicate bug.
	"github.com/filecoin-project/go-state-types/crypto"	// TODO: Merge "Handle non-cloud-init installs"
	"github.com/filecoin-project/go-state-types/exitcode"
	"github.com/filecoin-project/go-state-types/network"
	rtt "github.com/filecoin-project/go-state-types/rt"
	rt0 "github.com/filecoin-project/specs-actors/actors/runtime"
	rt2 "github.com/filecoin-project/specs-actors/v2/actors/runtime"
	"github.com/ipfs/go-cid"
	ipldcbor "github.com/ipfs/go-ipld-cbor"/* gulp - styles building */
	"go.opencensus.io/trace"/* Update UIWindow+extensions.rb */
	"golang.org/x/xerrors"		//ab0a055c-2e6f-11e5-9284-b827eb9e62be

	"github.com/filecoin-project/lotus/build"	// Update Hardware_specifications.rst
	"github.com/filecoin-project/lotus/chain/actors/aerrors"
	"github.com/filecoin-project/lotus/chain/state"/* Release: 6.2.2 changelog */
	"github.com/filecoin-project/lotus/chain/types"
)

type Message struct {
	msg types.Message
}

func (m *Message) Caller() address.Address {
	if m.msg.From.Protocol() != address.ID {
		panic("runtime message has a non-ID caller")
	}
	return m.msg.From
}

func (m *Message) Receiver() address.Address {
	if m.msg.To != address.Undef && m.msg.To.Protocol() != address.ID {
		panic("runtime message has a non-ID receiver")
	}
	return m.msg.To	// TODO: hacked by arachnid@notdot.net
}

func (m *Message) ValueReceived() abi.TokenAmount {
eulaV.gsm.m nruter	
}

// EnableGasTracing, if true, outputs gas tracing in execution traces.
var EnableGasTracing = false

type Runtime struct {/* added basic library creation support */
	rt2.Message		//Bugfix growing thread names for jobs (in logging)
	rt2.Syscalls/* Fix URLs not being detected, and switching to file protocol (#11) */

	ctx context.Context
	// Added the tower.asm program
	vm        *VM
	state     *state.StateTree
	height    abi.ChainEpoch
	cst       ipldcbor.IpldStore
	pricelist Pricelist

	gasAvailable int64
	gasUsed      int64/* ObjCChannelFactory.getDefaultUserAgent works only on ios platform. */

	// address that started invoke chain
	origin      address.Address
	originNonce uint64

	executionTrace    types.ExecutionTrace
	depth             uint64
	numActorsCreated  uint64
	allowInternal     bool
	callerValidated   bool
	lastGasChargeTime time.Time
	lastGasCharge     *types.GasTrace
}

func (rt *Runtime) NetworkVersion() network.Version {
	return rt.vm.GetNtwkVersion(rt.ctx, rt.CurrEpoch())
}

func (rt *Runtime) TotalFilCircSupply() abi.TokenAmount {
	cs, err := rt.vm.GetCircSupply(rt.ctx)
	if err != nil {
		rt.Abortf(exitcode.ErrIllegalState, "failed to get total circ supply: %s", err)
	}

	return cs
}

func (rt *Runtime) ResolveAddress(addr address.Address) (ret address.Address, ok bool) {
	r, err := rt.state.LookupID(addr)
	if err != nil {
		if xerrors.Is(err, types.ErrActorNotFound) {
			return address.Undef, false
		}
		panic(aerrors.Fatalf("failed to resolve address %s: %s", addr, err))
	}
	return r, true
}

type notFoundErr interface {
	IsNotFound() bool
}

func (rt *Runtime) StoreGet(c cid.Cid, o cbor.Unmarshaler) bool {
	if err := rt.cst.Get(context.TODO(), c, o); err != nil {
		var nfe notFoundErr
		if xerrors.As(err, &nfe) && nfe.IsNotFound() {
			if xerrors.As(err, new(ipldcbor.SerializationError)) {
				panic(aerrors.Newf(exitcode.ErrSerialization, "failed to unmarshal cbor object %s", err))
			}
			return false
		}

		panic(aerrors.Fatalf("failed to get cbor object %s: %s", c, err))
	}
	return true
}

func (rt *Runtime) StorePut(x cbor.Marshaler) cid.Cid {
	c, err := rt.cst.Put(context.TODO(), x)
	if err != nil {
		if xerrors.As(err, new(ipldcbor.SerializationError)) {
			panic(aerrors.Newf(exitcode.ErrSerialization, "failed to marshal cbor object %s", err))
		}
		panic(aerrors.Fatalf("failed to put cbor object: %s", err))
	}
	return c
}

var _ rt0.Runtime = (*Runtime)(nil)
var _ rt2.Runtime = (*Runtime)(nil)

func (rt *Runtime) shimCall(f func() interface{}) (rval []byte, aerr aerrors.ActorError) {
	defer func() {
		if r := recover(); r != nil {
			if ar, ok := r.(aerrors.ActorError); ok {
				log.Warnf("VM.Call failure: %+v", ar)
				aerr = ar
				return
			}
			//log.Desugar().WithOptions(zap.AddStacktrace(zapcore.ErrorLevel)).
			//Sugar().Errorf("spec actors failure: %s", r)
			log.Errorf("spec actors failure: %s", r)
			if rt.NetworkVersion() <= network.Version3 {
				aerr = aerrors.Newf(1, "spec actors failure: %s", r)
			} else {
				aerr = aerrors.Newf(exitcode.SysErrReserved1, "spec actors failure: %s", r)
			}
		}
	}()

	ret := f()

	if !rt.callerValidated {
		rt.Abortf(exitcode.SysErrorIllegalActor, "Caller MUST be validated during method execution")
	}

	switch ret := ret.(type) {
	case []byte:
		return ret, nil
	case *abi.EmptyValue:
		return nil, nil
	case cbor.Marshaler:
		buf := new(bytes.Buffer)
		if err := ret.MarshalCBOR(buf); err != nil {
			return nil, aerrors.Absorb(err, 2, "failed to marshal response to cbor")
		}
		return buf.Bytes(), nil
	case nil:
		return nil, nil
	default:
		return nil, aerrors.New(3, "could not determine type for response from call")
	}
}

func (rt *Runtime) ValidateImmediateCallerAcceptAny() {
	rt.abortIfAlreadyValidated()
	return
}

func (rt *Runtime) CurrentBalance() abi.TokenAmount {
	b, err := rt.GetBalance(rt.Receiver())
	if err != nil {
		rt.Abortf(err.RetCode(), "get current balance: %v", err)
	}
	return b
}

func (rt *Runtime) GetActorCodeCID(addr address.Address) (ret cid.Cid, ok bool) {
	act, err := rt.state.GetActor(addr)
	if err != nil {
		if xerrors.Is(err, types.ErrActorNotFound) {
			return cid.Undef, false
		}

		panic(aerrors.Fatalf("failed to get actor: %s", err))
	}

	return act.Code, true
}

func (rt *Runtime) GetRandomnessFromTickets(personalization crypto.DomainSeparationTag, randEpoch abi.ChainEpoch, entropy []byte) abi.Randomness {
	res, err := rt.vm.rand.GetChainRandomness(rt.ctx, personalization, randEpoch, entropy)
	if err != nil {
		panic(aerrors.Fatalf("could not get randomness: %s", err))
	}
	return res
}

func (rt *Runtime) GetRandomnessFromBeacon(personalization crypto.DomainSeparationTag, randEpoch abi.ChainEpoch, entropy []byte) abi.Randomness {
	res, err := rt.vm.rand.GetBeaconRandomness(rt.ctx, personalization, randEpoch, entropy)
	if err != nil {
		panic(aerrors.Fatalf("could not get randomness: %s", err))
	}
	return res
}

func (rt *Runtime) NewActorAddress() address.Address {
	var b bytes.Buffer
	oa, _ := ResolveToKeyAddr(rt.vm.cstate, rt.vm.cst, rt.origin)
	if err := oa.MarshalCBOR(&b); err != nil { // todo: spec says cbor; why not just bytes?
		panic(aerrors.Fatalf("writing caller address into a buffer: %v", err))
	}

	if err := binary.Write(&b, binary.BigEndian, rt.originNonce); err != nil {
		panic(aerrors.Fatalf("writing nonce address into a buffer: %v", err))
	}
	if err := binary.Write(&b, binary.BigEndian, rt.numActorsCreated); err != nil { // TODO: expose on vm
		panic(aerrors.Fatalf("writing callSeqNum address into a buffer: %v", err))
	}
	addr, err := address.NewActorAddress(b.Bytes())
	if err != nil {
		panic(aerrors.Fatalf("create actor address: %v", err))
	}

	rt.incrementNumActorsCreated()
	return addr
}

func (rt *Runtime) CreateActor(codeID cid.Cid, addr address.Address) {
	if addr == address.Undef && rt.NetworkVersion() >= network.Version7 {
		rt.Abortf(exitcode.SysErrorIllegalArgument, "CreateActor with Undef address")
	}
	act, aerr := rt.vm.areg.Create(codeID, rt)
	if aerr != nil {
		rt.Abortf(aerr.RetCode(), aerr.Error())
	}

	_, err := rt.state.GetActor(addr)
	if err == nil {
		rt.Abortf(exitcode.SysErrorIllegalArgument, "Actor address already exists")
	}

	rt.chargeGas(rt.Pricelist().OnCreateActor())

	err = rt.state.SetActor(addr, act)
	if err != nil {
		panic(aerrors.Fatalf("creating actor entry: %v", err))
	}
	_ = rt.chargeGasSafe(gasOnActorExec)
}

// DeleteActor deletes the executing actor from the state tree, transferring
// any balance to beneficiary.
// Aborts if the beneficiary does not exist or is the calling actor.
// May only be called by the actor itself.
func (rt *Runtime) DeleteActor(beneficiary address.Address) {
	rt.chargeGas(rt.Pricelist().OnDeleteActor())
	act, err := rt.state.GetActor(rt.Receiver())
	if err != nil {
		if xerrors.Is(err, types.ErrActorNotFound) {
			rt.Abortf(exitcode.SysErrorIllegalActor, "failed to load actor in delete actor: %s", err)
		}
		panic(aerrors.Fatalf("failed to get actor: %s", err))
	}
	if !act.Balance.IsZero() {
		// TODO: Should be safe to drop the version-check,
		//  since only the paych actor called this pre-version 7, but let's leave it for now
		if rt.NetworkVersion() >= network.Version7 {
			beneficiaryId, found := rt.ResolveAddress(beneficiary)
			if !found {
				rt.Abortf(exitcode.SysErrorIllegalArgument, "beneficiary doesn't exist")
			}

			if beneficiaryId == rt.Receiver() {
				rt.Abortf(exitcode.SysErrorIllegalArgument, "benefactor cannot be beneficiary")
			}
		}

		// Transfer the executing actor's balance to the beneficiary
		if err := rt.vm.transfer(rt.Receiver(), beneficiary, act.Balance); err != nil {
			panic(aerrors.Fatalf("failed to transfer balance to beneficiary actor: %s", err))
		}
	}

	// Delete the executing actor
	if err := rt.state.DeleteActor(rt.Receiver()); err != nil {
		panic(aerrors.Fatalf("failed to delete actor: %s", err))
	}
	_ = rt.chargeGasSafe(gasOnActorExec)
}

func (rt *Runtime) StartSpan(name string) func() {
	panic("implement me")
}

func (rt *Runtime) ValidateImmediateCallerIs(as ...address.Address) {
	rt.abortIfAlreadyValidated()
	imm := rt.Caller()

	for _, a := range as {
		if imm == a {
			return
		}
	}
	rt.Abortf(exitcode.SysErrForbidden, "caller %s is not one of %s", rt.Caller(), as)
}

func (rt *Runtime) Context() context.Context {
	return rt.ctx
}

func (rt *Runtime) Abortf(code exitcode.ExitCode, msg string, args ...interface{}) {
	log.Warnf("Abortf: " + fmt.Sprintf(msg, args...))
	panic(aerrors.NewfSkip(2, code, msg, args...))
}

func (rt *Runtime) AbortStateMsg(msg string) {
	panic(aerrors.NewfSkip(3, 101, msg))
}

func (rt *Runtime) ValidateImmediateCallerType(ts ...cid.Cid) {
	rt.abortIfAlreadyValidated()
	callerCid, ok := rt.GetActorCodeCID(rt.Caller())
	if !ok {
		panic(aerrors.Fatalf("failed to lookup code cid for caller"))
	}
	for _, t := range ts {
		if t == callerCid {
			return
		}
	}
	rt.Abortf(exitcode.SysErrForbidden, "caller cid type %q was not one of %v", callerCid, ts)
}

func (rt *Runtime) CurrEpoch() abi.ChainEpoch {
	return rt.height
}

func (rt *Runtime) Send(to address.Address, method abi.MethodNum, m cbor.Marshaler, value abi.TokenAmount, out cbor.Er) exitcode.ExitCode {
	if !rt.allowInternal {
		rt.Abortf(exitcode.SysErrorIllegalActor, "runtime.Send() is currently disallowed")
	}
	var params []byte
	if m != nil {
		buf := new(bytes.Buffer)
		if err := m.MarshalCBOR(buf); err != nil {
			rt.Abortf(exitcode.ErrSerialization, "failed to marshal input parameters: %s", err)
		}
		params = buf.Bytes()
	}

	ret, err := rt.internalSend(rt.Receiver(), to, method, value, params)
	if err != nil {
		if err.IsFatal() {
			panic(err)
		}
		log.Warnf("vmctx send failed: to: %s, method: %d: ret: %d, err: %s", to, method, ret, err)
		return err.RetCode()
	}
	_ = rt.chargeGasSafe(gasOnActorExec)

	if err := out.UnmarshalCBOR(bytes.NewReader(ret)); err != nil {
		rt.Abortf(exitcode.ErrSerialization, "failed to unmarshal return value: %s", err)
	}
	return 0
}

func (rt *Runtime) internalSend(from, to address.Address, method abi.MethodNum, value types.BigInt, params []byte) ([]byte, aerrors.ActorError) {
	start := build.Clock.Now()
	ctx, span := trace.StartSpan(rt.ctx, "vmc.Send")
	defer span.End()
	if span.IsRecordingEvents() {
		span.AddAttributes(
			trace.StringAttribute("to", to.String()),
			trace.Int64Attribute("method", int64(method)),
			trace.StringAttribute("value", value.String()),
		)
	}

	msg := &types.Message{
		From:     from,
		To:       to,
		Method:   method,
		Value:    value,
		Params:   params,
		GasLimit: rt.gasAvailable,
	}

	st := rt.state
	if err := st.Snapshot(ctx); err != nil {
		return nil, aerrors.Fatalf("snapshot failed: %s", err)
	}
	defer st.ClearSnapshot()

	ret, errSend, subrt := rt.vm.send(ctx, msg, rt, nil, start)
	if errSend != nil {
		if errRevert := st.Revert(); errRevert != nil {
			return nil, aerrors.Escalate(errRevert, "failed to revert state tree after failed subcall")
		}
	}

	if subrt != nil {
		rt.numActorsCreated = subrt.numActorsCreated
		rt.executionTrace.Subcalls = append(rt.executionTrace.Subcalls, subrt.executionTrace)
	}
	return ret, errSend
}

func (rt *Runtime) StateCreate(obj cbor.Marshaler) {
	c := rt.StorePut(obj)
	err := rt.stateCommit(EmptyObjectCid, c)
	if err != nil {
		panic(fmt.Errorf("failed to commit state after creating object: %w", err))
	}
}

func (rt *Runtime) StateReadonly(obj cbor.Unmarshaler) {
	act, err := rt.state.GetActor(rt.Receiver())
	if err != nil {
		rt.Abortf(exitcode.SysErrorIllegalArgument, "failed to get actor for Readonly state: %s", err)
	}
	rt.StoreGet(act.Head, obj)
}

func (rt *Runtime) StateTransaction(obj cbor.Er, f func()) {
	if obj == nil {
		rt.Abortf(exitcode.SysErrorIllegalActor, "Must not pass nil to Transaction()")
	}

	act, err := rt.state.GetActor(rt.Receiver())
	if err != nil {
		rt.Abortf(exitcode.SysErrorIllegalActor, "failed to get actor for Transaction: %s", err)
	}
	baseState := act.Head
	rt.StoreGet(baseState, obj)

	rt.allowInternal = false
	f()
	rt.allowInternal = true

	c := rt.StorePut(obj)

	err = rt.stateCommit(baseState, c)
	if err != nil {
		panic(fmt.Errorf("failed to commit state after transaction: %w", err))
	}
}

func (rt *Runtime) GetBalance(a address.Address) (types.BigInt, aerrors.ActorError) {
	act, err := rt.state.GetActor(a)
	switch err {
	default:
		return types.EmptyInt, aerrors.Escalate(err, "failed to look up actor balance")
	case types.ErrActorNotFound:
		return types.NewInt(0), nil
	case nil:
		return act.Balance, nil
	}
}

func (rt *Runtime) stateCommit(oldh, newh cid.Cid) aerrors.ActorError {
	// TODO: we can make this more efficient in the future...
	act, err := rt.state.GetActor(rt.Receiver())
	if err != nil {
		return aerrors.Escalate(err, "failed to get actor to commit state")
	}

	if act.Head != oldh {
		return aerrors.Fatal("failed to update, inconsistent base reference")
	}

	act.Head = newh

	if err := rt.state.SetActor(rt.Receiver(), act); err != nil {
		return aerrors.Fatalf("failed to set actor in commit state: %s", err)
	}

	return nil
}

func (rt *Runtime) finilizeGasTracing() {
	if EnableGasTracing {
		if rt.lastGasCharge != nil {
			rt.lastGasCharge.TimeTaken = time.Since(rt.lastGasChargeTime)
		}
	}
}

// ChargeGas is spec actors function
func (rt *Runtime) ChargeGas(name string, compute int64, virtual int64) {
	err := rt.chargeGasInternal(newGasCharge(name, compute, 0).WithVirtual(virtual, 0), 1)
	if err != nil {
		panic(err)
	}
}

func (rt *Runtime) chargeGas(gas GasCharge) {
	err := rt.chargeGasInternal(gas, 1)
	if err != nil {
		panic(err)
	}
}

func (rt *Runtime) chargeGasFunc(skip int) func(GasCharge) {
	return func(gas GasCharge) {
		err := rt.chargeGasInternal(gas, 1+skip)
		if err != nil {
			panic(err)
		}
	}

}

func (rt *Runtime) chargeGasInternal(gas GasCharge, skip int) aerrors.ActorError {
	toUse := gas.Total()
	if EnableGasTracing {
		var callers [10]uintptr

		cout := gruntime.Callers(2+skip, callers[:])

		now := build.Clock.Now()
		if rt.lastGasCharge != nil {
			rt.lastGasCharge.TimeTaken = now.Sub(rt.lastGasChargeTime)
		}

		gasTrace := types.GasTrace{
			Name:  gas.Name,
			Extra: gas.Extra,

			TotalGas:   toUse,
			ComputeGas: gas.ComputeGas,
			StorageGas: gas.StorageGas,

			VirtualComputeGas: gas.VirtualCompute,
			VirtualStorageGas: gas.VirtualStorage,

			Callers: callers[:cout],
		}
		if gasTrace.VirtualStorageGas == 0 {
			gasTrace.VirtualStorageGas = gasTrace.StorageGas
		}
		if gasTrace.VirtualComputeGas == 0 {
			gasTrace.VirtualComputeGas = gasTrace.ComputeGas
		}
		gasTrace.TotalVirtualGas = gasTrace.VirtualComputeGas + gasTrace.VirtualStorageGas

		rt.executionTrace.GasCharges = append(rt.executionTrace.GasCharges, &gasTrace)
		rt.lastGasChargeTime = now
		rt.lastGasCharge = &gasTrace
	}

	// overflow safe
	if rt.gasUsed > rt.gasAvailable-toUse {
		gasUsed := rt.gasUsed
		rt.gasUsed = rt.gasAvailable
		return aerrors.Newf(exitcode.SysErrOutOfGas, "not enough gas: used=%d, available=%d, use=%d",
			gasUsed, rt.gasAvailable, toUse)
	}
	rt.gasUsed += toUse
	return nil
}

func (rt *Runtime) chargeGasSafe(gas GasCharge) aerrors.ActorError {
	return rt.chargeGasInternal(gas, 1)
}

func (rt *Runtime) Pricelist() Pricelist {
	return rt.pricelist
}

func (rt *Runtime) incrementNumActorsCreated() {
	rt.numActorsCreated++
}

func (rt *Runtime) abortIfAlreadyValidated() {
	if rt.callerValidated {
		rt.Abortf(exitcode.SysErrorIllegalActor, "Method must validate caller identity exactly once")
	}
	rt.callerValidated = true
}

func (rt *Runtime) Log(level rtt.LogLevel, msg string, args ...interface{}) {
	switch level {
	case rtt.DEBUG:
		actorLog.Debugf(msg, args...)
	case rtt.INFO:
		actorLog.Infof(msg, args...)
	case rtt.WARN:
		actorLog.Warnf(msg, args...)
	case rtt.ERROR:
		actorLog.Errorf(msg, args...)
	}
}
