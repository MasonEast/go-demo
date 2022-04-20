package main

import (
	"fmt"
)

/*
range循环可以用于冲通道接收值，直到通道关闭
*/

func rangeFn() {
	ch := make(chan int)
	go sendData(ch)

	for v:= range ch {
		fmt.Println("读取的数据：",v)
	}
	fmt.Println("main...")
}

func sendData(c chan int){
	for i:= 0; i < 10; i++ {
		// time.Sleep(time.Second)
		c <- i
	}
	// 通知对方，通道已关闭
	close(c)
}