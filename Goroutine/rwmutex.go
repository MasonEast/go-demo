package main

import (
	"fmt"
	"sync"
	"time"
)

/*
同时只能有一个 goroutine 能够获得写锁定。
同时可以有任意多个 gorouinte 获得读锁定。
同时只能存在写锁定或读锁定（读和写互斥）。
*/

var rwMutex *sync.RWMutex
var wg2 *sync.WaitGroup

func rwmutexFn(){
	rwMutex = new(sync.RWMutex)
	wg2 = new(sync.WaitGroup)

	wg2.Add(3)
	go write(1)
	go read(2)
	go write(3)

	wg2.Wait()

	fmt.Println("程序结束")
}

func write(i int){
	defer wg2.Done()
	fmt.Println(i, "开始写")
	rwMutex.Lock()
	fmt.Println(i, "正在写")
	time.Sleep(3 * time.Second)
	rwMutex.Unlock()
	fmt.Println("写结束")
}

func read(i int){
	defer wg2.Done()

	fmt.Println(i, "开始读")

	rwMutex.RLock()
	fmt.Println("正在读")
	time.Sleep(3 * time.Second)

	rwMutex.RUnlock()
	fmt.Println("读结束")
}