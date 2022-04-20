package main

import "fmt"

type Human struct {
	name string
	age int
}
// 方法也是函数，不同的是它是被一个接收者调用的函数
func (h Human)say(p string){
	fmt.Println(p)
}

func (h *Human)reName(name string){
	h.name = name
	fmt.Println(name)
}

func methodFn() {
	man := Human{
		name: "mason",
		age: 18,
	}

	man.say("hello")
	man.reName("hh")

}