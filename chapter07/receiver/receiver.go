package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total      int
	lastUpdate time.Time
}

func (c *Counter) Increment() {
	if c == nil {
		fmt.Println("Warning: Increment called on nil counter")
		return
	}
	c.total++
	c.lastUpdate = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("{total: %d, lastUpdate: %v}", c.total, c.lastUpdate)
}

func main() {
	c := Counter{}
	fmt.Println(c.String())
	c.Increment() // (&c).Increment() と書かなくてもOK
	fmt.Println(c.String())

	var c2 *Counter
	fmt.Println(c2)
	// nilを引数にしてポインタレシーバを呼び出す
	c2.Increment()
	fmt.Println(c2)
	// nilを引数にして値レシーバを呼び出すとpanicが発生する
	fmt.Println(c2.String())
}
