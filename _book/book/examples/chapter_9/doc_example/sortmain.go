// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This package gives an example of how to use a custom package with interfaces   	
package main
     	
import (
     	"fmt"
     	"./sort"
)
    	
// sorting of slice of integers
func ints() {
    	data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
    	a := sort.IntArray(data)  //conversion to type IntArray
    	sort.Sort(a)
    	if !sort.IsSorted(a) {
    			panic("fail")
    	}
	    fmt.Printf("The sorted array is: %v\n", a)
}
    	
// sorting of slice of strings
func strings() {
    	data := []string{"monday", "friday", "tuesday", "wednesday", "sunday","thursday", "", "saturday"}
    	a := sort.StringArray(data)
    	sort.Sort(a)
    	if !sort.IsSorted(a) {
    			panic("fail")
		}
	    fmt.Printf("The sorted array is: %v\n", a)
}
    	
// a type which describes a day of the week
type day struct {
    	num        int
    	shortName  string
    	longName   string
}
    	
type dayArray struct {
    	data []*day
}
    	
func (p *dayArray) Len() int            { return len(p.data) }
func (p *dayArray) Less(i, j int) bool  { return p.data[i].num < p.data[j].num }
func (p *dayArray) Swap(i, j int)       { p.data[i], p.data[j] = p.data[j], p.data[i] }
    	
// sorting of custom type day
func days() {
    	Sunday :=    day{0, "SUN", "Sunday"}
    	Monday :=    day{1, "MON", "Monday"}
    	Tuesday :=   day{2, "TUE", "Tuesday"}
    	Wednesday := day{3, "WED", "Wednesday"}
    	Thursday :=  day{4, "THU", "Thursday"}
    	Friday :=    day{5, "FRI", "Friday"}
    	Saturday :=  day{6, "SAT", "Saturday"}
    	data := []*day{&Tuesday, &Thursday, &Wednesday, &Sunday, &Monday, &Friday, &Saturday}
    	a := dayArray{data}
    	sort.Sort(&a)
    	if !sort.IsSorted(&a) {
    			panic("fail")
    	}
    	for _, d := range data {
    			fmt.Printf("%s ", d.longName)
    	}
    	fmt.Printf("\n")
}
    	
    	
func main() {
    ints()
    strings()
    days()
}

/* Output:
The sorted array is: [-5467984 -784 0 0 42 59 74 238 905 959 7586 7586 9845]
The sorted array is: [ friday monday saturday sunday thursday tuesday wednesday]
Sunday Monday Tuesday Wednesday Thursday Friday Saturday 
*/