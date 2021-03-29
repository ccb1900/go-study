package main

import (
	"fmt"
	"os"
	"reflect"
)

type demo struct {
}

type IDemo interface {
}

func main() {
	print(3)
	print("123")
	print(os.Stdout)
	print(make(chan string))
	print(make(map[int]string))
	print(make([]string, 10))
	print(true)
	print(1.22)
	print(demo{})
	print([2]int{12})
	print(func() {})
	var vv reflect.Value
	print(vv)
}

func print(a interface{}) {
	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)
	fmt.Println("type " + t.String())
	fmt.Printf("type is %T..\n", a)
	fmt.Printf("type is %T..\n", t)

	fmt.Println("#v", fmt.Sprintf("value is %v", v))
	fmt.Println("value=", v)
	kt := t.Kind()
	kv := v.Kind()
	fmt.Println("value string", v.String())
	fmt.Println("kind type==", kt.String(), "kind value", kv.String())
}
