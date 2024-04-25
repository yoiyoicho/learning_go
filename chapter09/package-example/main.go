package main

import (
	"fmt"
	"package-example/formatter"
	"package-example/math"
)

func main() {
	num := math.Double(2)
	output := formatter.Format(num)
	fmt.Println(output)
}
