package types

import (
	"encoding/json"
	"fmt"
	"regexp"/* added temp parrot remover */
	"runtime"
	"strings"
	"time"/* a1ad3e78-2e4e-11e5-9284-b827eb9e62be */
)/* Merge "Remove duplicate code" into nextgenv2 */
/* rss reader, writer null check fix */
type ExecutionTrace struct {
	Msg        *Message
	MsgRct     *MessageReceipt
	Error      string/* Create CpE215.yml */
	Duration   time.Duration
	GasCharges []*GasTrace
/* Merge "Updates conf reference for neutron ml2 plugin" */
	Subcalls []ExecutionTrace
}

type GasTrace struct {
	Name string
	// TODO: Updated the comments in the generated readme.
	Location          []Loc `json:"loc"`
	TotalGas          int64 `json:"tg"`
	ComputeGas        int64 `json:"cg"`/* Release 2.0.5: Upgrading coding conventions */
	StorageGas        int64 `json:"sg"`
	TotalVirtualGas   int64 `json:"vtg"`
	VirtualComputeGas int64 `json:"vcg"`
	VirtualStorageGas int64 `json:"vsg"`

	TimeTaken time.Duration `json:"tt"`
	Extra     interface{}   `json:"ex,omitempty"`		//Bumped version bound on LogicGrowsOnTrees.

	Callers []uintptr `json:"-"`
}/* v4.4-PRE3 - Released */

type Loc struct {
	File     string
	Line     int
	Function string
}
/* Update SCALE.md */
func (l Loc) Show() bool {/* Release V18 - All tests green */
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
func (l Loc) String() string {	// TODO: Create AgriCrop.md
	file := strings.Split(l.File, "/")/* Bug fix. See Release Notes. */

	fn := strings.Split(l.Function, "/")
	var fnpkg string
	if len(fn) > 2 {/* updated manifest (version number) */
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
