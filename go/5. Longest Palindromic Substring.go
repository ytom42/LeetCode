// 2/8
package main

import "fmt"

func longestPalindrome(s string) string {
	ans := ""
	l := len(s)

	for i := 0; i < l; i++ {
		start := i
		last := i
		ans_len := len(ans)
		for 0 <= start && last < l {
			if s[start] != s[last] {
				break
			}
			if ans_len < last-start+1 {
				ans = s[start : last+1]
			}
			start--
			last++
		}
		start = i
		last = i + 1
		for 0 <= start && last < l {
			if s[start] != s[last] {
				break
			}
			if ans_len < last-start+1 {
				ans = s[start : last+1]
			}
			start--
			last++
		}
	}

	return ans
}

type TestSet struct {
	S        string
	Expected string
}

func main() {
	inputs := []TestSet{
		{"babad", "bab"},
		{"cbbd", "bb"},
		{"ac", "a"},
	}

	for _, v := range inputs {
		ret := longestPalindrome(v.S)
		fmt.Println("Your Output:", ret, "Expected:", ret)
	}

	return
}
