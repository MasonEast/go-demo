package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func readFile () {
	dirname := "/Users/manbang/Desktop/go/go-demo"
	listFiles(dirname, 0)
}

func listFiles(dirname string, level int) {
	s := "|--"
	for i :=0; i < level; i++ {
		s = "|     " + s
	}
	fileInfos, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}

	for _, fi := range fileInfos {
		filename := dirname + "/" + fi.Name()
		fmt.Printf("%s%s\n", s, filename)
		if fi.IsDir(){
			listFiles(filename, level + 1)
		}
	}
}