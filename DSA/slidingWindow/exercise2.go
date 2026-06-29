package main

import "fmt"

/*
2. Count windows with average ≥ threshold [1343]

Given nums, k, and threshold, count how many length-k contiguous subarrays have an average >= threshold.
func numOfSubarrays(nums []int, k, threshold int) int
Example: nums = [2, 5, 5, 1, 3], k = 3, threshold = 4 → need sum >= 12. [2,5,5]=12 ✓, [5,5,1]=11 ✗, [5,1,3]=9 ✗ → return 1.

*,

*/

func numOfSubarrays(nums []int, k, threshold int) int {
	count := 0
	sum := 0

	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	if sum >= threshold*k {
		count++
	}

	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]

		if sum >= threshold*k {
			count++
		}
	}

	return count
}

func main() {
	nums := []int{2, 5, 5, 1, 3}

	fmt.Println(numOfSubarrays(nums, 3, 4))
}
