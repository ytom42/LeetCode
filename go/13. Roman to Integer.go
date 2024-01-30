// 1/25
package main

func romanToInt(s string) int {

	r := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	pre := s[0]
	ans := 0
	skip := false

	for i := 1; i < len(s); i++ {
		curr := s[i]
		if pre == 'I' && (curr == 'V' || curr == 'X') {
			ans = ans + r[curr] - r[pre]
			skip = true
		} else if pre == 'X' && (curr == 'L' || curr == 'C') {
			ans = ans + r[curr] - r[pre]
			skip = true
		} else if pre == 'C' && (curr == 'D' || curr == 'M') {
			ans = ans + r[curr] - r[pre]
			skip = true
		} else if skip {
			skip = false
		} else {
			ans = ans + r[pre]
		}
		pre = curr
	}
	if !skip {
		ans = ans + r[pre]
	}
	return ans
}

func main() {
	inputs := []string{
		"III",
		"LVIII",
		"MCMXCIV",
	}

	for _, v := range inputs {
		ret := romanToInt(v)
		println(ret, ":", v)
	}

	return
}
