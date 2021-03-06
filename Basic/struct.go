package main

import "fmt"

// 结构体大写字母开头可以导出，从其他包访问
// 结构体是值类型

type Books struct {
	title string
	author string
	book_id int
}

func structFn() {
	var Book1 Books
	Book1.title = "go语言"
	Book1.author = "mason"
	Book1.book_id = 1
	fmt.Println(Book1)

	var Book2 *Books = &Book1

	Book2.title = "go语言2"
	fmt.Println(Book1)

}

// 结构体的匿名字段
/*
type Human struct {
	name string
	age int
}

type Student struct {
	Human
	sex string
}
*/

// 结构体嵌套
/*
type Address struct {  
    city, state string
}
type Person struct {  
    name string
    age int
    address Address
}
*/

/*
make、new操作

make用于内建类型（map、slice 和channel）的内存分配。new用于各种类型的内存分配 内建函数new本质上说跟其它语言中的同名函数功能一样：new(T)分配了零值填充的T类型的内存空间，并且返回其地址，即一个*T类型的值。用Go的术语说，它返回了一个指针，指向新分配的类型T的零值。有一点非常重要：new返回指针

内建函数make(T, args)与new(T)有着不同的功能，make只能创建slice、map和channel，并且返回一个有初始值(非零)的T类型，而不是*T。本质来讲，导致这三个类型有所不同的原因是指向数据结构的引用在使用前必须被初始化。例如，一个slice，是一个包含指向数据（内部array）的指针、长度和容量的三项描述符；在这些项目被初始化之前，slice为nil。对于slice、map和channel来说，make初始化了内部的数据结构，填充适当的值。

make返回初始化后的（非零）值。
*/
