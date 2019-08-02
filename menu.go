package main

import (
	"fmt"

	"LeetCodeBook/util"
)

func menu() []string {
	fmt.Printf("> Question Number: ")
	order := util.GetChar()
	return order
}
