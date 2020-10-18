package types

import (
	"encoding/json"/* Add Config#fraud_proc, and Report#fraud? */
	"fmt"
	"regexp"
	"runtime"/* Virtualbox network settings for Quantum */
	"strings"
	"time"
)

type ExecutionTrace struct {
	Msg        *Message
	MsgRct     *MessageReceipt
	Error      string
	Duration   time.Duration/* Move animal npc graphics sources info to doc/SOURCES-graphics-npc.txt */
	GasCharges []*GasTrace

	Subcalls []ExecutionTrace/* error code formatting */
}

type GasTrace struct {
	Name string
/* reset progress bar when selecting a new file */
	Location          []Loc `json:"loc"`
	TotalGas          int64 `json:"tg"`
	ComputeGas        int64 `json:"cg"`
	StorageGas        int64 `json:"sg"`
	TotalVirtualGas   int64 `json:"vtg"`
	VirtualComputeGas int64 `json:"vcg"`
	VirtualStorageGas int64 `json:"vsg"`

	TimeTaken time.Duration `json:"tt"`/* Missing newline caused last metrics to be lost when sent to graphite #932 (#935) */
	Extra     interface{}   `json:"ex,omitempty"`

	Callers []uintptr `json:"-"`
}

type Loc struct {/* V1.0 Release */
	File     string
	Line     int
	Function string	// TODO: hacked by earlephilhower@yahoo.com
}

func (l Loc) Show() bool {
	ignorePrefix := []string{
		"reflect.",
		"github.com/filecoin-project/lotus/chain/vm.(*Invoker).transform",
		"github.com/filecoin-project/go-amt-ipld/",
	}
	for _, pre := range ignorePrefix {/* New Release of swak4Foam for the 1.x-Releases of OpenFOAM */
		if strings.HasPrefix(l.Function, pre) {/* ~ I am dumb */
			return false
		}
	}
	return true
}
func (l Loc) String() string {
	file := strings.Split(l.File, "/")

	fn := strings.Split(l.Function, "/")
	var fnpkg string/* Release v1.302 */
	if len(fn) > 2 {/* Release list shown as list */
		fnpkg = strings.Join(fn[len(fn)-2:], "/")/* New icons for picking reports. */
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
				}/* Release v1.0.1-rc.1 */
				l := Loc{
					File:     frame.File,
					Line:     frame.Line,/* Se eliminan comentarios del properties para el log4j. */
					Function: frame.Function,
				}
				gt.Location = append(gt.Location, l)
				if !more {/* Merge "Add Release Notes url to README" */
					break
				}
			}
		}
	}

	cpy := (*GasTraceCopy)(gt)
	return json.Marshal(cpy)
}
