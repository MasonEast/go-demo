package main

import "fmt"

func channelFn() {
	ch1 := make(chan bool)

	go func() {
		for i :=0; i < 10; i++ {
			fmt.Println("子goroutine中i:", i)
		}
		// 阻塞，直到循环结束后，向通道中写数据，表示要结束了
		ch1 <- true

		fmt.Println("write over...")
	}()

	fmt.Println("main....")
	// 阻塞，直到从通道中读取到数据
	data := <- ch1
	fmt.Println("read over:", data)
}