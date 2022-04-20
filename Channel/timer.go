package main

import (
	"fmt"
	"time"
)

func timerFn() {
	timer := time.NewTimer(3 * time.Second)
	fmt.Println(time.Now())

	// 此处在等待channel中的信号，执行此段代码会阻塞3秒
	ch := timer.C
	fmt.Println(<-ch)

	timer2Fn()
}

func timer2Fn() {
	timer2 := time.NewTimer(5 * time.Second)

	// 新开启一个线程来处理触发后的事件
	go func() {
		// 因为在下面的代码提前终止了计时器，所以这里接收不到，下面的代码就不会执行
		<- timer2.C
		fmt.Println("Timer2 over...")
	}()

	// 上面的等待信号是在新的线程，所以不会阻塞这里的代码执行
	time.Sleep(2 * time.Second)
	stop := timer2.Stop()

	if stop {
		fmt.Println("Timer2 stop...")
	}
}