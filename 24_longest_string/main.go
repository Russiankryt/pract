package main

import (
	"fmt"
)

func longestString(strings ...string) string {
	longest := ""
	for _, s := range strings {
		if len(s) > len(longest) {
			longest = s
		}
	}
	return longest
}

func main() {
	word := longestString("дом", "восьмидесятичетырёхлетний", "код", "технология")
	fmt.Println("Самая длинная строка:", word)
}
