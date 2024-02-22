package piscine

import "fmt"

// QuadC function takes two integers x and y as parameters and prints a pattern based on their values.
func QuadC(x, y int) {
	// Check if both x and y are greater than 0
	if x > 0 && y > 0 {
		// Print 'A' if conditions are met
		fmt.Print("A")

		// Loop to print 'B' characters for x-2 times (excluding first and last columns)
		for i := 0; i < x-2; i++ {
			fmt.Print("B")
		}

		// Print 'A' in the last column if x is greater than 1, otherwise, print a newline
		if x > 1 {
			fmt.Println("A")
		} else {
			fmt.Println()
		}

		// Loop for printing rows (excluding first and last rows)
		for j := 0; j < y-2; j++ {
			// Print 'B' in the first column
			fmt.Print("B")

			// Loop to print spaces for x-2 times (excluding first and last columns)
			for i := 0; i < x-2; i++ {
				fmt.Print(" ")
			}

			// Print 'B' in the last column if x is greater than 1, otherwise, print a newline
			if x > 1 {
				fmt.Println("B")
			} else {
				fmt.Println()
			}
		}

		// Print 'C' in the first column of the last row if y is greater than 1
		if y > 1 {
			fmt.Print("C")

			// Loop to print 'B' characters for x-2 times (excluding first and last columns)
			for i := 0; i < x-2; i++ {
				fmt.Print("B")
			}

			// Print 'C' in the last column if x is greater than 1, otherwise, print a newline
			if x > 1 {
				fmt.Println("C")
			} else {
				fmt.Println()
			}
		}
	}
}
