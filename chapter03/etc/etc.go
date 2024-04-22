package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x = make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4) // [1 2 3 4] cap5
	y := x[:2]                // [1 2] cap5
	z := x[2:]                // [3 4] cap3
	fmt.Println(cap(x), cap(y), cap(z))
	y = append(y, 30, 40, 50)
	fmt.Println("x:", x, "y:", y, "z:", z)
	// x: [1 2 30 40 50] y: [1 2 30 40 50] z: [30 40]
	x = append(x, 60)
	fmt.Println("x:", x, "y:", y, "z:", z)
	// x: [1 2 30 40 50 60] y: [1 2 30 40 50 60] z: [30 40]
	z = append(z, 70)
	fmt.Println("x:", x, "y:", y, "z:", z)
	// x: [1 2 30 40 70 60] y: [1 2 30 40 70 60] z: [30 40 70]

	x2 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(reflect.TypeOf(x2))
	y2 := x2[:2]
	fmt.Println(reflect.TypeOf(y2))

	x3 := []int{1, 2, 3, 4}
	d4 := [4]int{5, 6, 7, 8}
	copy(d4[:], x3)
	fmt.Println("x3:", x3, "d4:", d4)

	var nilMap map[string]int
	fmt.Println(nilMap)
	fmt.Println(nilMap == nil)
	totalMins := map[string]int{}
	fmt.Println(totalMins)
	fmt.Println(totalMins == nil)

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "3でも5でも割り切れる")
			continue
		}
		if i%3 == 0 {
			fmt.Println(i, "3で割り切れる")
			continue
		}
		if i%5 == 0 {
			fmt.Println(i, "5で割り切れる")
			continue
		}
		fmt.Println(i)
	}
}
