package main

import "fmt"

// 接口指定了类型应该具有的方法，类型决定了如何实现这些方法。
// 接口对象不能调用接口实现对象的属性

type Phone interface {
	call()
}

type Xiaomi struct {}

func (x Xiaomi) call() {
	fmt.Println("I am xiaomi")
}

type Iphone struct {}

func (i Iphone) call() {
	fmt.Println("I am Iphone")
}

func intefaceFn() {
	var phone Phone
	phone = new(Xiaomi)
	phone.call()

	phone = new(Iphone)
	phone.call()
}