package main

import (
	"fmt"
	"reflect"
)

const c = "world"

func main() {
	x := "hello"
	pointerToX := &x
	fmt.Println(pointerToX)
	fmt.Println(reflect.TypeOf(pointerToX))
	// fmt.Println(&c) 定数はアドレスを取得できない
}
