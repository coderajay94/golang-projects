package main

import "fmt"

func main() {
	nums := []int{3, 5, 7, 9, 10, 90, 100, 130, 140, 160, 170, 189, 200, 213}
	target := 100

	fmt.Println("answer is ", FindInInfiniteArray(nums, target))

}

//Infinite basically means we don't know the start and end
// we need to start doubling the size and see if target lies between that range

func FindInInfiniteArray(nums []int, target int) int {

	//{3, 5,  start+end = 1
	// 7, 9, 10, 90, start+end = 10 - (start+end-1)
	// 130, 140, 160, 170, 180, 190, 200, 230 }
	start := 0
	end := 1

	//size := 2

	for target > nums[end] {
		// start should be end+1
		temp := end + 1 // this is my new start, keeping temp so that original value should be intact
		end = start + (end-start+1)*2
		start = temp
		fmt.Println("running for loop", end)
	}

	return BinarySearchInInfiniteArray(nums, target, start, end)
}

func BinarySearchInInfiniteArray(nums []int, target int, start int, end int) int {

	for start <= end {
		mid := (start + end) / 2

		if target > nums[mid] {
			start = mid + 1
		} else if target < nums[mid] {
			end = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
