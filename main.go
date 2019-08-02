package main

import (
	"LeetCodeBook/util"
	"fmt"

	. "LeetCodeBook/solution/1.TwoSum"
	. "LeetCodeBook/solution/2.AddTwoNumbers"
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
			fmt.Println(TwoSum())
			break

		case "2":
			fmt.Println(AddTwoNumbers())
			break

		default:
			fmt.Println("Error Code !\n")
		}

	}

}
