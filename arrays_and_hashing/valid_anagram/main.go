package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	fmt.Println(isAnagram("rat", "car"))
}

// Given two strings s and t, return true if t is an anagram of s, and false otherwise.

// Example 1:

// Input: s = "anagram", t = "nagaram"

// Output: true

// Example 2:

// Input: s = "rat", t = "car"

// Output: false

// Constraints:

// 1 <= s.length, t.length <= 5 * 104
// s and t consist of lowercase English letters.

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	arr1 := strings.Split(sortString(s), "")
	arr2 := strings.Split(sortString(t), "")

	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func sortString(s string) string {
	arr := strings.Split(s, "")
	slices.Sort(arr)
	return strings.Join(arr, "")
}
