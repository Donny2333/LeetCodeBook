package main

import (
	"fmt"

	"leetcode.com/algorithms/solution"
	"leetcode.com/algorithms/util"
)

func main() {
	util.CallClear()
	fmt.Println("---------------------------------------")
	fmt.Println()
	fmt.Println("\tWelcome to leetcode.com")
	fmt.Println()
	fmt.Println("---------------------------------------")

	for true {
		menu := menu()

		if menu[0] == "q" {
			fmt.Println("Bye ~\n")
			break
		}

		switch menu[0] {
		case "1":
			fmt.Println(solution.TwoSum())
			break

		case "2":
			fmt.Println(solution.AddTwoNumbers())
			break

		default:
			fmt.Println("Error Code !\n")
		}

	}

}
