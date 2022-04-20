package main

import "fmt"

func stringFn() {

	fmt.Println("string-------")

	const str = "abcdefg"
	for k, v := range str {
		fmt.Println(k, v)
		fmt.Printf("%c", v)
	}
	fmt.Println(str[0])

}
