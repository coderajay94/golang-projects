package main

import "fmt"

func main() {

	fmt.Println("Running Binary Search...")

	elements := []int{12, 40, 56, 122, 134, 245, 4567, 5678}
	target := 245

	fmt.Println("Finding ", target, " using binary search on index:", BinarySearch(elements, target))
}

func BinarySearch(elements []int, target int) int {

	start := 0
	end := len(elements) - 1
	for start <= end {
		fmt.Println("checking start and end is", start, end)
		mid := (start + end) / 2
		fmt.Println("mid is", mid)
		if elements[mid] == target {
			return mid
		} else if elements[mid] > target {
			end = mid - 1
		} else if elements[mid] < target {
			start = mid + 1
		}
	}

	return -1
}

func BinarySearchMax(elements []int, target int) int {
	max := 0
	start := 0
	end := len(elements) - 1
	for start <= end {
		mid := (start + end) / 2

		if elements[mid] > target {
			if elements[mid] < max {
				max = target
			}
			end = mid - 1
		} else if elements[mid] < target {
			start = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
