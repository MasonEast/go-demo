package main

import "fmt"

func funcFn() {

	fmt.Println("func---------")

	res := max(1, 2)
	fmt.Println(res)

	a := 100
	fmt.Println(a)

	fn3(&a)
	fmt.Println(a)

	x, y := fn4("a", "b")
	fmt.Println(x, y)

	m, n := fn5(6, 7)
	fmt.Println(m, n)

	fn6()
}

func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

// 可变参
// func fn1(arg ...int){}

// 值传递
// func fn2(a int){}

// 引用传递
// 传指针可以使得多个函数能操作同一个对象
// 传指针只是传地址，比较轻量
func fn3(a *int) int {
	*a += *a
	return *a
}

// 函数多返回值
func fn4(x, y string) (string, string) {
	return y, x
}

func fn5(x, y int) (a int, b int) {
	a = x + y
	b = y - x
	return
}

// 延迟函数defer（先进后出）
// 延迟函数的参数是执行延迟语句时的参数，不是延迟函数调用时去获取
// 当执行外围函数中的return语句时，只有其中所有的延迟函数都执行完毕后，外围函数才会真正返回
func fn6() bool {
	a := 1
	b := 2
	defer fmt.Println(b)
	b = 4
	fmt.Println(a)
	return true
}
