package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("welcome to find max and min using linear search")

	elements := []int{12, 33, -11, 30000, 1, 534, 66}

	fmt.Println("find Min:", FindMin(elements))
	fmt.Println("find Min:", FindMax(elements))

}

func FindMin(elements []int) int {
	min := math.MinInt
	if len(elements) == 0 {
		return min
	}
	min = elements[0]

	for i := 1; i < len(elements); i++ {
		if elements[i] < min {
			min = elements[i]
		}
	}

	return min
}

func FindMax(elements []int) int {
	max := math.MaxInt

	if len(elements) == 0 {
		return max
	}

	max = elements[0]
	for i := 1; i < len(elements); i++ {
		if elements[i] > max {
			max = elements[i]
		}
	}
	return max
}
