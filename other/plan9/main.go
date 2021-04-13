package main

import "fmt"

func init() {
	fmt.Println("init")
}
func main() {
	fmt.Println("test")
	defer func() {
		fmt.Println("test1")
	}()
	defer func() {
		fmt.Println("test2")
	}()
	defer func() {
		fmt.Println("test3")
	}()
}
