package main

import (
	"fmt"
	"sort"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	people := []Person{
		{"Taro", "Yamada", 25},
		{"Hanako", "Tanaka", 30},
		{"Sachiko", "Yamagata", 35},
	}

	fmt.Println("初期データ")
	fmt.Println(people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].LastName < people[j].LastName
	})

	fmt.Println("姓でソート")
	fmt.Println(people)
}
