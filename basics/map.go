package main

import (
	"fmt"
)

func main() {
	employeeID := map[string]int{
		"steve": 100,
		"jamie": 500,
		"mike":  900,
	}
	fmt.Println("Contents of the map")
	for key, value := range employeeID {
		fmt.Printf("employeeID[%s] = %d\n", key, value)
	}

}
