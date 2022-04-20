package main

import (
	"fmt"
	"runtime"
)

func goSchedfn() {
	go func() {
		for i := 0; i< 5; i++ {
			fmt.Println("goroutine...")
		}
	}()

	for i := 0; i < 4; i++ {
		if i % 2 != 0 {
			// 让出时间片，先让别的协程执行，别的执行完，再回来执行此协程
			runtime.Gosched()
		} 
		fmt.Println("main....")
	}
}