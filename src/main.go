package main

import (
	"fmt"
)

func main() {
	fmt.Println("Calculator")

	var result = Sum(5, 5)

	fmt.Println(result)

}

func Sum(val1 int, val2 int) int {
	return val1 + val2
}
