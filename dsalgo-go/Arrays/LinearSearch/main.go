package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Hello World from Linear Search!!")

	elements := [9]int{12, 14, 1, 34, 61, 88, 56, -11, 11}
	elementsSlice := []int{14, 1, 34, 61, 88, 56, -11, 11}

	fmt.Println("Search returns index:", LinearSearch(elements, 11))
	fmt.Println("Search returns element if exists:", LinearSearch2(elements, 111))
	fmt.Println("Search returns bool:", LinearSearch3(elements, 11))
	index, found := LinearSearch4(elementsSlice, 11, 3, 7)
	fmt.Println("Searches between range returns found:", found, "at ", index)

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

func LinearSearch4(elements []int, target int, start int, end int) (int, bool) {

	if len(elements) == 0 || end > len(elements)-1 || start < 0 || start > end {
		return -1, false
	}

	//end+1 as end is not inclusive start to end-1 would be returned otherwise
	elements = elements[start : end+1]
	fmt.Println("updated slice", elements)

	for index, element := range elements {
		if element == target {
			//return index of the range start array
			return start + index, true
		}
	}
	return -1, false
}
