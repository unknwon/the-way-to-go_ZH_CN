// goroutines2.go
package main

import "fmt"

// integer producer:
//译者注：numGen函数中第三个参数所指定的 chan<- 类型的相关知识在 “14.2.11 通道的方向”，链接如下：
func numGen(start, count int, out chan<- int) { 
	for i := 0; i < count; i++ {
		out <- start
		start = start + count
	}
	close(out)
}

// integer consumer:
func numEchoRange(in <-chan int, done chan<- bool) {
	//译者注：下面的for 循环所涉及的相关知识在 “14.2.10 给通道使用 for 循环”
	for num := range in {
		fmt.Printf("%d\n", num)
	}
	done <- true
}

func main() {
	numChan := make(chan int)
	done := make(chan bool)
	go numGen(0, 10, numChan)
	go numEchoRange(numChan, done)

	<-done
}

/* Output:
0
10
20
30
40
50
60
70
80
90
*/
