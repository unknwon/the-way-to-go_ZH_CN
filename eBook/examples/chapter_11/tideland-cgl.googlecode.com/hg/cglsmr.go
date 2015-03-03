/*
	Tideland Common Go Library - Sorting and Map/Reduce

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
	"hash/adler32"
	"sort"
)

//--------------------
// CONTROL VALUES
//--------------------

// Threshold for switching from parallel to sequential quick sort.
var QuickSortParallelThreshold int = 4095

// Threshold for switching from sequential quick sort to insertion sort.
var QuickSortSequentialThreshold int = 63

//--------------------
// HELPING FUNCS
//--------------------

// Simple insertion sort for smaller data collections.
func insertionSort(data sort.Interface, lo, hi int) {
	for i := lo + 1; i < hi+1; i++ {
		for j := i; j > lo && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}

// Get the median based on Tukey's ninther.
func median(data sort.Interface, lo, hi int) int {
	m := (lo + hi) / 2
	d := (hi - lo) / 8

	// Move median into the middle.

	mot := func(ml, mm, mh int) {
		if data.Less(mm, ml) {
			data.Swap(mm, ml)
		}
		if data.Less(mh, mm) {
			data.Swap(mh, mm)
		}
		if data.Less(mm, ml) {
			data.Swap(mm, ml)
		}
	}

	// Get low, middle, and high median.

	if hi-lo > 40 {
		mot(lo+d, lo, lo+2*d)
		mot(m-d, m, m+d)
		mot(hi-d, hi, hi-2*d)
	}

	// Get combined median.

	mot(lo, m, hi)

	return m
}

// Partition the data based on the median.
func partition(data sort.Interface, lo, hi int) (int, int) {
	med := median(data, lo, hi)
	idx := lo

	data.Swap(med, hi)

	for i := lo; i < hi; i++ {
		if data.Less(i, hi) {
			data.Swap(i, idx)

			idx++
		}
	}

	data.Swap(idx, hi)

	return idx - 1, idx + 1
}

// Sequential quicksort using itself recursively.
func sequentialQuickSort(data sort.Interface, lo, hi int) {
	if hi-lo > QuickSortSequentialThreshold {
		// Use sequential quicksort.

		plo, phi := partition(data, lo, hi)

		sequentialQuickSort(data, lo, plo)
		sequentialQuickSort(data, phi, hi)
	} else {
		// Use insertion sort.

		insertionSort(data, lo, hi)
	}
}

// Parallel quicksort using itself recursively
// and concurrent.
func parallelQuickSort(data sort.Interface, lo, hi int, done chan bool) {
	if hi-lo > QuickSortParallelThreshold {
		// Parallel QuickSort.

		plo, phi := partition(data, lo, hi)
		partDone := make(chan bool)

		go parallelQuickSort(data, lo, plo, partDone)
		go parallelQuickSort(data, phi, hi, partDone)

		// Wait for the end of both sorts.

		<-partDone
		<-partDone
	} else {
		// Sequential QuickSort.

		sequentialQuickSort(data, lo, hi)
	}

	// Signal that it's done.

	done <- true
}

//--------------------
// PARALLEL QUICKSORT
//--------------------

func Sort(data sort.Interface) {
	done := make(chan bool)

	go parallelQuickSort(data, 0, data.Len()-1, done)

	<-done
}

//--------------------
// BASIC KEY/VALUE TYPES
//--------------------

// Data processing is based on key/value pairs.
type KeyValue struct {
	Key   string
	Value interface{}
}

// Channel for the transfer of key/value pairs.
type KeyValueChan chan *KeyValue

// Slice of key/value channels.
type KeyValueChans []KeyValueChan

// Map a key/value pair, emit to the channel.
type MapFunc func(*KeyValue, KeyValueChan)

// Reduce the key/values of the first channel, emit to the second channel.
type ReduceFunc func(KeyValueChan, KeyValueChan)

// Channel for closing signals.
type SigChan chan bool

//--------------------
// HELPING FUNCS
//--------------------

// Close given channel after a number of signals.
func closeSignalChannel(kvc KeyValueChan, size int) SigChan {
	sigChan := make(SigChan)

	go func() {
		ctr := 0

		for {
			<-sigChan

			ctr++

			if ctr == size {
				close(kvc)

				return
			}
		}
	}()

	return sigChan
}

// Perform the reducing.
func performReducing(mapEmitChan KeyValueChan, reduceFunc ReduceFunc, reduceSize int, reduceEmitChan KeyValueChan) {
	// Start a closer for the reduce emit chan.

	sigChan := closeSignalChannel(reduceEmitChan, reduceSize)

	// Start reduce funcs.

	reduceChans := make(KeyValueChans, reduceSize)

	for i := 0; i < reduceSize; i++ {
		reduceChans[i] = make(KeyValueChan)

		go func(inChan KeyValueChan) {
			reduceFunc(inChan, reduceEmitChan)

			sigChan <- true
		}(reduceChans[i])
	}

	// Read map emitted data.

	for kv := range mapEmitChan {
		hash := adler32.Checksum([]byte(kv.Key))
		idx := hash % uint32(reduceSize)

		reduceChans[idx] <- kv
	}

	// Close reduce channels.

	for _, reduceChan := range reduceChans {
		close(reduceChan)
	}
}

// Perform the mapping.
func performMapping(mapInChan KeyValueChan, mapFunc MapFunc, mapSize int, mapEmitChan KeyValueChan) {
	// Start a closer for the map emit chan.

	sigChan := closeSignalChannel(mapEmitChan, mapSize)

	// Start mapping goroutines.

	mapChans := make(KeyValueChans, mapSize)

	for i := 0; i < mapSize; i++ {
		mapChans[i] = make(KeyValueChan)

		go func(inChan KeyValueChan) {
			for kv := range inChan {
				mapFunc(kv, mapEmitChan)
			}

			sigChan <- true
		}(mapChans[i])
	}

	// Dispatch input data to map channels.

	idx := 0

	for kv := range mapInChan {
		mapChans[idx%mapSize] <- kv

		idx++
	}

	// Close mapping channels channel.

	for i := 0; i < mapSize; i++ {
		close(mapChans[i])
	}
}

//--------------------
// MAP/REDUCE
//--------------------

// Simple map/reduce function.
func MapReduce(inChan KeyValueChan, mapFunc MapFunc, mapSize int, reduceFunc ReduceFunc, reduceSize int) KeyValueChan {
	mapEmitChan := make(KeyValueChan)
	reduceEmitChan := make(KeyValueChan)

	// Perform operations.

	go performReducing(mapEmitChan, reduceFunc, reduceSize, reduceEmitChan)
	go performMapping(inChan, mapFunc, mapSize, mapEmitChan)

	return reduceEmitChan
}

//--------------------
// RESULT SORTING
//--------------------

// Less function for sorting.
type KeyValueLessFunc func(*KeyValue, *KeyValue) bool

// Sortable set of key/value pairs.
type SortableKeyValueSet struct {
	data     []*KeyValue
	lessFunc KeyValueLessFunc
}

// Constructor for the sortable set.
func NewSortableKeyValueSet(kvChan KeyValueChan, kvLessFunc KeyValueLessFunc) *SortableKeyValueSet {
	s := &SortableKeyValueSet{
		data:     make([]*KeyValue, 0, 1024),
		lessFunc: kvLessFunc,
	}

	for kv := range kvChan {
		l := len(s.data)

		if l == cap(s.data) {
			tmp := make([]*KeyValue, l, l+1024)

			copy(tmp, s.data)

			s.data = tmp
		}

		s.data = s.data[0 : l+1]
		s.data[l] = kv
	}

	return s
}

// Sort interface: Return the len of the data.
func (s *SortableKeyValueSet) Len() int {
	return len(s.data)
}

// Sort interface: Return which element is less.
func (s *SortableKeyValueSet) Less(a, b int) bool {
	return s.lessFunc(s.data[a], s.data[b])
}

// Sort interface: Swap two elements.
func (s *SortableKeyValueSet) Swap(a, b int) {
	s.data[a], s.data[b] = s.data[b], s.data[a]
}

// Return the data using a channel.
func (s *SortableKeyValueSet) DataChan() KeyValueChan {
	kvChan := make(KeyValueChan)

	go func() {
		for _, kv := range s.data {
			kvChan <- kv
		}

		close(kvChan)
	}()

	return kvChan
}

// SortedMapReduce performes a map/reduce and sorts the result.
func SortedMapReduce(inChan KeyValueChan, mapFunc MapFunc, mapSize int, reduceFunc ReduceFunc, reduceSize int, lessFunc KeyValueLessFunc) KeyValueChan {
	kvChan := MapReduce(inChan, mapFunc, mapSize, reduceFunc, reduceSize)
	s := NewSortableKeyValueSet(kvChan, lessFunc)

	Sort(s)

	return s.DataChan()
}

// KeyLessFunc compares the keys of two key/value
// pairs. It returns true if the key of a is less
// the key of b.
func KeyLessFunc(a *KeyValue, b *KeyValue) bool {
	return a.Key < b.Key
}

/*
	EOF
*/
