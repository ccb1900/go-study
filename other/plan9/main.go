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

	s := []int{2, 2, 2}
	fmt.Printf("out addr is %p", &s)
	f(s)
	fmt.Println(s)
}

func f(s []int) {
	fmt.Printf("func addr is %p", &s)
	for i, i2 := range s {
		fmt.Println(i, i2)
		s[i] += 2
	}
}
