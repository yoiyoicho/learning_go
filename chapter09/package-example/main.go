package main

import (
	"fmt"
	print "package-example/formatter"
	"package-example/math"
)

func main() {
	num := math.Double(2)
	output := print.Format(num)
	fmt.Println(output)
}
