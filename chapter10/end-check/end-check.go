package main

import "fmt"

func countTo(max int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i <= max; i++ {
			fmt.Printf("Sending: %d\n", i)
			ch <- i
		}
		fmt.Printf("Closing channel\n")
		close(ch)
	}()
	return ch
}

func doSomethingTakingLongTime() {}

func main() {
	for i := range countTo(10) { // チャネルに値が送信されるとループが回る
		fmt.Printf("Received: %d\n", i)
		// 下記のようにループを途中で抜けると、doSomethingTakingLongTimeが実行されている間ずっと
		// 無名関数のゴルーチンは終了せずにブロックし続けることになる
		// if i > 5 {
		// 	break
		// }
	}
	fmt.Printf("Done\n")
	doSomethingTakingLongTime()
}
