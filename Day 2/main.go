package main

/*
今天跑`go run main.go`报错，昨天还好好的，很奇怪。
网上查询结果：
Go 中 main 包默认不会加载其他文件， 而其他包都是默认加载的。
如果 main 包有多个文件，则在执行的时候需要将其它文件都带上，即执行 `go run *.go`。
*/


func main () {
	fileFn()
	ioFn()
	readFile()
}