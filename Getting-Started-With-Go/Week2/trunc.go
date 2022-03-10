package main

import "fmt"

func main() {
	var input float64;
	fmt.Printf("Enter a float: ");
	_, err := fmt.Scan(&input);
    fmt.Println("The truncated number is: ", int64(input))

	// If scan returns error
	if err != nil {
		fmt.Println("Error:", err)
	}
}
