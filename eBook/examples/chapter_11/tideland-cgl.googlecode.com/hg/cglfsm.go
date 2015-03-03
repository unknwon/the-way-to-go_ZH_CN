/*
	Tideland Common Go Library - Finite State Machine

	Copyright (C) 2010-2011 Frank Mueller / Oldenburg / Germany

	Redistribution and use in source and binary forms, with or
	modification, are permitted provided that the following conditions are
	met:

	Redistributions of source code must retain the above copyright notice, this
	list of conditions and the following disclaimer.

	Redistributions in binary form must reproduce the above copyright notice,
	this list of conditions and the following disclaimer in the documentation
	and/or other materials provided with the distribution.

	Neither the name of Tideland nor the names of its contributors may be
	used to endorse or promote products derived from this software without
	specific prior written permission.

	THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
	AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
	IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
	ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE
	LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
	CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
	SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
	INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
	CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
	ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF
	THE POSSIBILITY OF SUCH DAMAGE.
*/

package cgl

//--------------------
// IMPORTS
//--------------------

import (
	"log"
	"reflect"
	"strings"
	"time"
)

//--------------------
// HELPER TYPES
//--------------------

// Condition type.
type Condition struct {
	Now     int64
	Payload interface{}
}

// Transition type.
type transition struct {
	payload    interface{}
	resultChan chan interface{}
}

// Timeout type.
type Timeout int64

//--------------------
// FINITE STATE MACHINE
//--------------------

// Handler interface.
type Handler interface {
	Init() string
	Terminate(string, interface{}) string
}

// State machine type.
type FSM struct {
	Handler        Handler
	handlerValue   reflect.Value
	handlerFuncs   map[string]reflect.Value
	state          string
	transitionChan chan *transition
	timeoutChan    <-chan int64
}

// Create a new finite state machine.
func NewFSM(h Handler, timeout int64) *FSM {
	var bufferSize int

	if timeout > 0 {
		bufferSize = int(timeout / 1e3)
	} else {
		bufferSize = 10
	}

	fsm := &FSM{
		Handler:        h,
		handlerFuncs:   make(map[string]reflect.Value),
		state:          h.Init(),
		transitionChan: make(chan *transition, bufferSize),
	}

	if timeout > 0 {
		fsm.timeoutChan = time.After(timeout)
	}

	fsm.analyze()

	go fsm.backend()

	return fsm
}

// Send a payload to handle and return the result.
func (fsm *FSM) SendWithResult(payload interface{}) interface{} {
	t := &transition{payload, make(chan interface{})}

	fsm.transitionChan <- t

	return <-t.resultChan
}

// Send a payload with no result.
func (fsm *FSM) Send(payload interface{}) {
	t := &transition{payload, nil}

	fsm.transitionChan <- t
}

// Send a payload with no result after a given time.
func (fsm *FSM) SendAfter(payload interface{}, ns int64) {
	saf := func() {
		time.Sleep(ns)

		fsm.Send(payload)
	}

	go saf()
}

// Return the current state.
func (fsm *FSM) State() string {
	return fsm.state
}

// Return the supervisor.
func (fsm *FSM) Supervisor() *Supervisor {
	return GlobalSupervisor()
}

// Recover after an error.
func (fsm *FSM) Recover(recoverable Recoverable, err interface{}) {
	log.Printf("[cgl] recovering finite state machine server backend after error '%v'!", err)

	go fsm.backend()
}

// Analyze the event handler and prepare the state table.
func (fsm *FSM) analyze() {
	prefix := "HandleState"

	fsm.handlerValue = reflect.ValueOf(fsm.Handler)

	num := fsm.handlerValue.Type().NumMethod()

	for i := 0; i < num; i++ {
		meth := fsm.handlerValue.Type().Method(i)

		if (meth.PkgPath == "") && (strings.HasPrefix(meth.Name, prefix)) {
			if (meth.Type.NumIn() == 2) && (meth.Type.NumOut() == 2) {
				state := meth.Name[len(prefix):len(meth.Name)]

				fsm.handlerFuncs[state] = meth.Func
			}
		}
	}
}

// State machine backend.
func (fsm *FSM) backend() {
	defer func() {
		HelpIfNeeded(fsm, recover())
	}()

	// Message loop.

	for {
		select {
		case t := <-fsm.transitionChan:
			// Regular transition.

			if nextState, ok := fsm.handle(t); ok {
				// Continue.

				fsm.state = nextState
			} else {
				// Stop processing.

				fsm.state = fsm.Handler.Terminate(fsm.state, nextState)

				return
			}
		case to := <-fsm.timeoutChan:
			// Timeout signal resent to let it be handled.

			t := &transition{Timeout(to), nil}

			fsm.transitionChan <- t
		}
	}
}

// Handle a transition.
func (fsm *FSM) handle(t *transition) (string, bool) {
	condition := &Condition{time.Nanoseconds(), t.payload}
	handlerFunc := fsm.handlerFuncs[fsm.state]
	handlerArgs := make([]reflect.Value, 2)

	handlerArgs[0] = fsm.handlerValue
	handlerArgs[1] = reflect.ValueOf(condition)

	results := handlerFunc.Call(handlerArgs)

	nextState := results[0].Interface().(string)
	result := results[1].Interface()

	// Return a result if wanted.

	if t.resultChan != nil {
		t.resultChan <- result
	}

	// Check for termination.

	if nextState == "Terminate" {
		return nextState, false
	}

	return nextState, true
}

/*
	EOF
*/
