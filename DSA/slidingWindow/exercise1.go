package main

import "fmt"

/*
1. Maximum average subarray [643]

Given an integer array nums (values may be negative) and an integer k, find the contiguous subarray of length exactly k with the largest average. Return that average.
func findMaxAverage(nums []int, k int) float64
Example: nums = [5, 2, 8, 1, 9], k = 2 → windows [5,2]=3.5, [2,8]=5.0, [8,1]=4.5, [1,9]=5.0 → return 5.0.

Constraint: 1 <= k <= len(nums). Watch the integer-division trap when computing the average.
*/

func findMaxAverage(nums []int, k int) float64 {
	var maxAvg float64
	var sum float64
	for i := 0; i < k; i++ {
		sum += float64(nums[i])
	}

	maxAvg = float64(sum / float64(k))

	for i := k; i < len(nums); i++ {
		sum += float64(nums[i]) - float64(nums[i-k])
		avg := float64(sum / float64(k))

		if avg > maxAvg {
			maxAvg = avg
		}
	}

	return maxAvg
}

func main() {
	nums := []int{5, 2, 8, 1, 9}
	fmt.Printf("%f", findMaxAverage(nums, 2))
}
