/*
	Tideland Common Go Library - Unit Tests

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
	"bytes"
	"fmt"
	"log"
	"rand"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"
)

//--------------------
// TESTS
//--------------------

// Test single recovering.
func TestSingleRecovering(t *testing.T) {
	s := NewSupervisor(nil)
	ra := NewRecoverableAction(s)

	pos := ra.Action(PositiveAction)

	if pos == "OK" {
		t.Logf("Single recovering (a) OK.")
	} else {
		t.Errorf("Single recovering (a) failed! Reply is '%v'.", pos)
	}

	neg := ra.Action(FailAction)

	if neg == "Recovered" {
		t.Logf("Single recovering (b) OK.")
	} else {
		t.Errorf("Single recovering (b) failed! Reply is '%v.", neg)
	}
}

// Test multiple recovering.
func TestMultipleRecovering(t *testing.T) {
	s := NewSupervisor(nil)
	raa := NewRecoverableAction(s)
	rab := NewRecoverableAction(s)
	rac := NewRecoverableAction(s)

	s.AddRecoverable("A", raa)
	s.AddRecoverable("B", rab)
	s.AddRecoverable("C", rac)

	t.Logf("(A) is '%v'.", raa.Action(FailAction))
	t.Logf("(B) is '%v'.", rab.Action(PositiveAction))
	t.Logf("(C) is '%v'.", rac.Action(PositiveAction))
}

// Test single heartbeat timeout.
func TestSingleHeartbeatTimeout(t *testing.T) {
	ra := NewRecoverableAction(nil)
	reply := ra.Action(TimeConsumingAction)

	if reply == "Recovered" {
		t.Logf("Heartbeat timeout recovering OK.")
	} else {
		t.Errorf("Heartbeat timeout recovering failed! Reply is '%v.", reply)
	}
}

// Test multiple heartbeat timeout.
func TestMultipleHeartbeatTimeout(t *testing.T) {
	s := NewSupervisor(nil)
	raa := NewRecoverableAction(s)
	rab := NewRecoverableAction(s)
	rac := NewRecoverableAction(s)

	s.AddRecoverable("A", raa)
	s.AddRecoverable("B", rab)
	s.AddRecoverable("C", rac)

	t.Logf("(A) is '%v'.", raa.Action(TimeConsumingAction))
	t.Logf("(B) is '%v'.", rab.Action(PositiveAction))
	t.Logf("(C) is '%v'.", rac.Action(PositiveAction))
}

// Test the finite state machine successfully.
func TestFsmSuccess(t *testing.T) {
	fsm := NewFSM(NewLoginHandler(), -1)

	fsm.Send(&LoginPayload{"yadda"})
	fsm.Send(&PreparePayload{"foo", "bar"})
	fsm.Send(&LoginPayload{"yaddaA"})
	fsm.Send(&LoginPayload{"yaddaB"})
	fsm.Send(&LoginPayload{"yaddaC"})
	fsm.Send(&LoginPayload{"yaddaD"})
	fsm.Send(&UnlockPayload{})
	fsm.Send(&LoginPayload{"bar"})

	time.Sleep(1e7)

	t.Logf("Status: '%v'.", fsm.State())
}

// Test the finite state machine with timeout.
func TestFsmTimeout(t *testing.T) {
	fsm := NewFSM(NewLoginHandler(), 1e5)

	fsm.Send(&LoginPayload{"yadda"})
	fsm.Send(&PreparePayload{"foo", "bar"})
	fsm.Send(&LoginPayload{"yaddaA"})
	fsm.Send(&LoginPayload{"yaddaB"})

	time.Sleep(1e8)

	fsm.Send(&LoginPayload{"yaddaC"})
	fsm.Send(&LoginPayload{"yaddaD"})
	fsm.Send(&UnlockPayload{})
	fsm.Send(&LoginPayload{"bar"})

	time.Sleep(1e7)

	t.Logf("Status: '%v'.", fsm.State())
}

// Test dispatching.
func TestDispatching(t *testing.T) {
	tt := new(TT)

	v1, ok1 := Dispatch(tt, "Add", 4, 5)
	v2, ok2 := Dispatch(tt, "Add", 4, 5, 6)
	v3, ok3 := Dispatch(tt, "Mul", 4, 5, 6)
	v4, ok4 := Dispatch(tt, "Mul", 4, 5, 6, 7, 8)

	t.Logf("Add 1: %v / %v\n", v1, ok1)
	t.Logf("Add 2: %v / %v\n", v2, ok2)
	t.Logf("Mul 1: %v / %v\n", v3, ok3)
	t.Logf("Mul 2: %v / %v\n", v4, ok4)
}

// Test debug statement.
func TestDebug(t *testing.T) {
	Debug("Hello, I'm debugging %v!", "here")
}

// Test nanoseconds calculation.
func TestNanoseconds(t *testing.T) {
	t.Logf("Microseconds: %v\n", NsMicroseconds(4711))
	t.Logf("Milliseconds: %v\n", NsMilliseconds(4711))
	t.Logf("Seconds     : %v\n", NsSeconds(4711))
	t.Logf("Minutes     : %v\n", NsMinutes(4711))
	t.Logf("Hours       : %v\n", NsHours(4711))
	t.Logf("Days        : %v\n", NsDays(4711))
	t.Logf("Weeks       : %v\n", NsWeeks(4711))
}

// Test time containments.
func TestTimeContainments(t *testing.T) {
	now := time.UTC()
	years := []int64{2008, 2009, 2010}
	months := []int{3, 6, 9, 12}
	days := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	hours := []int{20, 21, 22, 23}
	minutes := []int{5, 10, 15, 20, 25, 30}
	seconds := []int{0, 15, 30, 45}
	weekdays := []int{time.Saturday, time.Sunday}

	t.Logf("Time is %s\n", now.Format(time.RFC822))
	t.Logf("Year in list    : %t\n", YearInList(now, years))
	t.Logf("Year in range   : %t\n", YearInRange(now, 2000, 2005))
	t.Logf("Month in list   : %t\n", MonthInList(now, months))
	t.Logf("Month in range  : %t\n", MonthInRange(now, 1, 6))
	t.Logf("Day in list     : %t\n", DayInList(now, days))
	t.Logf("Day in range    : %t\n", DayInRange(now, 15, 25))
	t.Logf("Hour in list    : %t\n", HourInList(now, hours))
	t.Logf("Hour in range   : %t\n", HourInRange(now, 9, 17))
	t.Logf("Minute in list  : %t\n", MinuteInList(now, minutes))
	t.Logf("Minute in range : %t\n", MinuteInRange(now, 0, 29))
	t.Logf("Second in list  : %t\n", SecondInList(now, seconds))
	t.Logf("Second in range : %t\n", SecondInRange(now, 30, 59))
	t.Logf("Weekday in list : %t\n", WeekdayInList(now, weekdays))
	t.Logf("Weekday in range: %t\n", WeekdayInRange(now, time.Monday, time.Friday))
}

// Test the UUID.
func TestUuid(t *testing.T) {
	uuids := make(map[string]bool)

	t.Logf("Start generating UUIDs ...")

	for i := 0; i < 1000000; i++ {
		uuid := NewUUID().String()

		if uuids[uuid] {
			t.Fatalf("UUID collision")
		}

		uuids[uuid] = true
	}

	t.Logf("Done generating UUIDs!")
}

// Test the creation of an identifier.
func TestIdentifier(t *testing.T) {
	// Type as identifier.
	var kvlf KeyValueLessFunc

	idp := TypeAsIdentifierPart(kvlf)

	if idp != "key-value-less-func" {
		t.Errorf("Identifier part for KeyValueLessFunc is wrong, returned '%v'!", idp)
	}

	idp = TypeAsIdentifierPart(NewUUID())

	if idp != "u-u-i-d" {
		t.Errorf("Identifier part for UUID is wrong, returned '%v'!", idp)
	}

	// Identifier.
	id := Identifier("One", 2, "three four")

	if id != "one:2:three-four" {
		t.Errorf("First identifier is wrong! Id: %v", id)
	}

	id = Identifier(2011, 6, 22, "One, two, or  three things.")

	if id != "2011:6:22:one-two-or-three-things" {
		t.Errorf("Second identifier is wrong! Id: %v", id)
	}

	id = SepIdentifier("+", 1, "oNe", 2, "TWO", "3", "ÄÖÜ")

	if id != "1+one+2+two+3+äöü" {
		t.Errorf("Third identifier is wrong! Id: %v", id)
	}

	id = LimitedSepIdentifier("+", true, "     ", 1, "oNe", 2, "TWO", "3", "ÄÖÜ", "Four", "+#-:,")

	if id != "1+one+2+two+3+four" {
		t.Errorf("Fourth identifier is wrong! Id: %v", id)
	}
}

// Test the integer generator.
func TestLazyIntEvaluator(t *testing.T) {
	fibFunc := func(s interface{}) (interface{}, interface{}) {
		os := s.([]int)
		v1 := os[0]
		v2 := os[1]
		ns := []int{v2, v1 + v2}

		return v1, ns
	}

	fib := BuildLazyIntEvaluator(fibFunc, []int{0, 1})

	var fibs [25]int

	for i := 0; i < 25; i++ {
		fibs[i] = fib()
	}

	t.Logf("FIBS: %v", fibs)
}

// Test pivot.
func TestPivot(t *testing.T) {
	a := make(sort.IntSlice, 15)

	for i := 0; i < len(a); i++ {
		a[i] = rand.Intn(99)
	}

	plo, phi := partition(a, 0, len(a)-1)

	t.Logf("PLO  : %v", plo)
	t.Logf("PHI  : %v", phi)
	t.Logf("PDATA: %v", a[phi-1])
	t.Logf("PIVOT: %v", a)
}

// Test sort shootout.
func TestSort(t *testing.T) {
	ola := generateTestOrdersList(25000)
	olb := generateTestOrdersList(25000)
	olc := generateTestOrdersList(25000)
	old := generateTestOrdersList(25000)

	ta := time.Nanoseconds()
	Sort(ola)
	tb := time.Nanoseconds()
	sort.Sort(olb)
	tc := time.Nanoseconds()
	insertionSort(olc, 0, len(olc)-1)
	td := time.Nanoseconds()
	sequentialQuickSort(old, 0, len(olc)-1)
	te := time.Nanoseconds()

	t.Logf("PQS: %v", tb-ta)
	t.Logf(" QS: %v", tc-tb)
	t.Logf(" IS: %v", td-tc)
	t.Logf("SQS: %v", te-td)
}

// Test the parallel quicksort function.
func TestParallelQuickSort(t *testing.T) {
	ol := generateTestOrdersList(10000)

	Sort(ol)

	cn := 0

	for _, o := range ol {
		if cn > o.CustomerNo {
			t.Errorf("Customer No %v in wrong order!", o.CustomerNo)

			cn = o.CustomerNo
		} else {
			cn = o.CustomerNo
		}
	}
}

// Test the MapReduce function.
func TestMapReduce(t *testing.T) {
	// Start data producer.

	orderChan := generateTestOrders(2000)

	// Define map and reduce functions.

	mapFunc := func(in *KeyValue, mapEmitChan KeyValueChan) {
		o := in.Value.(*Order)

		// Emit analysis data for each item.

		for _, i := range o.Items {
			unitDiscount := (i.UnitPrice / 100.0) * i.DiscountPerc
			totalDiscount := unitDiscount * float64(i.Count)
			totalAmount := (i.UnitPrice - unitDiscount) * float64(i.Count)
			analysis := &OrderItemAnalysis{i.ArticleNo, i.Count, totalAmount, totalDiscount}
			articleNo := strconv.Itoa(i.ArticleNo)

			mapEmitChan <- &KeyValue{articleNo, analysis}
		}
	}

	reduceFunc := func(inChan KeyValueChan, reduceEmitChan KeyValueChan) {
		memory := make(map[string]*OrderItemAnalysis)

		// Collect emitted analysis data.

		for kv := range inChan {
			analysis := kv.Value.(*OrderItemAnalysis)

			if existing, ok := memory[kv.Key]; ok {
				existing.Quantity += analysis.Quantity
				existing.Amount += analysis.Amount
				existing.Discount += analysis.Discount
			} else {
				memory[kv.Key] = analysis
			}
		}

		// Emit it to map/reduce caller.

		for articleNo, analysis := range memory {
			reduceEmitChan <- &KeyValue{articleNo, analysis}
		}
	}

	// Now call MapReduce.

	for result := range SortedMapReduce(orderChan, mapFunc, 100, reduceFunc, 20, KeyLessFunc) {
		t.Logf("%v\n", result.Value)
	}
}

// Test job.
func TestJob(t *testing.T) {
	// Check function.

	cf := func(now *time.Time) (perform, delete bool) {
		perform = now.Day == 1 &&
			now.Hour == 22 &&
			now.Minute == 0 &&
			SecondInList(now, []int{0, 10, 20, 30, 40, 50})
		delete = false

		return perform, delete
	}

	// Task function.

	tf := func(id string) { t.Logf("Performed job %s\n", id) }

	// Job and time.

	job := NewJob("test-job-a", cf, tf)
	time := time.LocalTime()

	// Test with non-matching time.

	time.Second = 1

	job.checkAndPerform(time)

	// Test with matching time

	time.Day = 1
	time.Hour = 22
	time.Minute = 0
	time.Second = 0

	job.checkAndPerform(time)
}

// Test crontab keeping the job.
func TestCrontabKeep(t *testing.T) {
	ctb := NewCrontab()
	job := createJob(t, "keep", false)

	ctb.AddJob(job)
	time.Sleep(10 * 1e9)
	ctb.Stop()
}

// Test crontab deleting the job.
func TestCrontabDelete(t *testing.T) {
	ctb := NewCrontab()
	job := createJob(t, "delete", true)

	ctb.AddJob(job)
	time.Sleep(10 * 1e9)
	ctb.Stop()
}

// Test creating.
func TestSmlCreating(t *testing.T) {
	root := createSmlStructure()

	t.Logf("Root: %v", root)
}

// Test SML writer processing.
func TestSmlWriterProcessing(t *testing.T) {
	root := createSmlStructure()
	bufA := bytes.NewBufferString("")
	bufB := bytes.NewBufferString("")
	sppA := NewSmlWriterProcessor(bufA, true)
	sppB := NewSmlWriterProcessor(bufB, false)

	root.ProcessWith(sppA)
	root.ProcessWith(sppB)

	t.Logf("Print A: %v", bufA)
	t.Logf("Print B: %v", bufB)
}

// Test positive reading.
func TestSmlPositiveReading(t *testing.T) {
	sml := "Before!   {foo {bar:1:first Yadda ^{Test^} 1}  {inbetween}  {bar:2:last Yadda {Test ^^} 2}}   After!"
	reader := NewSmlReader(strings.NewReader(sml))

	root, err := reader.RootTagNode()

	if err == nil {
		t.Logf("Root:%v", root)
	} else {
		t.Errorf("Error: %v", err)
	}
}

// Test negative reading.
func TestSmlNegativeReading(t *testing.T) {
	sml := "{Foo {bar:1 Yadda {test} {} 1} {bar:2 Yadda 2}}"
	reader := NewSmlReader(strings.NewReader(sml))

	root, err := reader.RootTagNode()

	if err == nil {
		t.Errorf("Root: %v", root)
	} else {
		t.Logf("Error: %v", err)
	}
}

// Test of the ETM monitor.
func TestEtmMonitor(t *testing.T) {
	mon := Monitor()

	// Generate measurings.
	for i := 0; i < 500; i++ {
		n := rand.Intn(10)
		id := fmt.Sprintf("Work %d", n)
		m := mon.BeginMeasuring(id)

		work(n * 5000)

		m.EndMeasuring()
	}

	// Print, process with error, and print again.
	mon.MeasuringPointsPrintAll()

	mon.MeasuringPointsDo(func(mp *MeasuringPoint) {
		if mp.Count >= 25 {
			// Divide by zero.
			mp.Count = mp.Count / (mp.Count - mp.Count)
		}
	})

	mon.MeasuringPointsPrintAll()
}

// Test of the SSI monitor.
func TestSsiMonitor(t *testing.T) {
	mon := Monitor()

	// Generate values.
	for i := 0; i < 500; i++ {
		n := rand.Intn(10)
		id := fmt.Sprintf("Work %d", n)

		mon.SetValue(id, rand.Int63n(2001)-1000)
	}

	// Print, process with error, and print again.
	mon.StaySetVariablesPrintAll()

	mon.StaySetVariablesDo(func(ssv *StaySetVariable) {
		if ssv.Count >= 25 {
			// Divide by zero.
			ssv.Count = ssv.Count / (ssv.Count - ssv.Count)
		}
	})

	mon.StaySetVariablesPrintAll()
}

// Test of the DSR monitor.
func TestDsrMonitor(t *testing.T) {
	mon := Monitor()

	mon.Register("monitor:dsr:a", func() string { return "A" })
	mon.Register("monitor:dsr:b", func() string { return "4711" })
	mon.Register("monitor:dsr:c", func() string { return "2011-05-07" })

	mon.DynamicStatusValuesPrintAll()
}

//--------------------
// HELPERS
//--------------------

// Test type.
type TT struct{}

func (tt *TT) Add(a, b int) int { return a + b }

func (tt *TT) Mul(a, b, c, d, e int) int { return a * b * c * d * e }

// Order item type.
type OrderItem struct {
	ArticleNo    int
	Count        int
	UnitPrice    float64
	DiscountPerc float64
}

// Order type.
type Order struct {
	OrderNo    UUID
	CustomerNo int
	Items      []*OrderItem
}

func (o *Order) String() string {
	msg := "ON: %v / CN: %4v / I: %v"

	return fmt.Sprintf(msg, o.OrderNo, o.CustomerNo, len(o.Items))
}

// Order item analysis type.
type OrderItemAnalysis struct {
	ArticleNo int
	Quantity  int
	Amount    float64
	Discount  float64
}

func (oia *OrderItemAnalysis) String() string {
	msg := "AN: %5v / Q: %4v / A: %10.2f € / D: %10.2f €"

	return fmt.Sprintf(msg, oia.ArticleNo, oia.Quantity, oia.Amount, oia.Discount)
}

// Order list.
type OrderList []*Order

func (l OrderList) Len() int {
	return len(l)
}

func (l OrderList) Less(i, j int) bool {
	return l[i].CustomerNo < l[j].CustomerNo
}

func (l OrderList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

// Action function.
type Action func() string

// A positive action.
func PositiveAction() string {
	log.Printf("Perform positive action.")

	return "OK"
}

// A failing action.
func FailAction() string {
	log.Printf("Perform failing action!")

	panic("Fail!")

	return "Fail!"
}

// A time consuming action.
func TimeConsumingAction() string {
	log.Printf("Perform time consuming action!")

	time.Sleep(3e8)

	return "Time consumed"
}

// Recoverable action type.
type RecoverableAction struct {
	actionChan chan Action
	replyChan  chan string
	supervisor *Supervisor
	heartbeat  *Heartbeat
}

// Create a new recoverable action.
func NewRecoverableAction(supervisor *Supervisor) *RecoverableAction {
	ra := &RecoverableAction{
		actionChan: make(chan Action),
		replyChan:  make(chan string, 5),
		supervisor: supervisor,
	}

	ra.heartbeat = NewHeartbeat(ra, 1e8)

	go ra.backend()

	return ra
}

// Send an action to perform.
func (ra *RecoverableAction) Action(action Action) string {
	ra.actionChan <- action

	return <-ra.replyChan
}

// Implement Supervisor() of the recoverable interface.
func (ra *RecoverableAction) Supervisor() *Supervisor {
	return ra.supervisor
}

// Implement Recover() of the recoverable interface.
func (ra *RecoverableAction) Recover(r Recoverable, err interface{}) {
	if ra == r {
		log.Printf("Recovering error '%v'!", err)

		ra.replyChan <- "Recovered"

		go ra.backend()
	}
}

// Backend of the recoverable action.
func (ra *RecoverableAction) backend() {
	defer func() {
		HelpIfNeeded(ra, recover())
	}()

	for {
		select {
		case action := <-ra.actionChan:
			ra.replyChan <- action()
		case h := <-ra.heartbeat.HeartbeatChan:
			ImAlive(h)
		}
	}
}

// Generate a test order.
func generateTestOrders(count int) KeyValueChan {
	articleMaxNo := 9999
	unitPrices := make([]float64, articleMaxNo+1)

	for i := 0; i < articleMaxNo+1; i++ {
		unitPrices[i] = rand.Float64() * 100.0
	}

	kvc := make(KeyValueChan)

	go func() {
		for i := 0; i < count; i++ {
			order := new(Order)

			order.OrderNo = NewUUID()
			order.CustomerNo = rand.Intn(999) + 1
			order.Items = make([]*OrderItem, rand.Intn(9)+1)

			for j := 0; j < len(order.Items); j++ {
				articleNo := rand.Intn(articleMaxNo)

				order.Items[j] = &OrderItem{
					ArticleNo:    articleNo,
					Count:        rand.Intn(9) + 1,
					UnitPrice:    unitPrices[articleNo],
					DiscountPerc: rand.Float64() * 4.0,
				}
			}

			kvc <- &KeyValue{order.OrderNo.String(), order}
		}

		close(kvc)
	}()

	return kvc
}

// Generate a list with test orders.
func generateTestOrdersList(count int) OrderList {
	l := make(OrderList, count)
	idx := 0

	for kv := range generateTestOrders(count) {
		l[idx] = kv.Value.(*Order)

		idx++
	}

	return l
}

// Create a job that leads to a an event every 2 seconds.
func createJob(t *testing.T, descr string, delete bool) *Job {
	cf := func(now *time.Time) (bool, bool) { return now.Seconds()%2 == 0, delete }
	tf := func(id string) { t.Logf("Performed job %s\n", id) }

	return NewJob("test-server-"+descr, cf, tf)
}

// Create a SML structure.
func createSmlStructure() *TagNode {
	root := NewTagNode("root")

	root.AppendText("Text A")
	root.AppendText("Text B")

	root.AppendTaggedText("comment", "A first comment.")

	subA := root.AppendTag("sub-a:1st:important")

	subA.AppendText("Text A.A")

	root.AppendTaggedText("comment", "A second comment.")

	subB := root.AppendTag("sub-b:2nd")

	subB.AppendText("Text B.A")
	subB.AppendTaggedText("raw", "Any raw text with {, }, and ^.")

	return root
}

// Do some work.
func work(n int) int {
	if n < 0 {
		return 0
	}

	return n * work(n-1)
}

//--------------------
// TEST LOGIN EVENT HANDLER
//--------------------

// Prepare payload.
type PreparePayload struct {
	userId   string
	password string
}

// Login payload.
type LoginPayload struct {
	password string
}

// Reset payload.
type ResetPayload struct{}

// Unlock payload.
type UnlockPayload struct{}

// Login handler tyoe.
type LoginHandler struct {
	userId              string
	password            string
	illegalLoginCounter int
	locked              bool
}

// Create a new login handler.
func NewLoginHandler() *LoginHandler {
	return new(LoginHandler)
}

// Return the initial state.
func (lh *LoginHandler) Init() string {
	return "New"
}

// Terminate the handler.
func (lh *LoginHandler) Terminate(string, interface{}) string {
	return "LoggedIn"
}

// Handler for state: "New".
func (lh *LoginHandler) HandleStateNew(c *Condition) (string, interface{}) {
	switch pld := c.Payload.(type) {
	case *PreparePayload:
		lh.userId = pld.userId
		lh.password = pld.password
		lh.illegalLoginCounter = 0
		lh.locked = false

		log.Printf("User '%v' prepared.", lh.userId)

		return "Authenticating", nil
	case *LoginPayload:
		log.Printf("Illegal login, handler not initialized!")

		return "New", false
	case Timeout:
		log.Printf("Timeout, terminate handler!")

		return "Terminate", nil
	}

	log.Printf("Illegal payload '%v' during state 'new'!", c.Payload)

	return "New", nil
}

// Handler for state: "Authenticating".
func (lh *LoginHandler) HandleStateAuthenticating(c *Condition) (string, interface{}) {
	switch pld := c.Payload.(type) {
	case *LoginPayload:
		if pld.password == lh.password {
			lh.illegalLoginCounter = 0
			lh.locked = false

			log.Printf("User '%v' logged in.", lh.userId)

			return "Terminate", true
		}

		log.Printf("User '%v' used illegal password.", lh.userId)

		lh.illegalLoginCounter++

		if lh.illegalLoginCounter == 3 {
			lh.locked = true

			log.Printf("User '%v' locked!", lh.userId)

			return "Locked", false
		}

		return "Authenticating", false
	case *UnlockPayload:
		log.Printf("No need to unlock user '%v'!", lh.userId)

		return "Authenticating", nil
	case *ResetPayload, Timeout:
		lh.illegalLoginCounter = 0
		lh.locked = false

		log.Printf("User '%v' resetted.", lh.userId)

		return "Authenticating", nil
	}

	log.Printf("Illegal payload '%v' during state 'authenticating'!", c.Payload)

	return "Authenticating", nil
}

// Handler for state: "Locked".
func (lh *LoginHandler) HandleStateLocked(c *Condition) (string, interface{}) {
	switch pld := c.Payload.(type) {
	case *LoginPayload:
		log.Printf("User '%v' login rejected, user is locked!", lh.userId)

		return "Locked", false
	case *ResetPayload, *UnlockPayload, Timeout:
		lh.illegalLoginCounter = 0
		lh.locked = false

		log.Printf("User '%v' resetted / unlocked.", lh.userId)

		return "Authenticating", nil
	}

	log.Printf("Illegal payload '%v' during state 'loacked'!", c.Payload)

	return "Locked", nil
}

/*
	EOF
*/
