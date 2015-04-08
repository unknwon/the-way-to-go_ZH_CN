/*
	Tideland Common Go Library

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

// The package 'cgl' contains a lot of helpful stuff for many kinds of software. Those
// are a map/reduce, sorting, time handling, UUIDs, lazy evaluation, chronological jobs,
// state machines, monitoring, suervision, a simple markup language and much more.
package cgl

//--------------------
// IMPORTS
//--------------------

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"encoding/hex"
	"io"
	"log"
	"reflect"
	"runtime"
	"strings"
	"unicode"
)

//--------------------
// CONST
//--------------------

const RELEASE = "Tideland Common Go Library Release 2011-08-02"

//--------------------
// DEBUGGING
//--------------------

// Debug prints a debug information to the log with file and line.
func Debug(format string, a ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	info := fmt.Sprintf(format, a...)

	log.Printf("[cgl] debug %s:%d %v", file, line, info)
}

//--------------------
// UUID
//--------------------

// UUID represent a universal identifier with 16 bytes.
type UUID []byte

// NewUUID generates a new UUID based on version 4.
func NewUUID() UUID {
	uuid := make([]byte, 16)

	_, err := io.ReadFull(rand.Reader, uuid)

	if err != nil {
		log.Fatal(err)
	}

	// Set version (4) and variant (2).

	var version byte = 4 << 4
	var variant byte = 2 << 4

	uuid[6] = version | (uuid[6] & 15)
	uuid[8] = variant | (uuid[8] & 15)

	return uuid
}

// Raw returns a copy of the UUID bytes.
func (uuid UUID) Raw() []byte {
	raw := make([]byte, 16)

	copy(raw, uuid[0:16])

	return raw
}

// String returns a hexadecimal string representation with
// standardized separators.
func (uuid UUID) String() string {
	base := hex.EncodeToString(uuid.Raw())

	return base[0:8] + "-" + base[8:12] + "-" + base[12:16] + "-" + base[16:20] + "-" + base[20:32]
}

//--------------------
// MORE ID FUNCTIONS
//--------------------

// LimitedSepIdentifier builds an identifier out of multiple parts, 
// all as lowercase strings and concatenated with the separator
// Non letters and digits are exchanged with dashes and
// reduced to a maximum of one each. If limit is true only
// 'a' to 'z' and '0' to '9' are allowed.
func LimitedSepIdentifier(sep string, limit bool, parts ...interface{}) string {
	iparts := make([]string, 0)

	for _, p := range parts {
		tmp := strings.Map(func(r int) int {
			// Check letter and digit.
			if unicode.IsLetter(r) || unicode.IsDigit(r) {
				lcr := unicode.ToLower(r)

				if limit {
					// Only 'a' to 'z' and '0' to '9'.
					if lcr <= unicode.MaxASCII {
						return lcr
					} else {
						return ' '
					}
				} else {
					// Every char is allowed.
					return lcr
				}
			}

			return ' '
		}, fmt.Sprintf("%v", p))

		// Only use non-empty identifier parts.
		if ipart := strings.Join(strings.Fields(tmp), "-"); len(ipart) > 0 {
			iparts = append(iparts, ipart)
		}
	}

	return strings.Join(iparts, sep)
}

// SepIdentifier builds an identifier out of multiple parts, all
// as lowercase strings and concatenated with the separator
// Non letters and digits are exchanged with dashes and
// reduced to a maximum of one each.
func SepIdentifier(sep string, parts ...interface{}) string {
	return LimitedSepIdentifier(sep, false, parts...)
}

// Identifier works like SepIdentifier but the seperator
// is set to be a colon.
func Identifier(parts ...interface{}) string {
	return SepIdentifier(":", parts...)
}

// TypeAsIdentifierPart transforms the name of the arguments type into 
// a part for identifiers. It's splitted at each uppercase char, 
// concatenated with dashes and transferred to lowercase.
func TypeAsIdentifierPart(i interface{}) string {
	var buf bytes.Buffer

	fullTypeName := reflect.TypeOf(i).String()
	lastDot := strings.LastIndex(fullTypeName, ".")
	typeName := fullTypeName[lastDot+1:]

	for i, r := range typeName {
		if unicode.IsUpper(r) {
			if i > 0 {
				buf.WriteRune('-')
			}
		}

		buf.WriteRune(r)
	}

	return strings.ToLower(buf.String())
}

//--------------------
// METHOD DISPATCHING
//--------------------

// Dispatch a string to a method of a type.
func Dispatch(variable interface{}, name string, args ...interface{}) ([]interface{}, bool) {
	numArgs := len(args)
	value := reflect.ValueOf(variable)
	valueType := value.Type()
	numMethods := valueType.NumMethod()

	for i := 0; i < numMethods; i++ {
		method := valueType.Method(i)

		if (method.PkgPath == "") && (method.Type.NumIn() == numArgs+1) {

			if method.Name == name {
				// Prepare all args with variable and args.

				callArgs := make([]reflect.Value, numArgs+1)

				callArgs[0] = value

				for i, a := range args {
					callArgs[i+1] = reflect.ValueOf(a)
				}

				// Make the function call.

				results := method.Func.Call(callArgs)

				// Transfer results into slice of interfaces.

				allResults := make([]interface{}, len(results))

				for i, v := range results {
					allResults[i] = v.Interface()
				}

				return allResults, true
			}
		}
	}

	return nil, false
}

//--------------------
// LAZY EVALUATOR BUILDERS
//--------------------

// Function to evaluate.
type EvalFunc func(interface{}) (interface{}, interface{})

// Generic builder for lazy evaluators.
func BuildLazyEvaluator(evalFunc EvalFunc, initState interface{}) func() interface{} {
	retValChan := make(chan interface{})

	loopFunc := func() {
		var actState interface{} = initState
		var retVal interface{}

		for {
			retVal, actState = evalFunc(actState)

			retValChan <- retVal
		}
	}

	retFunc := func() interface{} {
		return <-retValChan
	}

	go loopFunc()

	return retFunc
}

// Builder for lazy evaluators with ints as result.
func BuildLazyIntEvaluator(evalFunc EvalFunc, initState interface{}) func() int {
	ef := BuildLazyEvaluator(evalFunc, initState)

	return func() int {
		return ef().(int)
	}
}

/*
	EOF
*/
