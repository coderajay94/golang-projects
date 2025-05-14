package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Hello World from Linear Search!!")

	elements := [9]int{12, 14, 1, 34, 61, 88, 56, -11, 11}

	fmt.Println("returns index:", LinearSearch(elements, 11))
	fmt.Println("returns element if exists:", LinearSearch2(elements, 111))
	fmt.Println("returns bool:", LinearSearch3(elements, 11))

}

// returns the index of the target elment
func LinearSearch(elements [9]int, item int) int {
	if len(elements) == 0 {
		return -1
	}

	for index, _ := range elements {
		if elements[index] == item {
			return index
		}
	}

	return -1

}

// in case want to return the value
// return intmax64 if target is not there, -1 can't be returned as it can be also one of the elment
func LinearSearch2(elements [9]int, item int) int {
	if len(elements) == 0 {
		return -1
	}

	for _, element := range elements {
		if element == item {
			return element
		}
	}

	return math.MaxInt64

}

// checks and returns bool if target exists or not
func LinearSearch3(elements [9]int, item int) bool {
	if len(elements) == 0 {
		return false
	}

	for _, element := range elements {
		if element == item {
			return true
		}
	}

	return false

}
