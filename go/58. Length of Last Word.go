// 2/2
package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	words := strings.Fields(s)

	if len(words) == 0 {
		return 0
	}

	return len(words[len(words)-1])
}

type TestSet struct {
	S        string
	Expected int
}

func main() {
	inputs := []TestSet{
		{"Hello World", 5},
		{"   fly me   to   the moon  ", 4},
		{"luffy is still joyboy", 6},
		{" ", 0},
	}

	for _, v := range inputs {
		ret := lengthOfLastWord(v.S)
		fmt.Println("Your Output:", ret, "Expected:", v.Expected)
	}

	return
}
