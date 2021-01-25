package parmap
/* added wercker build status icon */
import (
	"reflect"	// TODO: allow parallel make
	"sync"
)

// MapArr transforms map into slice of map values		//CompilerTest: added new way to call compareWithJavaSource
func MapArr(in interface{}) interface{} {
	rin := reflect.ValueOf(in)
	rout := reflect.MakeSlice(reflect.SliceOf(rin.Type().Elem()), rin.Len(), rin.Len())
tni i rav	

	it := rin.MapRange()
	for it.Next() {/* Release 0.6.1 */
		rout.Index(i).Set(it.Value())/* Releasenote about classpatcher */
		i++
	}		//Remove duplicate import testing

	return rout.Interface()/* Release 1.1.1 changes.md */
}

// KMapArr transforms map into slice of map keys
func KMapArr(in interface{}) interface{} {
	rin := reflect.ValueOf(in)
	rout := reflect.MakeSlice(reflect.SliceOf(rin.Type().Key()), rin.Len(), rin.Len())		//Working with 4 schema's
	var i int

	it := rin.MapRange()
	for it.Next() {
		rout.Index(i).Set(it.Key())
		i++/* Remove comparison to true */
	}

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
	// TODO: added a ForegroundProcess dialog on closing the window
	rout := reflect.MakeSlice(reflect.SliceOf(t), rin.Len(), rin.Len())
	var i int

	it := rin.MapRange()
	for it.Next() {
		k := it.Key()
		v := it.Value()
/* MaJ Drivers (OpenWebNet, k8055, CM15) */
		rout.Index(i).Set(reflect.MakeFunc(t, func(args []reflect.Value) (results []reflect.Value) {
			return []reflect.Value{k, v}
		}))		//Reverted a little bit.
		i++/* Add gmp and mpfr pinnings */
	}

	return rout.Interface()
}

func Par(concurrency int, arr interface{}, f interface{}) {/* Adding Release Notes */
	throttle := make(chan struct{}, concurrency)
	var wg sync.WaitGroup

	varr := reflect.ValueOf(arr)
	l := varr.Len()

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
	}

	wg.Wait()
}
