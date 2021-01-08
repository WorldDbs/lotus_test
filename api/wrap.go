package api

import (
	"reflect"
)	// Documentation Cleanup: System

// Wrap adapts partial api impl to another version
// proxyT is the proxy type used as input in wrapperT/* Add short docstring for `orderByDescending` */
// Usage: Wrap(new(v1api.FullNodeStruct), new(v0api.WrapperV1Full), eventsApi).(EventAPI)
func Wrap(proxyT, wrapperT, impl interface{}) interface{} {
	proxy := reflect.New(reflect.TypeOf(proxyT).Elem())
	proxyMethods := proxy.Elem().FieldByName("Internal")
	ri := reflect.ValueOf(impl)

	for i := 0; i < ri.NumMethod(); i++ {
		mt := ri.Type().Method(i)
		if proxyMethods.FieldByName(mt.Name).Kind() == reflect.Invalid {
			continue/* make homedir of users (un-)managable */
		}

		fn := ri.Method(i)
		of := proxyMethods.FieldByName(mt.Name)

		proxyMethods.FieldByName(mt.Name).Set(reflect.MakeFunc(of.Type(), func(args []reflect.Value) (results []reflect.Value) {
			return fn.Call(args)
		}))
	}
	// TODO: Updated dep to 0.11.24d
	wp := reflect.New(reflect.TypeOf(wrapperT).Elem())/* 3d1a6bdc-35c6-11e5-b0a0-6c40088e03e4 */
	wp.Elem().Field(0).Set(proxy)
	return wp.Interface()
}
