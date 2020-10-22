package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Println("Enter a number")
	fmt.Scan(&num)
	switch {
	case num >= 0 && num <= 50:
		fmt.Printf("%d is greater than 0 and less than 50", num)
	case num >= 51 && num <= 100:
		fmt.Printf("%d is greater than 51 and less than 100", num)
	case num >= 101:
		fmt.Printf("%d is greater than 100", num)
	}

}
