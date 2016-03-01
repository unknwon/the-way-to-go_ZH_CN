/*
	Tideland Common Go Library - Time and Crontab

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
	"log"
	"time"
)

//--------------------
// DATE AND TIME
//--------------------

// Calc nanoseconds from microseconds.
func NsMicroseconds(count int64) int64 { return count * 1e3 }

// Calc nanoseconds from milliseconds.
func NsMilliseconds(count int64) int64 { return NsMicroseconds(count * 1e3) }

// Calc nanoseconds from seconds.
func NsSeconds(count int64) int64 { return NsMilliseconds(count * 1e3) }

// Calc nanoseconds from minutes.
func NsMinutes(count int64) int64 { return NsSeconds(count * 60) }

// Calc nanoseconds from hours.
func NsHours(count int64) int64 { return NsMinutes(count * 60) }

// Calc nanoseconds from days.
func NsDays(count int64) int64 { return NsHours(count * 24) }

// Calc nanoseconds from weeks.
func NsWeeks(count int64) int64 { return NsDays(count * 7) }

// Test if the year of a time is in a given list.
func YearInList(time *time.Time, years []int64) bool {
	for _, year := range years {
		if time.Year == year {
			return true
		}
	}

	return false
}

// Test if a year of a time is in a given range.
func YearInRange(time *time.Time, minYear, maxYear int64) bool {
	return (minYear <= time.Year) && (time.Year <= maxYear)
}

// Test if the month of a time is in a given list.
func MonthInList(time *time.Time, months []int) bool {
	return fieldInList(time.Month, months)
}

// Test if a month of a time is in a given range.
func MonthInRange(time *time.Time, minMonth, maxMonth int) bool {
	return fieldInRange(time.Month, minMonth, maxMonth)
}

// Test if the day of a time is in a given list.
func DayInList(time *time.Time, days []int) bool {
	return fieldInList(time.Day, days)
}

// Test if a day of a time is in a given range.
func DayInRange(time *time.Time, minDay, maxDay int) bool {
	return fieldInRange(time.Day, minDay, maxDay)
}

// Test if the hour of a time is in a given list.
func HourInList(time *time.Time, hours []int) bool {
	return fieldInList(time.Hour, hours)
}

// Test if a hour of a time is in a given range.
func HourInRange(time *time.Time, minHour, maxHour int) bool {
	return fieldInRange(time.Hour, minHour, maxHour)
}

// Test if the minute of a time is in a given list.
func MinuteInList(time *time.Time, minutes []int) bool {
	return fieldInList(time.Minute, minutes)
}

// Test if a minute of a time is in a given range.
func MinuteInRange(time *time.Time, minMinute, maxMinute int) bool {
	return fieldInRange(time.Minute, minMinute, maxMinute)
}

// Test if the second of a time is in a given list.
func SecondInList(time *time.Time, seconds []int) bool {
	return fieldInList(time.Second, seconds)
}

// Test if a second of a time is in a given range.
func SecondInRange(time *time.Time, minSecond, maxSecond int) bool {
	return fieldInRange(time.Second, minSecond, maxSecond)
}

// Test if the weekday of a time is in a given list.
func WeekdayInList(time *time.Time, weekdays []int) bool {
	return fieldInList(time.Weekday, weekdays)
}

// Test if a weekday of a time is in a given range.
func WeekdayInRange(time *time.Time, minWeekday, maxWeekday int) bool {
	return fieldInRange(time.Weekday, minWeekday, maxWeekday)
}

//--------------------
// JOB
//--------------------

// Check function type.
type CheckFunc func(*time.Time) (bool, bool)

// Tast function type.
type TaskFunc func(string)

// Job type.
type Job struct {
	id    string
	check CheckFunc
	task  TaskFunc
}

// Create a new job.
func NewJob(id string, check CheckFunc, task TaskFunc) *Job {
	return &Job{id, check, task}
}

// Check if the job has to be performed at a given time.
func (job *Job) checkAndPerform(time *time.Time) bool {
	perform, delete := job.check(time)

	if perform {
		go job.task(job.id)
	}

	return perform && delete
}

//--------------------
// CRONTAB
//--------------------

const (
	opJobAdd = iota
	opJobDel
	opCrontabStop
)

// Crontab control type.
type crontabControl struct {
	opCode int
	args   interface{}
}

// Crontab.
type Crontab struct {
	jobs    map[string]*Job
	control chan *crontabControl
	ticker  *time.Ticker
}

// Start a crontab server.
func NewCrontab() *Crontab {
	ctb := &Crontab{
		jobs:    make(map[string]*Job),
		control: make(chan *crontabControl),
		ticker:  time.NewTicker(1e9),
	}

	go ctb.backend()

	return ctb
}

// Stop the server.
func (ctb *Crontab) Stop() {
	ctb.control <- &crontabControl{opCrontabStop, nil}
}

// Add a job to the server.
func (ctb *Crontab) AddJob(job *Job) {
	ctb.control <- &crontabControl{opJobAdd, job}
}

// Delete a job from the server.
func (ctb *Crontab) DeleteJob(id string) {
	ctb.control <- &crontabControl{opJobDel, id}
}

// Return the supervisor.
func (src *Crontab) Supervisor() *Supervisor {
	return GlobalSupervisor()
}

// Recover after an error.
func (ctb *Crontab) Recover(recoverable Recoverable, err interface{}) {
	log.Printf("Recovering crontab backend after error '%v'!", err)

	go ctb.backend()
}

// Crontab backend.
func (ctb *Crontab) backend() {
	defer func() {
		HelpIfNeeded(ctb, recover())
	}()

	for {
		select {
		case sc := <-ctb.control:
			// Control the server.

			switch sc.opCode {
			case opJobAdd:
				job := sc.args.(*Job)
				ctb.jobs[job.id] = job
			case opJobDel:
				id := sc.args.(string)
				job, _ := ctb.jobs[id]
				ctb.jobs[id] = job, false
			case opCrontabStop:
				ctb.ticker.Stop()
			}
		case <-ctb.ticker.C:
			// One tick every second.

			ctb.tick()
		}
	}
}

// Handle one server tick.
func (ctb *Crontab) tick() {
	now := time.UTC()
	deletes := make(map[string]*Job)

	// Check and perform jobs.

	for id, job := range ctb.jobs {
		delete := job.checkAndPerform(now)

		if delete {
			deletes[id] = job
		}
	}

	// Delete those marked for deletion.

	for id, job := range deletes {
		ctb.jobs[id] = job, false
	}
}

//--------------------
// HELPERS
//--------------------

// Test if an int is in a list of ints.
func fieldInList(field int, list []int) bool {
	for _, item := range list {
		if field == item {
			return true
		}
	}

	return false
}

// Test if an int is in a given range.
func fieldInRange(field int, min, max int) bool {
	return (min <= field) && (field <= max)
}

/*
	EOF
*/
