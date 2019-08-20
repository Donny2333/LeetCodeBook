package main

import (
	"fmt"

	"github.com/Donny2333/LeetCodeBook/util"
)

func menu() []string {
	fmt.Printf("> Question Number: ")
	order := util.GetChar()
	return order
}
