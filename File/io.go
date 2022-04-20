package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func ioFn() {
	fileName := "/Users/manbang/Desktop/go/go-demo/File/file.txt"
	fileName2 := "/Users/manbang/Desktop/go/go-demo/File/file2.txt"
	file, err := os.Open(fileName)
	if err !=nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	bs := make([]byte, 4)

	n := -1
	for{
		n, err = file.Read(bs)
		if n == 0 || err == io.EOF{
			fmt.Println("读取到文件末尾了")
			break
		}
		fmt.Println(string(bs[:n]))
	}

	file2, err := os.OpenFile(fileName2, os.O_CREATE | os.O_APPEND | os.O_WRONLY, os.ModePerm)
	HandleErr(err)
	defer file2.Close()

	// 直接写字符串
	// w, err := file2.WriteString("Hello, World!")
	// fmt.Println(w)
	// HandleErr(err)

	// 写入字节
	// n,err =file.Write([]byte("today"))

	// copy
	io.Copy(file2, file)
}

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}