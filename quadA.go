package piscine

import "fmt"

// QuadA function takes two parameters x and y, representing the number of columns and rows, respectively.
func QuadA(x, y int) {
	// Check if x and y are both greater than 0
	if x > 0 && y > 0 {
		// Print the top side of the rectangle 'o'
		fmt.Print("o")
		// Prints middle elements of the top row by subtracting the width by 2
		for i := 0; i < x-2; i++ {
			fmt.Print("-")
		}
		// Print the last 'o' if there is more than one column, otherwise, print a new line
		if x > 1 {
			fmt.Println("o")
		} else {
			fmt.Println()
		}

		// Loop through the rows
		for j := 0; j < y-2; j++ {
			// Print the left side of the rectangle
			fmt.Print("|")
			// Print spaces for the length of row subtracting 2 elements  in the interior of the rectangle
			for i := 0; i < x-2; i++ {
				fmt.Print(" ")
			}
			// Print the right side of the rectangle if there is more than one column,
			// otherwise, print a new line
			if x > 1 {
				fmt.Println("|")
			} else {
				fmt.Println()
			}
		}

		// Print the bottom side of the rectangle if there is more than one row
		if y > 1 {
			fmt.Print("o")
			for i := 0; i < x-2; i++ {
				fmt.Print("-")
			}
			// Print the last 'o' if there is more than one column, otherwise, print a new line
			if x > 1 {
				fmt.Println("o")
			} else {
				fmt.Println()
			}
		}
	}
}
