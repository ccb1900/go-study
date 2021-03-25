package main

import "fmt"

func manyDefer() {
	fmt.Println("defer start")
	defer fmt.Println("first")
	defer fmt.Println("second")
	defer fmt.Println("third")
	fmt.Println("defer finished")
}
func main() {
	manyDefer()
}
