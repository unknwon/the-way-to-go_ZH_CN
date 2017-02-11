package main

import (
	"fmt"
)

type Any interface{}
type EvalFunc func(Any) (Any, Any)

func main() {
	fibFunc := func(state Any) (Any, Any) {
		os := state.([]uint64)
		v1 := os[0]
		v2 := os[1]
		ns := []uint64{v2, v1 + v2}
		return v1, ns
	}
	fib := BuildLazyUInt64Evaluator(fibFunc, []uint64{0, 1})

	for i := 0; i < 10; i++ {
		fmt.Printf("Fib nr %v: %v\n", i, fib())
	}
}

func BuildLazyEvaluator(evalFunc EvalFunc, initState Any) func() Any {
	retValChan := make(chan Any)
	loopFunc := func() {
		var actState Any = initState
		var retVal Any
		for {
			retVal, actState = evalFunc(actState)
			retValChan <- retVal
		}
	}
	retFunc := func() Any {
		return <-retValChan
	}
	go loopFunc()
	return retFunc
}

func BuildLazyUInt64Evaluator(evalFunc EvalFunc, initState Any) func() uint64 {
	ef := BuildLazyEvaluator(evalFunc, initState)
	return func() uint64 {
		return ef().(uint64)
	}
}

/* Output:
Fib nr 0: 0
Fib nr 1: 1
Fib nr 2: 1
Fib nr 3: 2
Fib nr 4: 3
Fib nr 5: 5
Fib nr 6: 8
Fib nr 7: 13
Fib nr 8: 21
Fib nr 9: 34
*/
