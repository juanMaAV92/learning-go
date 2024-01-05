package main

import "fmt"

func sum(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func getValues(x int) (double int, triple int, quad int) {
	double = 2 * x
	triple = 3 * x
	quad = 4 * x
	return
}

func main() {
	fmt.Println(sum(2, 2))
	fmt.Println(sum(2, 2, 3, 10))

	fmt.Println(getValues(2))
}
