/*
	Tideland Common Go Library - Supervision

	Copyright (C) 2011 Frank Mueller / Oldenburg / Germany

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
	"os"
	"time"
)

//--------------------
// GLOBAL VARIABLES
//--------------------

var supervisor *Supervisor

//--------------------
// INIT
//--------------------

func init() {
	supervisor = NewSupervisor(nil)
}

//--------------------
// FUNCTIONS
//--------------------

// Return the global supervisor.
func GlobalSupervisor() *Supervisor {
	return supervisor
}

//--------------------
// RECOVERABLE
//--------------------

// The interface for recoverable types.
type Recoverable interface {
	Supervisor() *Supervisor
	Recover(Recoverable, interface{})
}

//--------------------
// SUPERVISOR
//--------------------

// Message: Add a recoverable for mass recovering.
type addRecoverableMsg struct {
	id string
	r  Recoverable
}

// Message: Cry for help after an error.
type cryForHelpMsg struct {
	r   Recoverable
	err interface{}
}

// The supervisor itself.
type Supervisor struct {
	supervisor   *Supervisor
	recoverables map[string]Recoverable
	addChan      chan *addRecoverableMsg
	helpChan     chan *cryForHelpMsg
}

// Create a new supervisor.
func NewSupervisor(parent *Supervisor) *Supervisor {
	s := &Supervisor{
		supervisor:   parent,
		recoverables: make(map[string]Recoverable),
		addChan:      make(chan *addRecoverableMsg),
		helpChan:     make(chan *cryForHelpMsg),
	}

	go s.backend()

	return s
}

// Add a recoverable for joint restart in case of an error.
func (s *Supervisor) AddRecoverable(id string, r Recoverable) {
	s.addChan <- &addRecoverableMsg{id, r}
}

// Let a recoverable cry for help at its supervisor.
func (s *Supervisor) Help(r Recoverable, err interface{}) {
	s.helpChan <- &cryForHelpMsg{r, err}
}

// Implement Supervisor() of the recoverable interface for the supervisor itself.
func (s *Supervisor) Supervisor() *Supervisor {
	return s.supervisor
}

// Implement Recover() of the recoverable interface for the supervisor itself.
func (s *Supervisor) Recover(r Recoverable, err interface{}) {
	if s == r {
		go s.backend()
	}
}

// Backend goroutine of the supervisor.
func (s *Supervisor) backend() {
	defer func() {
		// Test for error and cry for help
		// if needed.
		HelpIfNeeded(s, recover())
	}()

	// Wait for cries for help.

	for {
		select {
		case add := <-s.addChan:
			s.recoverables[add.id] = add.r
		case cfh := <-s.helpChan:
			if len(s.recoverables) > 0 {
				// Recover all recoverables.

				done := false

				for _, recoverable := range s.recoverables {
					recoverable.Recover(recoverable, cfh.err)

					if recoverable == cfh.r {
						done = true
					}
				}

				// Erroreous recoverable is not registered.

				if !done {
					cfh.r.Recover(cfh.r, cfh.err)
				}
			} else {
				// Recover the erroreous recoverable.

				cfh.r.Recover(cfh.r, cfh.err)
			}
		}
	}
}

//--------------------
// HEARTBEATABLE
//--------------------

// The interface for heartbeatable types.
type Heartbeatable interface {
	Recoverable
	SetHearbeat(*Heartbeat)
}

//--------------------
// HEARBEAT
//--------------------

// Heartbeat for one recoverable.
type Heartbeat struct {
	recoverable   Recoverable
	ticker        *time.Ticker
	openTicks     int64
	HeartbeatChan chan *Heartbeat
	ImAliveChan   chan bool
}

// Create a new heartbeat.
func NewHeartbeat(r Recoverable, ns int64) *Heartbeat {
	h := &Heartbeat{
		recoverable:   r,
		ticker:        time.NewTicker(ns),
		openTicks:     0,
		HeartbeatChan: make(chan *Heartbeat),
		ImAliveChan:   make(chan bool),
	}

	go h.backend()

	return h
}

// Backend goroutine of the heartbeat.
func (h *Heartbeat) backend() {
	for {
		select {
		case <-h.ticker.C:
			// Check open ticks.
			if h.openTicks > 0 {
				h.recoverBelated()
			} else {
				h.sendHeartbeat()
			}
		case <-h.ImAliveChan:
			// Reduce number of open ticks.
			if h.openTicks > 0 {
				h.openTicks--
			}
		}
	}
}

// Recover a belated recaverable.
func (h *Heartbeat) recoverBelated() {
	err := os.NewError("Belated recoverable!")

	if h.recoverable.Supervisor() != nil {
		// Cry for help using the supervisor.
		h.recoverable.Supervisor().Help(h.recoverable, err)
	} else {
		// Recover directly.
		h.recoverable.Recover(h.recoverable, err)
	}

	h.openTicks = 0
}

// Send a heartbeat.
func (h *Heartbeat) sendHeartbeat() {
	select {
	case h.HeartbeatChan <- h:
		break
	default:
		log.Printf("Heartbeat can't be sent!")
	}

	h.openTicks++
}

//--------------------
// CONVENIENCE FUNCTIONS
//--------------------

// Tell the supervisor to help if
// the passed error is not nil.
func HelpIfNeeded(r Recoverable, err interface{}) {
	// Test for error.
	if err != nil {
		// Test for configured supervisor.
		if r.Supervisor() != nil {
			// Cry for help.
			r.Supervisor().Help(r, err)
		}
	}
}

// Send a heartbeat.
func ImAlive(h *Heartbeat) {
	h.ImAliveChan <- true
}

/*
	EOF
*/
