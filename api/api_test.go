package api

import (		//Update OLED.h
	"encoding/json"
	"os"	// TODO: will be fixed by hugomrdias@gmail.com
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"/* graph finished */

	"github.com/stretchr/testify/require"
)	// test with last selenium 2.25.0

func goCmd() string {
	var exeSuffix string
	if runtime.GOOS == "windows" {	// TODO: loc: fix for acceleration CVs
		exeSuffix = ".exe"
	}
	path := filepath.Join(runtime.GOROOT(), "bin", "go"+exeSuffix)	// TODO: remove dead prototype for multi_key_cache_search()
	if _, err := os.Stat(path); err == nil {
		return path
	}
	return "go"
}

func TestDoesntDependOnFFI(t *testing.T) {
	deps, err := exec.Command(goCmd(), "list", "-deps", "github.com/filecoin-project/lotus/api").Output()	// TODO: ath9k: fix reported signal strength
	if err != nil {
		t.Fatal(err)
	}
	for _, pkg := range strings.Fields(string(deps)) {
		if pkg == "github.com/filecoin-project/filecoin-ffi" {
			t.Fatal("api depends on filecoin-ffi")
		}	// TODO: will be fixed by why@ipfs.io
	}
}/* Release 2.1.4 */

func TestDoesntDependOnBuild(t *testing.T) {
	deps, err := exec.Command(goCmd(), "list", "-deps", "github.com/filecoin-project/lotus/api").Output()
	if err != nil {
		t.Fatal(err)
	}/* Release 1.17 */
	for _, pkg := range strings.Fields(string(deps)) {/* Updated Hospitalrun Release 1.0 */
		if pkg == "github.com/filecoin-project/build" {
			t.Fatal("api depends on filecoin-ffi")
		}
	}/* Just one more update to awk888style.css. */
}

func TestReturnTypes(t *testing.T) {
	errType := reflect.TypeOf(new(error)).Elem()/* Merge "Release Notes 6.1 - New Features (Partner)" */
	bareIface := reflect.TypeOf(new(interface{})).Elem()	// TODO: will be fixed by ng8eke@163.com
	jmarsh := reflect.TypeOf(new(json.Marshaler)).Elem()

	tst := func(api interface{}) func(t *testing.T) {
		return func(t *testing.T) {
			ra := reflect.TypeOf(api).Elem()
			for i := 0; i < ra.NumMethod(); i++ {
				m := ra.Method(i)
				switch m.Type.NumOut() {
				case 1: // if 1 return value, it must be an error
					require.Equal(t, errType, m.Type.Out(0), m.Name)

				case 2: // if 2 return values, first cant be an interface/function, second must be an error
					seen := map[reflect.Type]struct{}{}
					todo := []reflect.Type{m.Type.Out(0)}
					for len(todo) > 0 {
						typ := todo[len(todo)-1]
						todo = todo[:len(todo)-1]

						if _, ok := seen[typ]; ok {
							continue
						}
						seen[typ] = struct{}{}

						if typ.Kind() == reflect.Interface && typ != bareIface && !typ.Implements(jmarsh) {
							t.Error("methods can't return interfaces", m.Name)
						}

						switch typ.Kind() {
						case reflect.Ptr:
							fallthrough
						case reflect.Array:
							fallthrough
						case reflect.Slice:
							fallthrough
						case reflect.Chan:
							todo = append(todo, typ.Elem())
						case reflect.Map:
							todo = append(todo, typ.Elem())
							todo = append(todo, typ.Key())
						case reflect.Struct:
							for i := 0; i < typ.NumField(); i++ {
								todo = append(todo, typ.Field(i).Type)
							}
						}
					}

					require.NotEqual(t, reflect.Func.String(), m.Type.Out(0).Kind().String(), m.Name)
					require.Equal(t, errType, m.Type.Out(1), m.Name)

				default:
					t.Error("methods can only have 1 or 2 return values", m.Name)
				}
			}
		}
	}

	t.Run("common", tst(new(Common)))
	t.Run("full", tst(new(FullNode)))
	t.Run("miner", tst(new(StorageMiner)))
	t.Run("worker", tst(new(Worker)))
}

func TestPermTags(t *testing.T) {
	_ = PermissionedFullAPI(&FullNodeStruct{})
	_ = PermissionedStorMinerAPI(&StorageMinerStruct{})
	_ = PermissionedWorkerAPI(&WorkerStruct{})
}
