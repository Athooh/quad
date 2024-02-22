package piscine

import "fmt"

// QuadD is a function that prints a pattern based on the input values x and y.
func QuadD(x, y int) {
	// Check if both x and y are greater than 0
	if x > 0 && y > 0 {
		// Print 'A'
		fmt.Print("A")
		// Print 'B' x-2 times
		for i := 0; i < x-2; i++ {
			fmt.Print("B")
		}
		// If x is greater than 1, print 'C', otherwise, print a newline
		if x > 1 {
			fmt.Println("C")
		} else {
			fmt.Println()
		}

		// Loop to print the middle rows
		for j := 0; j < y-2; j++ {
			// Print 'B'
			fmt.Print("B")
			// Print spaces (x-2 times)
			for i := 0; i < x-2; i++ {
				fmt.Print(" ")
			}
			// If x is greater than 1, print 'B', otherwise, print a newline
			if x > 1 {
				fmt.Println("B")
			} else {
				fmt.Println()
			}
		}

		// If y is greater than 1, print 'A' at the bottom
		if y > 1 {
			// Print 'A'
			fmt.Print("A")
			// Print 'B' x-2 times
			for i := 0; i < x-2; i++ {
				fmt.Print("B")
			}
			// If x is greater than 1, print 'C', otherwise, print a newline
			if x > 1 {
				fmt.Println("C")
			} else {
				fmt.Println()
			}
		}
	}
}
