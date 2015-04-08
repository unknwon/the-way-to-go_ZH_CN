package fibo

/*
func Fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = Fibonacci(n-1) + Fibonacci(n-2)
	}
	return
}
*/
// accepts a general operation op:
func Fibonacci(op string, n int) (res int) {
	if n <= 1 {
		switch op {
			case "+":
				res = 1
			case "*":
				res = 2
			default: res = 0
		}
	} else {
		switch op {
			case "+":
				res = Fibonacci(op, n-1) + Fibonacci(op, n-2)
			case "*":
				res = Fibonacci(op, n-1) * Fibonacci(op, n-2)
			default: res = 0
		}
	}
	return
}
