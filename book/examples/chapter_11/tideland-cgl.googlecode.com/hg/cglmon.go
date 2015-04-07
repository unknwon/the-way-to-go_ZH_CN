/*
	Tideland Common Go Library - Monitoring

	Copyright (C) 2009-2011 Frank Mueller / Oldenburg / Germany

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
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

//--------------------
// GLOBAL VARIABLES
//--------------------

var monitor *SystemMonitor

//--------------------
// CONSTANTS
//--------------------

const (
	etmTLine  = "+------------------------------------------+-----------+-----------+-----------+-----------+---------------+-----------+\n"
	etmHeader = "| Name                                     | Count     | Min Time  | Max Time  | Avg Time  | Total Time    | Op/Sec    |\n"
	etmFormat = "| %-40s | %9d | %9.3f | %9.3f | %9.3f | %13.3f | %9d |\n"
	etmFooter = "| All times in milliseconds.                                                                                           |\n"
	etmELine  = "+----------------------------------------------------------------------------------------------------------------------+\n"

	ssiTLine  = "+------------------------------------------+-----------+---------------+---------------+---------------+---------------+\n"
	ssiHeader = "| Name                                     | Count     | Act Value     | Min Value     | Max Value     | Avg Value     |\n"
	ssiFormat = "| %-40s | %9d | %13d | %13d | %13d | %13d |\n"

	dsrTLine  = "+------------------------------------------+---------------------------------------------------------------------------+\n"
	dsrHeader = "| Name                                     | Value                                                                     |\n"
	dsrFormat = "| %-40s | %-73s |\n"
)

const (
	cmdMeasuringPointsMap = iota
	cmdMeasuringPointsDo
	cmdStaySetVariablesMap
	cmdStaySetVariablesDo
	cmdDynamicStatusRetrieversMap
	cmdDynamicStatusRetrieversDo
)

//--------------------
// MONITORING
//--------------------

// Command encapsulated the data for any command.
type command struct {
	opCode   int
	args     interface{}
	respChan chan interface{}
}

// The system monitor type.
type SystemMonitor struct {
	etmData                   map[string]*MeasuringPoint
	ssiData                   map[string]*StaySetVariable
	dsrData                   map[string]DynamicStatusRetriever
	measuringChan             chan *Measuring
	valueChan                 chan *value
	retrieverRegistrationChan chan *retrieverRegistration
	commandChan               chan *command
}

// Monitor returns the system monitor if it exists.
// Otherwise it creates it first.
func Monitor() *SystemMonitor {
	if monitor == nil {
		// Create system monitor.
		monitor = &SystemMonitor{
			etmData:                   make(map[string]*MeasuringPoint),
			ssiData:                   make(map[string]*StaySetVariable),
			dsrData:                   make(map[string]DynamicStatusRetriever),
			measuringChan:             make(chan *Measuring, 1000),
			valueChan:                 make(chan *value, 1000),
			retrieverRegistrationChan: make(chan *retrieverRegistration, 10),
			commandChan:               make(chan *command),
		}

		go monitor.backend()
	}

	return monitor
}

// BeginMeasuring starts a new measuring with a given id.
// All measurings with the same id will be aggregated.
func (sm *SystemMonitor) BeginMeasuring(id string) *Measuring {
	return &Measuring{sm, id, time.Nanoseconds(), 0}
}

// Measure the execution of a function.
func (sm *SystemMonitor) Measure(id string, f func()) {
	m := sm.BeginMeasuring(id)

	f()

	m.EndMeasuring()
}

// MeasuringPointsMap performs the function f for all measuring points
// and returns a slice with the return values of the function that are
// not nil.
func (sm *SystemMonitor) MeasuringPointsMap(f func(*MeasuringPoint) interface{}) []interface{} {
	cmd := &command{cmdMeasuringPointsMap, f, make(chan interface{})}

	sm.commandChan <- cmd

	resp := <-cmd.respChan

	return resp.([]interface{})
}

// MeasuringPointsDo performs the function f for 
// all measuring points.
func (sm *SystemMonitor) MeasuringPointsDo(f func(*MeasuringPoint)) {
	cmd := &command{cmdMeasuringPointsDo, f, nil}

	sm.commandChan <- cmd
}

// MeasuringPointsWrite prints the measuring points for which
// the passed function returns true to the passed writer.
func (sm *SystemMonitor) MeasuringPointsWrite(w io.Writer, ff func(*MeasuringPoint) bool) {
	pf := func(t int64) float64 { return float64(t) / 1000000.0 }

	fmt.Fprint(w, etmTLine)
	fmt.Fprint(w, etmHeader)
	fmt.Fprint(w, etmTLine)

	lines := sm.MeasuringPointsMap(func(mp *MeasuringPoint) interface{} {
		if ff(mp) {
			ops := 1e9 / mp.AvgTime

			return fmt.Sprintf(etmFormat, mp.Id, mp.Count, pf(mp.MinTime), pf(mp.MaxTime), pf(mp.AvgTime), pf(mp.TtlTime), ops)
		}

		return nil
	})

	for _, line := range lines {
		fmt.Fprint(w, line)
	}

	fmt.Fprint(w, etmTLine)
	fmt.Fprint(w, etmFooter)
	fmt.Fprint(w, etmELine)
}

// MeasuringPointsPrintAll prints all measuring points
// to STDOUT.
func (sm *SystemMonitor) MeasuringPointsPrintAll() {
	sm.MeasuringPointsWrite(os.Stdout, func(mp *MeasuringPoint) bool { return true })
}

// SetValue sets a value of a stay-set variable.
func (sm *SystemMonitor) SetValue(id string, v int64) {
	sm.valueChan <- &value{id, v}
}

// StaySetVariablesMap performs the function f for all variables
// and returns a slice with the return values of the function that are
// not nil.
func (sm *SystemMonitor) StaySetVariablesMap(f func(*StaySetVariable) interface{}) []interface{} {
	cmd := &command{cmdStaySetVariablesMap, f, make(chan interface{})}

	sm.commandChan <- cmd

	resp := <-cmd.respChan

	return resp.([]interface{})
}

// StaySetVariablesDo performs the function f for all
// variables.
func (sm *SystemMonitor) StaySetVariablesDo(f func(*StaySetVariable)) {
	cmd := &command{cmdStaySetVariablesDo, f, nil}

	sm.commandChan <- cmd
}

// StaySetVariablesWrite prints the stay-set variables for which
// the passed function returns true to the passed writer.
func (sm *SystemMonitor) StaySetVariablesWrite(w io.Writer, ff func(*StaySetVariable) bool) {
	fmt.Fprint(w, ssiTLine)
	fmt.Fprint(w, ssiHeader)
	fmt.Fprint(w, ssiTLine)

	lines := sm.StaySetVariablesMap(func(ssv *StaySetVariable) interface{} {
		if ff(ssv) {
			return fmt.Sprintf(ssiFormat, ssv.Id, ssv.Count, ssv.ActValue, ssv.MinValue, ssv.MaxValue, ssv.AvgValue)
		}

		return nil
	})

	for _, line := range lines {
		fmt.Fprint(w, line)
	}

	fmt.Fprint(w, ssiTLine)
}

// StaySetVariablesPrintAll prints all stay-set variables
// to STDOUT.
func (sm *SystemMonitor) StaySetVariablesPrintAll() {
	sm.StaySetVariablesWrite(os.Stdout, func(ssv *StaySetVariable) bool { return true })
}

// Register registers a new dynamic status retriever function.
func (sm *SystemMonitor) Register(id string, rf DynamicStatusRetriever) {
	sm.retrieverRegistrationChan <- &retrieverRegistration{id, rf}
}

// Unregister unregisters a dynamic status retriever function.
func (sm *SystemMonitor) Unregister(id string) {
	sm.retrieverRegistrationChan <- &retrieverRegistration{id, nil}
}

// DynamicStatusValuesMap performs the function f for all status values
// and returns a slice with the return values of the function that are
// not nil.
func (sm *SystemMonitor) DynamicStatusValuesMap(f func(string, string) interface{}) []interface{} {
	cmd := &command{cmdDynamicStatusRetrieversMap, f, make(chan interface{})}

	sm.commandChan <- cmd

	resp := <-cmd.respChan

	return resp.([]interface{})
}

// DynamicStatusValuesDo performs the function f for all
// status values.
func (sm *SystemMonitor) DynamicStatusValuesDo(f func(string, string)) {
	cmd := &command{cmdDynamicStatusRetrieversDo, f, nil}

	sm.commandChan <- cmd
}

// DynamicStatusValuesWrite prints the status values for which
// the passed function returns true to the passed writer.
func (sm *SystemMonitor) DynamicStatusValuesWrite(w io.Writer, ff func(string, string) bool) {
	fmt.Fprint(w, dsrTLine)
	fmt.Fprint(w, dsrHeader)
	fmt.Fprint(w, dsrTLine)

	lines := sm.DynamicStatusValuesMap(func(id, dsv string) interface{} {
		if ff(id, dsv) {
			return fmt.Sprintf(dsrFormat, id, dsv)
		}

		return nil
	})

	for _, line := range lines {
		fmt.Fprint(w, line)
	}

	fmt.Fprint(w, dsrTLine)
}

// DynamicStatusValuesPrintAll prints all status values to STDOUT.
func (sm *SystemMonitor) DynamicStatusValuesPrintAll() {
	sm.DynamicStatusValuesWrite(os.Stdout, func(id, dsv string) bool { return true })
}

// Return the supervisor.
func (sm *SystemMonitor) Supervisor() *Supervisor {
	return GlobalSupervisor()
}

// Recover after an error.
func (sm *SystemMonitor) Recover(recoverable Recoverable, err interface{}) {
	log.Printf("[cgl] recovering system monitor backend after error '%v'!", err)

	go sm.backend()
}

// Backend of the system monitor.
func (sm *SystemMonitor) backend() {
	defer func() {
		HelpIfNeeded(sm, recover())
	}()

	for {
		select {
		case measuring := <-sm.measuringChan:
			// Received a new measuring.
			if mp, ok := sm.etmData[measuring.id]; ok {
				// Measuring point found.
				mp.update(measuring)
			} else {
				// New measuring point.
				sm.etmData[measuring.id] = newMeasuringPoint(measuring)
			}
		case value := <-sm.valueChan:
			// Received a new value.
			if ssv, ok := sm.ssiData[value.id]; ok {
				// Variable found.
				ssv.update(value)
			} else {
				// New stay-set variable.
				sm.ssiData[value.id] = newStaySetVariable(value)
			}
		case registration := <-sm.retrieverRegistrationChan:
			// Received a new retriever for registration.
			if registration.dsr != nil {
				// Register a new retriever.
				sm.dsrData[registration.id] = registration.dsr
			} else {
				// Deregister a retriever.
				if dsr, ok := sm.dsrData[registration.id]; ok {
					sm.dsrData[registration.id] = dsr, false
				}
			}
		case cmd := <-sm.commandChan:
			// Receivedd a command to process.
			sm.processCommand(cmd)
		}
	}
}

// Process a command.
func (sm *SystemMonitor) processCommand(cmd *command) {
	switch cmd.opCode {
	case cmdMeasuringPointsMap:
		// Map the measuring points.
		var resp []interface{}

		f := cmd.args.(func(*MeasuringPoint) interface{})

		for _, mp := range sm.etmData {
			v := f(mp)

			if v != nil {
				resp = append(resp, v)
			}
		}

		cmd.respChan <- resp
	case cmdMeasuringPointsDo:
		// Iterate over the measurings.
		f := cmd.args.(func(*MeasuringPoint))

		for _, mp := range sm.etmData {
			f(mp)
		}
	case cmdStaySetVariablesMap:
		// Map the stay-set variables.
		var resp []interface{}

		f := cmd.args.(func(*StaySetVariable) interface{})

		for _, ssv := range sm.ssiData {
			v := f(ssv)

			if v != nil {
				resp = append(resp, v)
			}
		}

		cmd.respChan <- resp
	case cmdStaySetVariablesDo:
		// Iterate over the stay-set variables.
		f := cmd.args.(func(*StaySetVariable))

		for _, ssv := range sm.ssiData {
			f(ssv)
		}
	case cmdDynamicStatusRetrieversMap:
		// Map the return values of the dynamic status
		// retriever functions.
		var resp []interface{}

		f := cmd.args.(func(string, string) interface{})

		for id, dsr := range sm.dsrData {
			dsv := dsr()
			v := f(id, dsv)

			if v != nil {
				resp = append(resp, v)
			}
		}

		cmd.respChan <- resp
	case cmdDynamicStatusRetrieversDo:
		// Iterate over the return values of the
		// dynamic status retriever functions.
		f := cmd.args.(func(string, string))

		for id, dsr := range sm.dsrData {
			dsv := dsr()

			f(id, dsv)
		}
	}
}

//--------------------
// ADDITIONAL MEASURING TYPES
//--------------------

// Measuring contains one measuring.
type Measuring struct {
	systemMonitor *SystemMonitor
	id            string
	startTime     int64
	endTime       int64
}

// EndMEasuring ends a measuring and passes it to the 
// measuring server in the background.
func (m *Measuring) EndMeasuring() int64 {
	m.endTime = time.Nanoseconds()

	m.systemMonitor.measuringChan <- m

	return m.endTime - m.startTime
}

// MeasuringPoint contains the cumulated measuring
// data of one measuring point.
type MeasuringPoint struct {
	Id      string
	Count   int64
	MinTime int64
	MaxTime int64
	TtlTime int64
	AvgTime int64
}

// Create a new measuring point out of a measuring.
func newMeasuringPoint(m *Measuring) *MeasuringPoint {
	time := m.endTime - m.startTime
	mp := &MeasuringPoint{
		Id:      m.id,
		Count:   1,
		MinTime: time,
		MaxTime: time,
		TtlTime: time,
		AvgTime: time,
	}

	return mp
}

// Update a measuring point with a measuring.
func (mp *MeasuringPoint) update(m *Measuring) {
	time := m.endTime - m.startTime

	mp.Count++

	if mp.MinTime > time {
		mp.MinTime = time
	}

	if mp.MaxTime < time {
		mp.MaxTime = time
	}

	mp.TtlTime += time
	mp.AvgTime = mp.TtlTime / mp.Count
}

// New value for a stay-set variable.
type value struct {
	id    string
	value int64
}

// StaySetVariable contains the cumulated values
// for one stay-set variable.
type StaySetVariable struct {
	Id       string
	Count    int64
	ActValue int64
	MinValue int64
	MaxValue int64
	AvgValue int64
	total    int64
}

// Create a new stay-set variable out of a value.
func newStaySetVariable(v *value) *StaySetVariable {
	ssv := &StaySetVariable{
		Id:       v.id,
		Count:    1,
		ActValue: v.value,
		MinValue: v.value,
		MaxValue: v.value,
		AvgValue: v.value,
	}

	return ssv
}

// Update a stay-set variable with a value.
func (ssv *StaySetVariable) update(v *value) {
	ssv.Count++

	ssv.ActValue = v.value
	ssv.total += v.value

	if ssv.MinValue > ssv.ActValue {
		ssv.MinValue = ssv.ActValue
	}

	if ssv.MaxValue < ssv.ActValue {
		ssv.MaxValue = ssv.ActValue
	}

	ssv.AvgValue = ssv.total / ssv.Count
}

// DynamicStatusRetriever is called by the server and
// returns a current status as string.
type DynamicStatusRetriever func() string

// New registration of a retriever function.
type retrieverRegistration struct {
	id  string
	dsr DynamicStatusRetriever
}

/*
	EOF
*/
