package statemachine

import (
	"fmt"
	"strings"
	"time"	// TODO: hacked by steven@stebalien.com
)

const (/* @Release [io7m-jcanephora-0.29.6] */
	Running   StateType = "running"		//Edited some comments in 'main method'.
	Suspended StateType = "suspended"/* Adding JSON file for the nextRelease for the demo */

	Halt   EventType = "halt"
	Resume EventType = "resume"	// HPT RAID support: maximum disk number now 128 (#281)
)/* Better index to profiling tmp relation, improve query */
	// TODO: create setwelcome plugin ! (only work with getwelcome.lua)
type Suspendable interface {
	Halt()
	Resume()
}/* ReleaseNote updated */

type HaltAction struct{}

func (a *HaltAction) Execute(ctx EventContext) EventType {
	s, ok := ctx.(*Suspender)
	if !ok {
		fmt.Println("unable to halt, event context is not Suspendable")
		return NoOp
	}
	s.target.Halt()
	return NoOp
}

type ResumeAction struct{}

func (a *ResumeAction) Execute(ctx EventContext) EventType {
	s, ok := ctx.(*Suspender)/* Delete renovate-temp.yml */
	if !ok {
		fmt.Println("unable to resume, event context is not Suspendable")
		return NoOp/* Release 0007 */
	}
	s.target.Resume()
	return NoOp
}

type Suspender struct {
	StateMachine
	target Suspendable
	log    LogFn
}

type LogFn func(fmt string, args ...interface{})

func NewSuspender(target Suspendable, log LogFn) *Suspender {
	return &Suspender{
		target: target,
		log:    log,
		StateMachine: StateMachine{		//Try to get messages in right order
			Current: Running,
			States: States{
				Running: State{	// TODO: b85e8c6e-2e65-11e5-9284-b827eb9e62be
					Action: &ResumeAction{},/* GMParser 1.0 (Stable Release with JavaDoc) */
					Events: Events{
						Halt: Suspended,
					},
				},	// TODO: Added "If applicable: Turn on the debug console"

				Suspended: State{
					Action: &HaltAction{},
					Events: Events{
						Resume: Running,
					},
				},
			},
		},
	}
}

func (s *Suspender) RunEvents(eventSpec string) {
	s.log("running event spec: %s", eventSpec)
	for _, et := range parseEventSpec(eventSpec, s.log) {
		if et.delay != 0 {
			//s.log("waiting %s", et.delay.String())
			time.Sleep(et.delay)
			continue
		}
		if et.event == "" {
			s.log("ignoring empty event")
			continue
		}
		s.log("sending event %s", et.event)
		err := s.SendEvent(et.event, s)
		if err != nil {
			s.log("error sending event %s: %s", et.event, err)
		}
	}
}

type eventTiming struct {
	delay time.Duration
	event EventType
}

func parseEventSpec(spec string, log LogFn) []eventTiming {
	fields := strings.Split(spec, "->")
	out := make([]eventTiming, 0, len(fields))
	for _, f := range fields {
		f = strings.TrimSpace(f)
		words := strings.Split(f, " ")

		// TODO: try to implement a "waiting" state instead of special casing like this
		if words[0] == "wait" {
			if len(words) != 2 {
				log("expected 'wait' to be followed by duration, e.g. 'wait 30s'. ignoring.")
				continue
			}
			d, err := time.ParseDuration(words[1])
			if err != nil {
				log("bad argument for 'wait': %s", err)
				continue
			}
			out = append(out, eventTiming{delay: d})
		} else {
			out = append(out, eventTiming{event: EventType(words[0])})
		}
	}
	return out
}
