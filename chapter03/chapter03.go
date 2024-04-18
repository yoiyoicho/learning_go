package main

import "fmt"

func main() {
	var x []int
	fmt.Println(x, len(x), cap(x))
	x = append(x, 1)
	fmt.Println(x, len(x), cap(x))
}
