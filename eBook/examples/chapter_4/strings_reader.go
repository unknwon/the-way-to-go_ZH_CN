package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	var str string = "Hello, the-way-to-go_ZH_CN"
	rs := strings.NewReader(str)

	// 从rs中读取下一个byte
	ab, err := rs.ReadByte()
	if err != nil { // 关于更多异常处理的方法见5.2节
		log.Fatalf("ReadByte Failed: %v", err)
	}
	fmt.Println("读取到的byte值为: ", ab)

	aBytes := make([]byte, 5) // 初始化一个长度为5（len(aBytes)==5）的字节数组
	fmt.Println("字节数组aBytes初始化: ", aBytes)
	// 从rs中读取Max(len(aBytes), rs.Len())个byte到aBytes中
	fmt.Printf("len(aBytes): %d, rs.Len() %d \n", len(aBytes), rs.Len())
	n, err := rs.Read(aBytes)
	if err != nil {
		log.Fatalf("ReadBytes Failed: %v", err)
	}
	fmt.Printf("从rs中读取%d个byte \n", n)
	fmt.Printf("读取%d个byte到字节数组aBytes后 %v: \n", n, aBytes)
	fmt.Printf("len(aBytes): %d, rs.Len(): %d \n", len(aBytes), rs.Len())

	aBytes = make([]byte, 30)
	fmt.Println("字节数组aBytes初始化: ", aBytes)
	// 从rs中读取Max(len(aBytes), rs.Len())个byte到aBytes中
	fmt.Printf("len(aBytes): %d, rs.Len() %d \n", len(aBytes), rs.Len())
	n, err = rs.Read(aBytes)
	if err != nil {
		log.Fatalf("ReadBytes Failed: %v", err)
	}
	fmt.Printf("从rs中读取%d个byte \n", n)
	fmt.Printf("读取%d个byte到字节数组aBytes后 %v: \n", n, aBytes)
	fmt.Printf("len(aBytes): %d, rs.Len(): %d \n", len(aBytes), rs.Len())
}
