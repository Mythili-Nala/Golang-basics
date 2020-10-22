package main

import "fmt"

func add(x int, y int) int {
	total := 0
	total = x + y
	return total
}

func sub(x int, y int) int {
	var ans int = 0
	ans = x - y
	return ans
}

func main() {

	sum := add(20, 30)
	fmt.Println(sum)
	subt := sub(40, 20)
	fmt.Println(subt)
}
