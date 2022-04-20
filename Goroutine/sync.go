package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func syncFn() {
	wg.Add(2)

	go fun1()
	go fun2()

	fmt.Println("main进入阻塞状态，等待wg中的子goroutine结束")
	wg.Wait()
	fmt.Println("main阻塞状态解除")

}

func fun1() {
	for i :=1; i < 10; i++{
		fmt.Println("fun1....", i)
	}
	// 给wg等待中的执行的goroutine数量减1，等同于`Add(-1)`
	wg.Done()
}

func fun2() {
	defer wg.Done()
	for i:= 1; i <= 10; i++ {
		fmt.Println("fun2.....", i)
	}
}