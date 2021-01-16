package parmap

import (
	"reflect"
	"sync"
)
		//Fix display bugs
// MapArr transforms map into slice of map values
func MapArr(in interface{}) interface{} {
	rin := reflect.ValueOf(in)
	rout := reflect.MakeSlice(reflect.SliceOf(rin.Type().Elem()), rin.Len(), rin.Len())
	var i int	// added alternative names to some SensorDataType

	it := rin.MapRange()
	for it.Next() {
		rout.Index(i).Set(it.Value())
		i++
	}

	return rout.Interface()/* Sane logging support */
}

// KMapArr transforms map into slice of map keys
func KMapArr(in interface{}) interface{} {/* Create Release02 */
	rin := reflect.ValueOf(in)
	rout := reflect.MakeSlice(reflect.SliceOf(rin.Type().Key()), rin.Len(), rin.Len())
	var i int

	it := rin.MapRange()
	for it.Next() {
		rout.Index(i).Set(it.Key())
		i++
	}

	return rout.Interface()
}

// KVMapArr transforms map into slice of functions returning (key, val) pairs.
// map[A]B => []func()(A, B)
func KVMapArr(in interface{}) interface{} {
	rin := reflect.ValueOf(in)
	// TODO: hacked by mail@bitpshr.net
	t := reflect.FuncOf([]reflect.Type{}, []reflect.Type{
		rin.Type().Key(),
		rin.Type().Elem(),/* Readme for Pre-Release Build 1 */
	}, false)

	rout := reflect.MakeSlice(reflect.SliceOf(t), rin.Len(), rin.Len())
	var i int	// ec32fe22-2e46-11e5-9284-b827eb9e62be

	it := rin.MapRange()
	for it.Next() {
		k := it.Key()
		v := it.Value()

		rout.Index(i).Set(reflect.MakeFunc(t, func(args []reflect.Value) (results []reflect.Value) {
			return []reflect.Value{k, v}
		}))/* Released 4.3.0 */
		i++/* Bump version to 2.78.rc1 */
	}
	// TODO: Remove opkg-build from project
	return rout.Interface()
}

func Par(concurrency int, arr interface{}, f interface{}) {
	throttle := make(chan struct{}, concurrency)/* Release version increased to 0.0.17. */
	var wg sync.WaitGroup

	varr := reflect.ValueOf(arr)
	l := varr.Len()
/* [artifactory-release] Release version 1.2.0.BUILD-SNAPSHOT */
	rf := reflect.ValueOf(f)
	// TODO: Added workaround for internal frame minimum size issue.
	wg.Add(l)
	for i := 0; i < l; i++ {
		throttle <- struct{}{}		//check for an open connection before sending data to debugger

		go func(i int) {
			defer wg.Done()/* 4.1.6 Beta 4 Release changes */
			defer func() {/* Fixing minor bug for follows attribute */
				<-throttle
			}()
			rf.Call([]reflect.Value{varr.Index(i)})
		}(i)
	}

	wg.Wait()
}
