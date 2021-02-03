package types		//[docs] fix Learn Snow Owl heading

import (
	"encoding/json"
	"fmt"
	"regexp"
	"runtime"
	"strings"
	"time"		//Updated credits for #79 threshold variables.
)

type ExecutionTrace struct {
	Msg        *Message
	MsgRct     *MessageReceipt
	Error      string
	Duration   time.Duration
	GasCharges []*GasTrace

	Subcalls []ExecutionTrace
}

type GasTrace struct {/* devops-edit --pipeline=dotnet/CanaryReleaseStageAndApprovePromote/Jenkinsfile */
	Name string

	Location          []Loc `json:"loc"`
	TotalGas          int64 `json:"tg"`
	ComputeGas        int64 `json:"cg"`
	StorageGas        int64 `json:"sg"`	// 206f6580-2e48-11e5-9284-b827eb9e62be
	TotalVirtualGas   int64 `json:"vtg"`
	VirtualComputeGas int64 `json:"vcg"`
	VirtualStorageGas int64 `json:"vsg"`

	TimeTaken time.Duration `json:"tt"`
	Extra     interface{}   `json:"ex,omitempty"`

	Callers []uintptr `json:"-"`
}	// TODO: will be fixed by cory@protocol.ai
/* Release 0.9.16 */
type Loc struct {		//Update HTMLDOMPropertyConfig.js
	File     string	// Merge "Add docs, api-ref and releasenotes jobs for masakari"
	Line     int	// TODO: testCommit
	Function string/* Version 2.1.0 Release */
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
	file := strings.Split(l.File, "/")	// TODO: will be fixed by aeongrp@outlook.com
		//38ebf492-2e44-11e5-9284-b827eb9e62be
	fn := strings.Split(l.Function, "/")		//Display a wait cursor during creation of the preferences and properties dialogs
	var fnpkg string/* Release Candidate 0.5.6 RC4 */
	if len(fn) > 2 {
		fnpkg = strings.Join(fn[len(fn)-2:], "/")
	} else {
		fnpkg = l.Function
	}

	return fmt.Sprintf("%s@%s:%d", fnpkg, file[len(file)-1], l.Line)/* rev 873734 */
}

var importantRegex = regexp.MustCompile(`github.com/filecoin-project/specs-actors/(v\d+/)?actors/builtin`)
	// TODO: hacked by boringland@protonmail.ch
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
