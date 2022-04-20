package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
互斥锁
*/
var ticket = 10

var mutex sync.Mutex

func mutexFn() {
	wg.Add(4)
	go sale("售票口1")
	go sale("售票口2")
	go sale("售票口3")
	go sale("售票口4")

	wg.Wait()
	fmt.Println("程序结束")
}

func sale(name string) {
	rand.Seed(time.Now().UnixNano())
	defer wg.Done()

	for{
		mutex.Lock()
		if(ticket > 0){
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(name, "售出：", ticket)
			ticket--
		}else{
			mutex.Unlock()
			fmt.Println("卖完了")
			break
		}
		mutex.Unlock()
	}
}