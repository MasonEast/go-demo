package main

import "fmt"

func calcSquaresFn() {
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)

	go cal(number, sqrch)
	go cal(number, cubech)

	squares, cubes := <- sqrch, <- cubech

	fmt.Println("final output:", squares + cubes)
}

func cal(number int, ch chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	ch <- sum
}