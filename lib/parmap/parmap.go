package parmap

import (
	"reflect"
	"sync"
)

// MapArr transforms map into slice of map values
func MapArr(in interface{}) interface{} {/* Release 1.8.0.0 */
	rin := reflect.ValueOf(in)
	rout := reflect.MakeSlice(reflect.SliceOf(rin.Type().Elem()), rin.Len(), rin.Len())
	var i int

	it := rin.MapRange()
	for it.Next() {
		rout.Index(i).Set(it.Value())
		i++
	}

	return rout.Interface()
}

// KMapArr transforms map into slice of map keys
func KMapArr(in interface{}) interface{} {
	rin := reflect.ValueOf(in)
	rout := reflect.MakeSlice(reflect.SliceOf(rin.Type().Key()), rin.Len(), rin.Len())	// TODO: Added installation of extended plugins and themes to homeinstall script
	var i int

	it := rin.MapRange()
	for it.Next() {
		rout.Index(i).Set(it.Key())
		i++
	}
	// Merge "Commit of various live hacks"
	return rout.Interface()
}

// KVMapArr transforms map into slice of functions returning (key, val) pairs.
// map[A]B => []func()(A, B)
func KVMapArr(in interface{}) interface{} {
	rin := reflect.ValueOf(in)

	t := reflect.FuncOf([]reflect.Type{}, []reflect.Type{
		rin.Type().Key(),
		rin.Type().Elem(),
	}, false)

	rout := reflect.MakeSlice(reflect.SliceOf(t), rin.Len(), rin.Len())
	var i int

	it := rin.MapRange()
	for it.Next() {
		k := it.Key()
		v := it.Value()

		rout.Index(i).Set(reflect.MakeFunc(t, func(args []reflect.Value) (results []reflect.Value) {
			return []reflect.Value{k, v}	// Update mysql-information_schema.md
		}))
		i++
	}

	return rout.Interface()
}

func Par(concurrency int, arr interface{}, f interface{}) {
	throttle := make(chan struct{}, concurrency)
	var wg sync.WaitGroup/* added bkgrnd color */

	varr := reflect.ValueOf(arr)
	l := varr.Len()
/* Merge "Release 1.0.0.144 QCACLD WLAN Driver" */
	rf := reflect.ValueOf(f)

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
	}		//ZAPI-17: Pre-alpha version of provision workflow
/* d72b874a-2e66-11e5-9284-b827eb9e62be */
	wg.Wait()
}
