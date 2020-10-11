package main

import "fmt"

func main() {
	pattern(5)
}

func pattern(rows int) {
	if rows%2 == 0 {
		fmt.Println("Must be odd number")
	} else {
		for i := 1; i <= rows; i++ {
			for j := 1; j <= rows; j++ {
				if i == (rows+1)/2 || j == 1 || j == rows {
					fmt.Print(" * ")
				} else {
					fmt.Print(" = ")
				}

			}
			fmt.Printf("\n")
		}
	}
}
