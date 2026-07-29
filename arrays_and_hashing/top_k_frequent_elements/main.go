package main

import "slices"

// Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

// Example 1:

// Input: nums = [1,1,1,2,2,3], k = 2

// Output: [1,2]

// Example 2:

// Input: nums = [1], k = 1

// Output: [1]

// Example 3:

// Input: nums = [1,2,1,2,1,2,3,1,3,2], k = 2

// Output: [1,2]

// Constraints:

// 1 <= nums.length <= 105
// -104 <= nums[i] <= 104
// k is in the range [1, the number of unique elements in the array].
// It is guaranteed that the answer is unique.

// Follow up: Your algorithm's time complexity must be better than O(n log n), where n is the array's size.

func main() {}

func topKFrequent(nums []int, k int) []int {
	mapping := make(map[int]int)
	sliceInt := make([]int, 0)
	for _, num := range nums {
		mapping[num]++
	}
	for key := range mapping {
		sliceInt = append(sliceInt, key)
	}
	slices.SortFunc(sliceInt, func(i, j int) int {
		return mapping[j] - mapping[i]
	})

	return sliceInt[:k]
}
