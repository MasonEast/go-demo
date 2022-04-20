package main

import "fmt"

// Map是一种无序的键值对的集合，使用hash表实现的，是引用类型；

func mapFn() {
	a := map[string]float32{"b": 1, "c": 2}

	fmt.Println("map-----------")
	for v := range a {
		fmt.Println(v)
	}

	delete(a, "b")
	fmt.Println(a)

	c, ok := a["c"]
	fmt.Println(c, ok)

	b, ok := a["b"]
	fmt.Println(b, ok)

}
