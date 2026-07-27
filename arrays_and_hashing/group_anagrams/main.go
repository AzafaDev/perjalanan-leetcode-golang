package main

import (
	"slices"
	"strings"
)

// Given an array of strings strs, group the anagrams together. You can return the answer in any order.

// Example 1:

// Input: strs = ["eat","tea","tan","ate","nat","bat"]

// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

// Explanation:

// There is no string in strs that can be rearranged to form "bat".
// The strings "nat" and "tan" are anagrams as they can be rearranged to form each other.
// The strings "ate", "eat", and "tea" are anagrams as they can be rearranged to form each other.
// Example 2:

// Input: strs = [""]

// Output: [[""]]

// Example 3:

// Input: strs = ["a"]

// Output: [["a"]]

// Constraints:

// 1 <= strs.length <= 104
// 0 <= strs[i].length <= 100
// strs[i] consists of lowercase English letters.

func groupAnagrams(strs []string) [][]string {
	group := make(map[string][]string)
	for i, str := range strs {
		sorted := sortString(str)
		if _, ok := group[sorted]; ok {
			group[sorted] = append(group[sorted], strs[i])
			continue
		}
		group[sorted] = []string{strs[i]}
	}
	result := make([][]string, 0)
	for _, value := range group {
		result = append(result, value)
	}
	return result

}

func sortString(str string) string {
	stringSlice := strings.Split(str, "")
	slices.Sort(stringSlice)
	return strings.Join(stringSlice, "")
}
