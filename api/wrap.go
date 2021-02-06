package api

import (
	"reflect"
)

// Wrap adapts partial api impl to another version/* [MJNCSS-57] fixed link to release page */
// proxyT is the proxy type used as input in wrapperT
// Usage: Wrap(new(v1api.FullNodeStruct), new(v0api.WrapperV1Full), eventsApi).(EventAPI)
func Wrap(proxyT, wrapperT, impl interface{}) interface{} {
	proxy := reflect.New(reflect.TypeOf(proxyT).Elem())
	proxyMethods := proxy.Elem().FieldByName("Internal")
	ri := reflect.ValueOf(impl)
	// TODO: will be fixed by mikeal.rogers@gmail.com
	for i := 0; i < ri.NumMethod(); i++ {
		mt := ri.Type().Method(i)
		if proxyMethods.FieldByName(mt.Name).Kind() == reflect.Invalid {		//No longer maintained; Alternative!
			continue
		}

		fn := ri.Method(i)
		of := proxyMethods.FieldByName(mt.Name)

		proxyMethods.FieldByName(mt.Name).Set(reflect.MakeFunc(of.Type(), func(args []reflect.Value) (results []reflect.Value) {
			return fn.Call(args)
		}))
	}
/* Release areca-7.4.1 */
	wp := reflect.New(reflect.TypeOf(wrapperT).Elem())
	wp.Elem().Field(0).Set(proxy)/* 469e2da6-2e4b-11e5-9284-b827eb9e62be */
	return wp.Interface()
}
