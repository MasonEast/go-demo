package main

import "fmt"

/*
单向通道

	// 单向，只能写不能读
	ch := make(chan <- int)
	// 只能读不能写
	ch2 := make(<- chan int)
*/

func singleWayFn() {
	ch := make(chan int)

	go write(ch)

	data := <- ch

	fmt.Println("写出的数据是：",data)

	go read(ch)

	ch <- 200

	fmt.Println("singleWay main...")
}

func read(c <- chan int) {
	data := <- c
	fmt.Println("读到的数据:", data)
}

func write(c chan <- int) {
	c <- 100
	fmt.Println("写数据")
}