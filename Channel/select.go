package main

import (
	"fmt"
	"time"
)

/*
select类似于switch，但是select会随机执行一个可运行的case，如果没有case可以运行，它将阻塞，直到有case可以运行

- 每个case都必须是一个通道
- 所有channel表达式都会被求值
- 所有被发送的表达式都会被求值
- 如果有多个case可运行，select会随机公平选择一个执行，其他不会执行

*/

func selectFn() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 200
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- 100
	}()

	select {
	case num1 := <- ch1:
			fmt.Println("ch1中读取的数据：", num1)
	case num2, ok := <- ch2:
			if ok {
				fmt.Println("ch2中读取的数据：", num2)
			}else {
				fmt.Println("ch2通道已经关闭")
			}
	case <- time.After(3 * time.Second): 
		fmt.Println("case3...")
	default:
		fmt.Println("default")
	}
}