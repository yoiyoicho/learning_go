package main

import (
	"fmt"
	"reflect"
)

type Score int
type HighScore Score

const ci = 5

func main() {
	var s Score = 10
	var hs HighScore = 30
	s = ci
	fmt.Println(s)
	hhs := hs + 20
	fmt.Println(hhs)
	fmt.Println(reflect.TypeOf(hhs))
	ns := s + Score(hs)
	fmt.Println(reflect.TypeOf(ns))
}
