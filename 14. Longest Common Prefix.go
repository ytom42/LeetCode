// 1/26
package main

func longestCommonPrefix(strs []string) string {
	ans := strs[0]

	for i := 1; i < len(strs); i++ {
		tmp := strs[i]
		minlen := min(len(ans), len(tmp))
		if minlen < len(ans) {
			ans = ans[:minlen]
		}
		for j := 0; j < minlen; j++ {
			if ans[j] != tmp[j] {
				ans = ans[:j]
				break
			}
		}
		if ans == "" {
			break
		}
	}
	return ans
}

func main() {
	input := [][]string{
		{"flower", "flow", "flight"},
		{"dog", "racecar", "car"},
		{"ab", "a"},
	}

	for _, v := range input {
		ret := longestCommonPrefix(v)
		println(ret)
	}

	return
}
