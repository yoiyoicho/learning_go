package main

import (
	"fmt"
)

type Employee struct {
	Name string
	ID   int
}

type Manager struct {
	Employee            // 埋め込みフィールド
	Reports  []Employee // 部下
}

func (e Employee) String() string {
	return fmt.Sprintf("Name: %s, ID: %d", e.Name, e.ID)
}

func main() {
	e := Employee{Name: "Tanaka", ID: 100}
	fmt.Println(e.String())
	m := Manager{
		Employee: Employee{Name: "Suzuki", ID: 200},
		Reports:  []Employee{},
	}
	fmt.Println(m.String()) // 埋め込まれたフィールドのメソッドが使える
	var e2 Employee
	// e2 = m // コンパイルエラー
	e2 = m.Employee
	fmt.Println(e2.String())
}
