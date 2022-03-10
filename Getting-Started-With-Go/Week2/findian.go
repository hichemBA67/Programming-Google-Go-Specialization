package main

import (
	"fmt"
	"strings"
)
  

func main() {
	var input string;
	fmt.Printf("Enter a string: ");
	_, err := fmt.Scan(&input);
	if err != nil {
		fmt.Println("An error occured: ", err);
	}
	input = strings.ToLower(input)
	startsWithI := strings.HasPrefix(input, "i");
	endsWithN := strings.HasSuffix(input, "n");
	containsA := strings.Contains(input, "a");

	if startsWithI && endsWithN && containsA {
		fmt.Println("Found!")
	} else { 
		fmt.Println("Not found!")
	}

	
}
