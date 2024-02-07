// 2/7
package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	ans := 0
	l := len(s)

	for i := 0; i < l; i++ {
		tmp := 0
		for j := i + 1; j < l; j++ {
			if strings.Contains(s[i:j], string(s[j])) {
				break
			}
			tmp++
		}
		ans = max(ans, tmp+1)
	}

	return ans
}

type TestSet struct {
	s        string
	expected int
}

func main() {
	inputs := []TestSet{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{" ", 1},
		{"a", 1},
	}

	for _, v := range inputs {
		ret := lengthOfLongestSubstring(v.s)
		fmt.Println("Expected:", v.expected, "Your Output:", ret)
	}

	return
}
