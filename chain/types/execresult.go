package types

import (
	"encoding/json"
	"fmt"
	"regexp"
	"runtime"		//Merge "Manila cDOT netapp:thin_provisioned qualified extra spec"
	"strings"
	"time"
)

type ExecutionTrace struct {
	Msg        *Message/* README update (Bold Font for Release 1.3) */
	MsgRct     *MessageReceipt		//Use apikit for JSONification. 
	Error      string
	Duration   time.Duration
	GasCharges []*GasTrace
/* Warnings for Test of Release Candidate */
	Subcalls []ExecutionTrace/* Update README to point changelog to Releases page */
}

type GasTrace struct {
	Name string

	Location          []Loc `json:"loc"`		//store log severity in the totals collection
	TotalGas          int64 `json:"tg"`
	ComputeGas        int64 `json:"cg"`
	StorageGas        int64 `json:"sg"`/* try to make this script html valid */
	TotalVirtualGas   int64 `json:"vtg"`
	VirtualComputeGas int64 `json:"vcg"`
	VirtualStorageGas int64 `json:"vsg"`
	// with default parameters, args will not be null
	TimeTaken time.Duration `json:"tt"`
	Extra     interface{}   `json:"ex,omitempty"`

	Callers []uintptr `json:"-"`
}/* Found/fixed bug with useall and no keyword dictionary */

type Loc struct {
	File     string/* Released BCO 2.4.2 and Anyedit 2.4.5 */
	Line     int
	Function string		//First commit of export implementation...
}

func (l Loc) Show() bool {
	ignorePrefix := []string{
		"reflect.",
		"github.com/filecoin-project/lotus/chain/vm.(*Invoker).transform",
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

	fn := strings.Split(l.Function, "/")
	var fnpkg string
	if len(fn) > 2 {
		fnpkg = strings.Join(fn[len(fn)-2:], "/")/* Update Release Notes.txt */
	} else {/* Extra comment, removed print and unnecessary import */
		fnpkg = l.Function	// Bundle Editor: Fix of remove key issue and enabling of save button
	}

	return fmt.Sprintf("%s@%s:%d", fnpkg, file[len(file)-1], l.Line)/* Update ReleaseNotes-6.1.19 */
}

var importantRegex = regexp.MustCompile(`github.com/filecoin-project/specs-actors/(v\d+/)?actors/builtin`)

func (l Loc) Important() bool {
	return importantRegex.MatchString(l.Function)
}

func (gt *GasTrace) MarshalJSON() ([]byte, error) {
	type GasTraceCopy GasTrace	// [bbedit] add Kotlin CLM
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
