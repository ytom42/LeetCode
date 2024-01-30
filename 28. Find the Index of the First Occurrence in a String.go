// 1/30
package main

import "fmt"

func strStr(haystack string, needle string) int {

	tmp_i := 0
	tmp_j := 0
	lenNeedle := len(needle)
	lenHaystack := len(haystack)

	for i := 0; i < lenHaystack; i++ {
		if haystack[i] == needle[0] {
			tmp_i = i
			tmp_j = 0
			for tmp_i < lenHaystack && tmp_j < lenNeedle {
				if haystack[tmp_i] != needle[tmp_j] {
					break
				}
				tmp_i++
				tmp_j++
				if tmp_j == lenNeedle {
					return i
				}
			}
		}
	}

	return -1
}

type Input struct {
	Haystack string
	Needle   string
}

func main() {
	inputs := []Input{
		Input{"sadbutsad", "sad"},
		Input{"leetcode", "leeto"},
		Input{"a", "a"},
		Input{"a", "b"},
		Input{"leetcode", "leetcodee"},
		Input{"leetcode", "eet"},
		Input{"mississippi", "issip"},
	}

	for _, v := range inputs {
		ret := strStr(v.Haystack, v.Needle)
		fmt.Println("ret:", ret)
	}

	return
}
