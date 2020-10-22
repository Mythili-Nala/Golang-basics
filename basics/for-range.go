package main

import "fmt"

func main() {

	// Example 1
	strName := map[string]string{"Tina": "Manu", "Hina": "Benni", "Caterina": "Gita"}
	for index, element := range strName {
		fmt.Println("Index :", index, " Element :", element)
	}

	// Example 2
	for key := range strName {
		fmt.Println(key)
	}

	// Example 3
	for _, value := range strName {
		fmt.Println(value)
	}
}
