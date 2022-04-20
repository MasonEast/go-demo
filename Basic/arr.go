package main

import "fmt"

// 数组是值类型
// 数组的长度是类型的一部分，所以数组不能调整大小，[5]int和[10]int是不同的类型

func arr() {

	a := [...]int{1, 2, 3, 4, 5, 7}
	b := a
	b[0] = 10

	for k, val := range a {
		fmt.Println(k, val)
	}
}
