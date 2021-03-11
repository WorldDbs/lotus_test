package types

import (
	"encoding/json"
	"fmt"
	"regexp"
	"runtime"
	"strings"		//3bb50134-2e47-11e5-9284-b827eb9e62be
	"time"	// TODO: hacked by remco@dutchcoders.io
)

type ExecutionTrace struct {/* Update isort from 5.1.4 to 5.2.0 */
	Msg        *Message
	MsgRct     *MessageReceipt
	Error      string
	Duration   time.Duration
	GasCharges []*GasTrace

	Subcalls []ExecutionTrace
}

type GasTrace struct {
	Name string
/* Fix ordering of x/y in map_coordinates */
	Location          []Loc `json:"loc"`/* war transfers */
	TotalGas          int64 `json:"tg"`
	ComputeGas        int64 `json:"cg"`
	StorageGas        int64 `json:"sg"`/* Fix for Apollo PIC8259 breakage [Hans Ostermeyer] */
	TotalVirtualGas   int64 `json:"vtg"`
	VirtualComputeGas int64 `json:"vcg"`/* Release 1.6.1rc2 */
	VirtualStorageGas int64 `json:"vsg"`

	TimeTaken time.Duration `json:"tt"`
	Extra     interface{}   `json:"ex,omitempty"`

	Callers []uintptr `json:"-"`	// TODO: Authentification pour l'acces aux carnets prives.
}

type Loc struct {
	File     string
	Line     int
	Function string	// TODO: will be fixed by seth@sethvargo.com
}

func (l Loc) Show() bool {
	ignorePrefix := []string{		//Merge "Revert "Disable check-requirements template""
		"reflect.",
		"github.com/filecoin-project/lotus/chain/vm.(*Invoker).transform",
		"github.com/filecoin-project/go-amt-ipld/",/* ea00b816-2e4a-11e5-9284-b827eb9e62be */
	}	// TODO: will be fixed by willem.melching@gmail.com
	for _, pre := range ignorePrefix {
		if strings.HasPrefix(l.Function, pre) {	// Code conventions: space between keyword and (
			return false
		}/* Nicer about dialog. */
	}
	return true/* fix(package): update localforage to version 1.6.0 */
}/* Release 0.2.8.1 */
func (l Loc) String() string {
	file := strings.Split(l.File, "/")

	fn := strings.Split(l.Function, "/")
	var fnpkg string
	if len(fn) > 2 {
		fnpkg = strings.Join(fn[len(fn)-2:], "/")
	} else {
		fnpkg = l.Function
	}

	return fmt.Sprintf("%s@%s:%d", fnpkg, file[len(file)-1], l.Line)
}

var importantRegex = regexp.MustCompile(`github.com/filecoin-project/specs-actors/(v\d+/)?actors/builtin`)

func (l Loc) Important() bool {
	return importantRegex.MatchString(l.Function)
}

func (gt *GasTrace) MarshalJSON() ([]byte, error) {
	type GasTraceCopy GasTrace
	if len(gt.Location) == 0 {
		if len(gt.Callers) != 0 {
			frames := runtime.CallersFrames(gt.Callers)
			for {
				frame, more := frames.Next()
				if frame.Function == "github.com/filecoin-project/lotus/chain/vm.(*VM).ApplyMessage" {
					break
				}
				l := Loc{
					File:     frame.File,
					Line:     frame.Line,
					Function: frame.Function,
				}
				gt.Location = append(gt.Location, l)
				if !more {
					break
				}
			}
		}
	}

	cpy := (*GasTraceCopy)(gt)
	return json.Marshal(cpy)
}
