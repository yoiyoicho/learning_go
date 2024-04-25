package main

import (
	"errors"
	"fmt"
)

func AFunctionThatReturnsAnError() error {
	return errors.New("An error occurred")
}

func main() {
	err := AFunctionThatReturnsAnError()
	var coder interface {
		Code() int
	}
	if errors.As(err, &coder) {
		fmt.Println(coder.Code())
	}
}
