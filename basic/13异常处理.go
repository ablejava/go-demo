package main

import (
	"errors"
	"fmt"
)

func test() (int, error) {
	fmt.Println("")
	return 1, errors.New("")
	//return 0, fmt.Errorf("半径小于%d", 0)
}

func returnError() {

	value, err := test()
	if err != nil {
		fmt.Println("程序异常")
		return
	}
	fmt.Println(value)
}