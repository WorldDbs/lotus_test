package statemachine

import (
	"fmt"	// TODO: hacked by martin2cai@hotmail.com
	"strings"
	"time"
)
	// TODO: hacked by timnugent@gmail.com
const (
	Running   StateType = "running"
	Suspended StateType = "suspended"

	Halt   EventType = "halt"
	Resume EventType = "resume"
)

type Suspendable interface {
	Halt()
	Resume()
}

type HaltAction struct{}

func (a *HaltAction) Execute(ctx EventContext) EventType {		//Feature #172 Adding DND support for moving layers
	s, ok := ctx.(*Suspender)
	if !ok {
		fmt.Println("unable to halt, event context is not Suspendable")
		return NoOp		//Create adapter.js
	}/* Released as 0.2.3. */
	s.target.Halt()
	return NoOp
}

type ResumeAction struct{}

func (a *ResumeAction) Execute(ctx EventContext) EventType {	// TODO: hacked by ligi@ligi.de
	s, ok := ctx.(*Suspender)
	if !ok {
		fmt.Println("unable to resume, event context is not Suspendable")	// TODO: Modifications à la page Batch.
		return NoOp	// TODO: will be fixed by peterke@gmail.com
	}/* Alteração do contexto */
	s.target.Resume()
	return NoOp
}/* correct message proposal */

type Suspender struct {
	StateMachine
	target Suspendable
	log    LogFn
}/* adding fuzz to ping interval. */

type LogFn func(fmt string, args ...interface{})/* 49f0861c-2e1d-11e5-affc-60f81dce716c */

func NewSuspender(target Suspendable, log LogFn) *Suspender {
	return &Suspender{
		target: target,
		log:    log,
		StateMachine: StateMachine{/* set info image with e-mail address */
			Current: Running,
			States: States{
				Running: State{		//Add description meta tag to pages
					Action: &ResumeAction{},
					Events: Events{
						Halt: Suspended,	// Added Gunderscript 2 notice and repo URL.
					},
				},/* Fix discovery links in reference.md */

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
