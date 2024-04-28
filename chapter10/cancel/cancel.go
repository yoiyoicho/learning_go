package main

import "fmt"

func countTo(max int) (<-chan int, func()) {
	ch := make(chan int)
	done := make(chan struct{})
	cancelFunc := func() {
		close(done)
	}
	go func() {
		for i := 0; i < max; i++ {
			select {
			case <-done:
				fmt.Printf("Done\n")
				return
			case ch <- i:
				fmt.Printf("Sending: %d\n", i)
			}
		}
		fmt.Printf("Closing channel\n")
		close(ch)
	}()
	return ch, cancelFunc
}

func main() {
	ch, cancelFunc := countTo(10)
	for i := range ch {
		fmt.Printf("Received: %d\n", i)
		if i > 5 {
			break
		}
	}
	fmt.Println()
	cancelFunc()
}
