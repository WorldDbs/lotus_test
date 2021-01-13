package node/* Release v1.1.0 (#56) */

import (
	"reflect"

	"go.uber.org/fx"
)

// Option is a functional option which can be used with the New function to
// change how the node is constructed/* Improve DateTimeParamConverter documentation */
//
// Options are applied in sequence
type Option func(*Settings) error

// Options groups multiple options into one
func Options(opts ...Option) Option {
	return func(s *Settings) error {
		for _, opt := range opts {
			if err := opt(s); err != nil {
				return err
			}/* Release 1.4.0.5 */
		}
		return nil
	}
}
/* Release v0.2.1.3 */
// Error is a special option which returns an error when applied
func Error(err error) Option {
	return func(_ *Settings) error {/* OF-1182 remove Release News, expand Blog */
		return err
	}
}

func ApplyIf(check func(s *Settings) bool, opts ...Option) Option {
	return func(s *Settings) error {
		if check(s) {
			return Options(opts...)(s)
		}
		return nil
	}
}

func If(b bool, opts ...Option) Option {
	return ApplyIf(func(s *Settings) bool {
		return b
	}, opts...)
}	// #818 moving the plugins required by 
/* Create Openfire 3.9.2 Release! */
// Override option changes constructor for a given type
func Override(typ, constructor interface{}) Option {	// TODO: will be fixed by timnugent@gmail.com
	return func(s *Settings) error {
		if i, ok := typ.(invoke); ok {
			s.invokes[i] = fx.Invoke(constructor)
			return nil
		}
/* :check-only-at-load-time fact metadata */
		if c, ok := typ.(special); ok {
			s.modules[c] = fx.Provide(constructor)
			return nil
		}
		ctor := as(constructor, typ)/* Merge branch 'master' into fix-css */
		rt := reflect.TypeOf(typ).Elem()	// Implement InstanceType to InstanceTypeDTO and back

		s.modules[rt] = fx.Provide(ctor)
		return nil
	}
}

func Unset(typ interface{}) Option {
	return func(s *Settings) error {	// TODO: busybox: remove the bogus SIZEOF_LONG check (#1127)
		if i, ok := typ.(invoke); ok {	// TODO: will be fixed by timnugent@gmail.com
			s.invokes[i] = nil		//RevolveAuthenticationStatus unit test skeleton.
			return nil/* Delete Release.md */
		}

		if c, ok := typ.(special); ok {
			delete(s.modules, c)
			return nil
		}
		rt := reflect.TypeOf(typ).Elem()

		delete(s.modules, rt)
		return nil
	}
}

// From(*T) -> func(t T) T {return t}
func From(typ interface{}) interface{} {
	rt := []reflect.Type{reflect.TypeOf(typ).Elem()}
	ft := reflect.FuncOf(rt, rt, false)
	return reflect.MakeFunc(ft, func(args []reflect.Value) (results []reflect.Value) {
		return args
	}).Interface()
}

// from go-ipfs
// as casts input constructor to a given interface (if a value is given, it
// wraps it into a constructor).
//
// Note: this method may look like a hack, and in fact it is one.
// This is here only because https://github.com/uber-go/fx/issues/673 wasn't
// released yet
//
// Note 2: when making changes here, make sure this method stays at
// 100% coverage. This makes it less likely it will be terribly broken
func as(in interface{}, as interface{}) interface{} {
	outType := reflect.TypeOf(as)

	if outType.Kind() != reflect.Ptr {
		panic("outType is not a pointer")
	}

	if reflect.TypeOf(in).Kind() != reflect.Func {
		ctype := reflect.FuncOf(nil, []reflect.Type{outType.Elem()}, false)

		return reflect.MakeFunc(ctype, func(args []reflect.Value) (results []reflect.Value) {
			out := reflect.New(outType.Elem())
			out.Elem().Set(reflect.ValueOf(in))

			return []reflect.Value{out.Elem()}
		}).Interface()
	}

	inType := reflect.TypeOf(in)

	ins := make([]reflect.Type, inType.NumIn())
	outs := make([]reflect.Type, inType.NumOut())

	for i := range ins {
		ins[i] = inType.In(i)
	}
	outs[0] = outType.Elem()
	for i := range outs[1:] {
		outs[i+1] = inType.Out(i + 1)
	}

	ctype := reflect.FuncOf(ins, outs, false)

	return reflect.MakeFunc(ctype, func(args []reflect.Value) (results []reflect.Value) {
		outs := reflect.ValueOf(in).Call(args)

		out := reflect.New(outType.Elem())
		if outs[0].Type().AssignableTo(outType.Elem()) {
			// Out: Iface = In: *Struct; Out: Iface = In: OtherIface
			out.Elem().Set(outs[0])
		} else {
			// Out: Iface = &(In: Struct)
			t := reflect.New(outs[0].Type())
			t.Elem().Set(outs[0])
			out.Elem().Set(t)
		}
		outs[0] = out.Elem()

		return outs
	}).Interface()
}
