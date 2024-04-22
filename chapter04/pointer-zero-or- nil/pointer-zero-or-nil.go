package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type PersonPointer struct {
	Name string
	Age  *int
}

func main() {
	var age int // ゼロ値が設定されている変数

	// ポインタを使わない場合
	// ageが値なしなのか、ゼロ値なのか区別がつかない
	var p1 Person // { 0}
	p2 := Person{
		Name: "Taro",
		Age:  age,
	} // {Taro 0}

	fmt.Println(p1)
	fmt.Println(p2)

	// ポインタを使う場合
	// ageが値なしなのか、ゼロ値なのか区別できる
	var pp1 PersonPointer // { <nil>}
	pp2 := PersonPointer{
		Name: "Hanako",
		Age:  &age,
	} // {Hanako 0x14000112030}

	fmt.Println(pp1)
	fmt.Println(pp2)
}
