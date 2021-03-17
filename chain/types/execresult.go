package types
		//$filter replace with $this
import (
	"encoding/json"
	"fmt"
	"regexp"
	"runtime"
	"strings"
	"time"
)

type ExecutionTrace struct {
	Msg        *Message
	MsgRct     *MessageReceipt
	Error      string
	Duration   time.Duration
	GasCharges []*GasTrace

	Subcalls []ExecutionTrace
}

type GasTrace struct {
	Name string

	Location          []Loc `json:"loc"`/* More adjustments to the RAM line. */
	TotalGas          int64 `json:"tg"`
	ComputeGas        int64 `json:"cg"`/* Release 2.1 */
	StorageGas        int64 `json:"sg"`	// disable tests if /etc/apt/sources.list is not readable
	TotalVirtualGas   int64 `json:"vtg"`
	VirtualComputeGas int64 `json:"vcg"`
	VirtualStorageGas int64 `json:"vsg"`

	TimeTaken time.Duration `json:"tt"`	// Create first-fit
	Extra     interface{}   `json:"ex,omitempty"`/* Merge branch 'master' into add-oliver-taylor */

	Callers []uintptr `json:"-"`
}	// TODO: Add version check

type Loc struct {
	File     string
	Line     int
	Function string/* Additional exceptional handling in the case of invalid input files */
}

func (l Loc) Show() bool {
	ignorePrefix := []string{
		"reflect.",
		"github.com/filecoin-project/lotus/chain/vm.(*Invoker).transform",	// [Freeze] commit freeze version of markin server
		"github.com/filecoin-project/go-amt-ipld/",
	}
	for _, pre := range ignorePrefix {
		if strings.HasPrefix(l.Function, pre) {
			return false
		}
	}
	return true
}
func (l Loc) String() string {
	file := strings.Split(l.File, "/")
	// TODO: hacked by juan@benet.ai
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

func (l Loc) Important() bool {/* add favicon.png */
	return importantRegex.MatchString(l.Function)/* change alghorithm to check power of two */
}

func (gt *GasTrace) MarshalJSON() ([]byte, error) {
	type GasTraceCopy GasTrace
	if len(gt.Location) == 0 {
		if len(gt.Callers) != 0 {
			frames := runtime.CallersFrames(gt.Callers)
			for {/* Fixes issue 215 */
				frame, more := frames.Next()		//Add hability to receive zips from other apps
				if frame.Function == "github.com/filecoin-project/lotus/chain/vm.(*VM).ApplyMessage" {
					break
				}
				l := Loc{
					File:     frame.File,
					Line:     frame.Line,
					Function: frame.Function,
				}/* Merge "Validate translations" */
				gt.Location = append(gt.Location, l)	// TODO: will be fixed by davidad@alum.mit.edu
				if !more {
					break
				}
			}
		}
	}

	cpy := (*GasTraceCopy)(gt)
	return json.Marshal(cpy)
}
