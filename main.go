package main

import "fmt"

type MyData struct {
	Id   int
	Name string
}

func main() {
	fmt.Println("Hello, World!")
	x := MyData{Id: 1, Name: "test"}
	fmt.Println(x)
}
