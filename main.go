package main

import (
	"fmt"
)

func main() {
	numLen := 5000
	results := make(map[int]bool)

	for i := 1; i < numLen; i++ {
		value := i
		result := i
		
		for value > 0 {
			digit := value % 10
			value = value/10

			result += digit
		}

		results[result] = true
	}
	
	sumSelfNumbers := 0
	for i := 1; i < numLen; i++ {
		if !results[i] {
			sumSelfNumbers += i
		}
	}

	fmt.Println(sumSelfNumbers)
}