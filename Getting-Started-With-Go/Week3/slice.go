package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var input string
	sli := make([]int, 0, 3)

	for input != "X" {
		fmt.Printf("Enter an integer or 'X' to quit: ")
		_, err := fmt.Scan(&input)

		if input != "X" {
			inputInt, _ := strconv.Atoi(input)
			fmt.Println("Entered integer", inputInt)
			sli = append(sli, inputInt)
			sort.Ints(sli)
			printSlice("Slice:", sli)
		}

		// If scan returns error
		if err != nil {
			fmt.Println("Error:", err)
		}
	}

	fmt.Println("Quit successful")
}

func printSlice(s string, x []int) {
	fmt.Printf("%v\n", x)
}
