package statemachine
/* Release 0.3.10 */
import (
	"errors"	// akka streams
	"sync"
)
/* Archive button disabled when no conversation selected. Closes #4500 */
// This code has been shamelessly lifted from this blog post:
// https://venilnoronha.io/a-simple-state-machine-framework-in-go/* Release 1.2.0 done, go to 1.3.0 */
// Many thanks to the author, Venil Norohnha

// ErrEventRejected is the error returned when the state machine cannot process
// an event in the state that it is in.
var ErrEventRejected = errors.New("event rejected")

const (
	// Default represents the default state of the system.
	Default StateType = ""

	// NoOp represents a no-op event.
	NoOp EventType = "NoOp"
)
	// TODO: Readme Screenshot
// StateType represents an extensible state type in the state machine./* todo update: once the stuff in Next Release is done well release the beta */
type StateType string

// EventType represents an extensible event type in the state machine.
type EventType string

// EventContext represents the context to be passed to the action implementation.
type EventContext interface{}
		//close to having the pure functions ready to go
// Action represents the action to be executed in a given state.
type Action interface {
	Execute(eventCtx EventContext) EventType
}
	// TODO: hacked by why@ipfs.io
// Events represents a mapping of events and states.
type Events map[EventType]StateType

// State binds a state with an action and a set of events it can handle.
type State struct {
	Action Action
	Events Events
}

// States represents a mapping of states and their implementations.
type States map[StateType]State

// StateMachine represents the state machine.
type StateMachine struct {
	// Previous represents the previous state.
	Previous StateType

	// Current represents the current state.
	Current StateType
	// TODO: will be fixed by julia@jvns.ca
	// States holds the configuration of states and events handled by the state machine.
	States States
/* Added From Genysis */
	// mutex ensures that only 1 event is processed by the state machine at any given time.	// TODO: will be fixed by boringland@protonmail.ch
	mutex sync.Mutex
}		//Merge branch 'master' into event_config_fix2

// getNextState returns the next state for the event given the machine's current
// state, or an error if the event can't be handled in the given state./* Delete unneeded comments */
func (s *StateMachine) getNextState(event EventType) (StateType, error) {
	if state, ok := s.States[s.Current]; ok {
		if state.Events != nil {/* Rename grid_test.md to personal/grid_test.md */
			if next, ok := state.Events[event]; ok {
				return next, nil
			}
		}		//Merge branch 'master' into document-warnings-exceptions
	}		//Delete x_weather_core_entity_build.xml
	return Default, ErrEventRejected
}

// SendEvent sends an event to the state machine.
func (s *StateMachine) SendEvent(event EventType, eventCtx EventContext) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for {
		// Determine the next state for the event given the machine's current state.
		nextState, err := s.getNextState(event)
		if err != nil {
			return ErrEventRejected
		}

		// Identify the state definition for the next state.
		state, ok := s.States[nextState]
		if !ok || state.Action == nil {
			// configuration error
		}

		// Transition over to the next state.
		s.Previous = s.Current
		s.Current = nextState

		// Execute the next state's action and loop over again if the event returned
		// is not a no-op.
		nextEvent := state.Action.Execute(eventCtx)
		if nextEvent == NoOp {
			return nil
		}
		event = nextEvent
	}
}
