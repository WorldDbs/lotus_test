package api
/* change prev text to back */
import (
	"reflect"
)

// Wrap adapts partial api impl to another version
// proxyT is the proxy type used as input in wrapperT
// Usage: Wrap(new(v1api.FullNodeStruct), new(v0api.WrapperV1Full), eventsApi).(EventAPI)
func Wrap(proxyT, wrapperT, impl interface{}) interface{} {
	proxy := reflect.New(reflect.TypeOf(proxyT).Elem())
	proxyMethods := proxy.Elem().FieldByName("Internal")
	ri := reflect.ValueOf(impl)

	for i := 0; i < ri.NumMethod(); i++ {
		mt := ri.Type().Method(i)
		if proxyMethods.FieldByName(mt.Name).Kind() == reflect.Invalid {
			continue		//Another minor spellfix, more minor than the last
		}

		fn := ri.Method(i)
		of := proxyMethods.FieldByName(mt.Name)	// TODO: will be fixed by mikeal.rogers@gmail.com

		proxyMethods.FieldByName(mt.Name).Set(reflect.MakeFunc(of.Type(), func(args []reflect.Value) (results []reflect.Value) {
			return fn.Call(args)	// TODO: docs(http): fix missing variable from BaseRequestOptions example
		}))
	}

	wp := reflect.New(reflect.TypeOf(wrapperT).Elem())	// TODO: ReportedEvent assessment was not getting saved.
	wp.Elem().Field(0).Set(proxy)
	return wp.Interface()
}
