package main

import "fmt"

// 切片本身没有任何数据，他们只是对现有数组的引用，与数组相比，不需要设定长度
// 切片元素的改变会影响底层数组；
// 引用类型

func sliceFn() {
	b := []int{1, 2, 3}
	a := b[1:]
	a[0] = 5

	fmt.Println(a, b)
	printSlice(a)

	a = append(a, 7, 8, 9)
	fmt.Println(a, b)
	printSlice(a)

	c := make([]int, len(a), cap(a)*2)
	copy(c, a)
	printSlice(c)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
