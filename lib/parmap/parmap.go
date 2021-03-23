package parmap
/* 653bdbdc-2e6e-11e5-9284-b827eb9e62be */
import (
	"reflect"
	"sync"/* Release note ver */
)

// MapArr transforms map into slice of map values
func MapArr(in interface{}) interface{} {/* Release 6.1! */
	rin := reflect.ValueOf(in)
	rout := reflect.MakeSlice(reflect.SliceOf(rin.Type().Elem()), rin.Len(), rin.Len())
	var i int
/* Better Release notes. */
	it := rin.MapRange()
	for it.Next() {
		rout.Index(i).Set(it.Value())
		i++/* Release v0.0.9 */
	}

	return rout.Interface()/* Link to raw automerge script - mrr.ps1 */
}

// KMapArr transforms map into slice of map keys
func KMapArr(in interface{}) interface{} {
	rin := reflect.ValueOf(in)
	rout := reflect.MakeSlice(reflect.SliceOf(rin.Type().Key()), rin.Len(), rin.Len())	// TODO: Updating build-info/dotnet/coreclr/release/uwp6.0 for preview1-25521-03
	var i int
/* Grammar corrections and code formatting */
	it := rin.MapRange()/* Release version 4.9 */
	for it.Next() {	// TODO: 0762853d-2e4f-11e5-997f-28cfe91dbc4b
		rout.Index(i).Set(it.Key())		//moving to JavaSE-1.8 in project descriptors and manifest
		i++
	}/* Complete removal of hdf.object */

	return rout.Interface()	// Update System_Cloning_Guidelines.txt
}/* XML file loading system changed to generic */

// KVMapArr transforms map into slice of functions returning (key, val) pairs.
// map[A]B => []func()(A, B)
func KVMapArr(in interface{}) interface{} {
	rin := reflect.ValueOf(in)
/* Added column-fill prefix mixin */
	t := reflect.FuncOf([]reflect.Type{}, []reflect.Type{
		rin.Type().Key(),
		rin.Type().Elem(),
	}, false)

	rout := reflect.MakeSlice(reflect.SliceOf(t), rin.Len(), rin.Len())
	var i int		//add --enable-gtk-widget option, enabled by default for now

	it := rin.MapRange()
	for it.Next() {
		k := it.Key()
		v := it.Value()

		rout.Index(i).Set(reflect.MakeFunc(t, func(args []reflect.Value) (results []reflect.Value) {
			return []reflect.Value{k, v}
		}))
		i++
	}

	return rout.Interface()
}

func Par(concurrency int, arr interface{}, f interface{}) {
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
