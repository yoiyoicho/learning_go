package main

import (
	"fmt"
)

func doThings1(val1 int) (int, error) {
	return val1, nil
}

func doThings2(val11 int, val2 string) (string, error) {
	return fmt.Sprintf("doThings3(%d, %s)", val11, val2), nil
}

func DoSomeThings(val1 int, val2 string) (_ string, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("DoSomeThings failed: %w", err)
		}
	}()
	val3, err := doThings1(val1)
	if err != nil {
		return "", err
	}
	// 2つ目の戻り値はerrに代入される
	return doThings2(val3, val2)
}
