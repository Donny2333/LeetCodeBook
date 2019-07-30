package main

import (
	"fmt"

	"leetcode.com/algorithms/util"
)

func menu() []string {
	fmt.Printf("> Question Number: ")
	order := util.GetChar()
	return order
}
