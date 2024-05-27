package main

import "fmt"

func Sum(a, b int) int {
	return a + b
}

func main() {
	result := Sum(3, 5)
	fmt.Println("result:", result)
}
