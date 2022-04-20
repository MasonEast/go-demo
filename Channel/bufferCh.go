package main

import (
	"fmt"
	"strconv"
	"time"
)

func bufferChFn() {
	// ch := make(chan int)	// 非缓冲通道
	// ch2 := make(chan int, 5)	// 缓冲通道，缓存区大小为5
	ch3 := make(chan string, 4)

	go sendData2(ch3)

	for{
		v, ok := <- ch3
		time.Sleep(time.Second)
		if !ok {
			fmt.Println("读完了...")
			break
		}
		fmt.Println("读取的数据：", v)
	}

	
}

func sendData2(c chan string) {
	for i := 0; i < 10; i++ {
		c <- "数据" + strconv.Itoa(i)
		fmt.Println("子goroutine写出：", i)
	}
	close(c)
}