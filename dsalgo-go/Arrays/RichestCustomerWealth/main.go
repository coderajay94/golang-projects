package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/richest-customer-wealth/
func main() {
	fmt.Println("Richest Customer Wealth")
	accounts := [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}

	fmt.Println(maximumWealth(accounts))

}

func maximumWealth(accounts [][]int) int {
	maxWealth := math.MinInt
	for _, account := range accounts {
		totalWealth := 0
		for _, wealth := range account {
			totalWealth = totalWealth + wealth
		}
		if totalWealth > maxWealth {
			maxWealth = totalWealth
		}
	}
	return maxWealth
}
