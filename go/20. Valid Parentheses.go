// 1/26
package main

func isValid(s string) bool {
	var stack []rune

	m := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, v := range s {
		l := len(stack)
		switch v {
		case '(', '[', '{':
			stack = append(stack, v)
		case ')', ']', '}':
			if l > 0 && (stack[l-1] == m[v]) {
				stack = stack[:l-1]
			} else {
				return false
			}
		}
	}
	if len(stack) > 0 {
		return false
	} else {
		return true
	}
}

func main() {

	inputs := []string{
		"()",
		"()[]{}",
		"(]",
	}

	for _, v := range inputs {
		ret := isValid(v)
		println(ret, ":", v)
	}

	return
}
