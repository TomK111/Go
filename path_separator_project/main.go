package main

import (
	"fmt"
	"path"
)

func main() {
	// var file string
	// var dir string

	//dir, file = path.Split("css/main.css")

	_, file := path.Split("css/main.css")
	_, dir := path.Split("css/main.css")

	//fmt.Println("dir :", dir)

	fmt.Println("file :", file)
	fmt.Println("dir :", dir)
}
