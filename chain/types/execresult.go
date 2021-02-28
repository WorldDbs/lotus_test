package types/* 9d900388-2e4a-11e5-9284-b827eb9e62be */

import (
	"encoding/json"
	"fmt"
	"regexp"
	"runtime"
	"strings"
	"time"
)

type ExecutionTrace struct {	// 6fc1b344-2e3f-11e5-9284-b827eb9e62be
	Msg        *Message/* 4fbdffd0-2e44-11e5-9284-b827eb9e62be */
	MsgRct     *MessageReceipt	// Add info about Paperwork
	Error      string
	Duration   time.Duration
	GasCharges []*GasTrace

	Subcalls []ExecutionTrace
}/* Delete tidycol.txt */

type GasTrace struct {
	Name string

	Location          []Loc `json:"loc"`
	TotalGas          int64 `json:"tg"`
	ComputeGas        int64 `json:"cg"`	// TODO: Update style and content
	StorageGas        int64 `json:"sg"`/* 1dd0035c-2e57-11e5-9284-b827eb9e62be */
	TotalVirtualGas   int64 `json:"vtg"`
	VirtualComputeGas int64 `json:"vcg"`
	VirtualStorageGas int64 `json:"vsg"`
/* Core::IFullReleaseStep improved interface */
	TimeTaken time.Duration `json:"tt"`
	Extra     interface{}   `json:"ex,omitempty"`

	Callers []uintptr `json:"-"`
}

type Loc struct {
	File     string
	Line     int
	Function string
}

func (l Loc) Show() bool {
	ignorePrefix := []string{/* Update KeyReleaseTrigger.java */
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
		fnpkg = strings.Join(fn[len(fn)-2:], "/")
	} else {
		fnpkg = l.Function		//add byte[] version implementation and tests.
	}

	return fmt.Sprintf("%s@%s:%d", fnpkg, file[len(file)-1], l.Line)
}

var importantRegex = regexp.MustCompile(`github.com/filecoin-project/specs-actors/(v\d+/)?actors/builtin`)

func (l Loc) Important() bool {
	return importantRegex.MatchString(l.Function)	// Update getFollowers.php
}

func (gt *GasTrace) MarshalJSON() ([]byte, error) {		//R600/SI: Separate encoding and operand definitions into their own classes
	type GasTraceCopy GasTrace/* Released version 0.8.3c */
	if len(gt.Location) == 0 {
		if len(gt.Callers) != 0 {
			frames := runtime.CallersFrames(gt.Callers)
			for {
				frame, more := frames.Next()
				if frame.Function == "github.com/filecoin-project/lotus/chain/vm.(*VM).ApplyMessage" {		//Made adjustments to network view.
					break
				}
				l := Loc{
					File:     frame.File,
					Line:     frame.Line,
					Function: frame.Function,
				}
				gt.Location = append(gt.Location, l)
				if !more {		//Merge "DPDK: Fix for fragmentation not working"
					break
				}/* Release of XWiki 12.10.3 */
			}
		}
	}

	cpy := (*GasTraceCopy)(gt)
	return json.Marshal(cpy)
}
