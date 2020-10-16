package main

import "fmt"

func main() {
	rataKiri(8)
	rataKanan(8)

}

func rataKiri(rows int) {
	for i := 1; i <= rows; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Printf("\n")
	}
}

func rataKanan(rows int) {
	for i := 1; i <= rows; i++ {
		for j := 1; j <= rows; j++ {
			if i <= j {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Printf("\n")
	}

}
