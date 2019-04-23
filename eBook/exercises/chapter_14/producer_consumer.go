// goroutines2.go
package main

import "fmt"

// integer producer:
//译者注：numGen函数中第三个参数所指定的 chan<- 类型的相关知识在 14.2.11 “通道的方向”，链接如下：
//https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/14.2.md#14211-%E9%80%9A%E9%81%93%E7%9A%84%E6%96%B9%E5%90%91
func numGen(start, count int, out chan<- int) { 
	for i := 0; i < count; i++ {
		out <- start
		start = start + count
	}
	close(out)
}

// integer consumer:
func numEchoRange(in <-chan int, done chan<- bool) {
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
