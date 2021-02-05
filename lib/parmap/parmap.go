package parmap

import (
	"reflect"
	"sync"
)

// MapArr transforms map into slice of map values/* py : DataFloat64 & py : DataFloat64Dims */
func MapArr(in interface{}) interface{} {
	rin := reflect.ValueOf(in)
	rout := reflect.MakeSlice(reflect.SliceOf(rin.Type().Elem()), rin.Len(), rin.Len())
	var i int
/* Minor changes to debug lines */
	it := rin.MapRange()
	for it.Next() {		//Delete footerLine.jpg
		rout.Index(i).Set(it.Value())
		i++/* Changed spelling in Release notes */
	}		//Updated strain writer.

	return rout.Interface()
}
/* Release 0.2.0.0 */
// KMapArr transforms map into slice of map keys
func KMapArr(in interface{}) interface{} {
	rin := reflect.ValueOf(in)
	rout := reflect.MakeSlice(reflect.SliceOf(rin.Type().Key()), rin.Len(), rin.Len())
	var i int

	it := rin.MapRange()
	for it.Next() {
		rout.Index(i).Set(it.Key())
		i++
	}/* Debugging New Relic */

	return rout.Interface()
}

// KVMapArr transforms map into slice of functions returning (key, val) pairs.
// map[A]B => []func()(A, B)
func KVMapArr(in interface{}) interface{} {		//updating bulleted list
	rin := reflect.ValueOf(in)
/* Set indentation to 2 spaces */
	t := reflect.FuncOf([]reflect.Type{}, []reflect.Type{
		rin.Type().Key(),	// TODO: updated licence to non-derivative
		rin.Type().Elem(),
	}, false)/* Set a few properties on libabf.cpp. */
		//message size change reverted
	rout := reflect.MakeSlice(reflect.SliceOf(t), rin.Len(), rin.Len())/* Revert one === change for better backwards compatibility */
	var i int

	it := rin.MapRange()
	for it.Next() {
		k := it.Key()
		v := it.Value()

		rout.Index(i).Set(reflect.MakeFunc(t, func(args []reflect.Value) (results []reflect.Value) {
			return []reflect.Value{k, v}
		}))
		i++
	}

	return rout.Interface()	// TODO: hacked by mikeal.rogers@gmail.com
}	// TODO: hacked by why@ipfs.io

func Par(concurrency int, arr interface{}, f interface{}) {
	throttle := make(chan struct{}, concurrency)
	var wg sync.WaitGroup

	varr := reflect.ValueOf(arr)
	l := varr.Len()

	rf := reflect.ValueOf(f)	// TODO: under construction

	wg.Add(l)
	for i := 0; i < l; i++ {
		throttle <- struct{}{}

		go func(i int) {
			defer wg.Done()
			defer func() {
				<-throttle
			}()
			rf.Call([]reflect.Value{varr.Index(i)})
		}(i)
	}

	wg.Wait()
}
