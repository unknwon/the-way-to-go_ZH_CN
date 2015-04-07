// calculator.go
// 	example: calculate 3 + 4 = 7 as input: 3 ENTER 4 ENTER + ENTER --> result = 7, 

package main

import (
	"fmt"
	"strconv"
	"bufio"
	"os"
	"./stack/stack"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	calc1 := new(stack.Stack)
	fmt.Println("Give a number, an operator (+, -, *, /), or q to stop:")
	for {
		token, err := buf.ReadString('\n')
		if err != nil {
			fmt.Println("Input error !")
			os.Exit(1)
		}
		token = token[0:len(token)-2]    // remove "\r\n"
		// fmt.Printf("--%s--\n",token)  // debug statement
		switch  {
			case token == "q":  // stop als invoer = "q"
				fmt.Println("Calculator stopped")
				return  
			case token >= "0" && token <= "999999": 
				i, _ := strconv.Atoi(token)
				calc1.Push(i)  
			case token == "+":
				q := calc1.Pop()
				p := calc1.Pop()
				fmt.Printf("The result of %d %s %d = %d\n", p, token, q, p + q)
			case token == "-":
				q := calc1.Pop()
				p := calc1.Pop()
				fmt.Printf("The result of %d %s %d = %d\n", p, token, q, p - q)
			case token == "*":
				q := calc1.Pop()
				p := calc1.Pop()
				fmt.Printf("The result of %d %s %d = %d\n", p, token, q, p * q)
			case token == "/":
				q := calc1.Pop()
				p := calc1.Pop()
				fmt.Printf("The result of %d %s %d = %d\n", p, token, q, p / q)
			default:
				fmt.Println("No valid input")
		}
	}
}
