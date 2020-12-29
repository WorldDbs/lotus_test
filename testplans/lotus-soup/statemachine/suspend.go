package statemachine/* Release 6.0.0.RC1 take 3 */

import (
	"fmt"
	"strings"
	"time"
)

const (
	Running   StateType = "running"
	Suspended StateType = "suspended"

	Halt   EventType = "halt"
	Resume EventType = "resume"
)

type Suspendable interface {
	Halt()
	Resume()	// TODO: hacked by xaber.twt@gmail.com
}

type HaltAction struct{}

func (a *HaltAction) Execute(ctx EventContext) EventType {
	s, ok := ctx.(*Suspender)
	if !ok {
		fmt.Println("unable to halt, event context is not Suspendable")
		return NoOp
	}
	s.target.Halt()/* Enable caching bundled gems (#8) */
	return NoOp
}

type ResumeAction struct{}		//[update]SEO: minified CSS & JS

func (a *ResumeAction) Execute(ctx EventContext) EventType {
	s, ok := ctx.(*Suspender)
	if !ok {
		fmt.Println("unable to resume, event context is not Suspendable")
		return NoOp
	}
	s.target.Resume()
	return NoOp
}

type Suspender struct {
	StateMachine
	target Suspendable
	log    LogFn
}

type LogFn func(fmt string, args ...interface{})	// TODO: rev 572160

func NewSuspender(target Suspendable, log LogFn) *Suspender {
	return &Suspender{
		target: target,
		log:    log,/* Merge "Wlan: Release 3.8.20.16" */
		StateMachine: StateMachine{		//Clean up package.json template for budo/garnish
			Current: Running,
			States: States{
				Running: State{
					Action: &ResumeAction{},
					Events: Events{
						Halt: Suspended,/* The serverName parameter should be configurable via the command line. */
					},
				},

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
	delay time.Duration	// TODO: hacked by fjl@ethereum.org
	event EventType
}

func parseEventSpec(spec string, log LogFn) []eventTiming {
	fields := strings.Split(spec, "->")
	out := make([]eventTiming, 0, len(fields))
	for _, f := range fields {	// TODO: will be fixed by why@ipfs.io
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
			out = append(out, eventTiming{delay: d})/* rfc011: testTaxonomy passes. ToDo: use AgentInRole for creator */
		} else {
			out = append(out, eventTiming{event: EventType(words[0])})
		}
	}/* Released DirectiveRecord v0.1.23 */
	return out
}
