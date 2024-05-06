package lc349decodestring

import (
	"strconv"
	"strings"
	"unicode"
)

type factor struct {
	dpi int
	n   int
	si  int
}

func decodeString(s string) string {
	index := 0
	stack := []factor{}
	for index < len(s) {
		if string(s[index]) == "[" {
			backward := index - 1
			for backward-1 >= 0 && unicode.IsDigit(rune(s[backward-1])) {
				backward--
			}
			dups, _ := strconv.Atoi(s[backward:index])
			stack = append(stack, factor{backward, dups, index})
		} else if string(s[index]) == "]" {
			end := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			str := ""
			// dups, _ := strconv.Atoi(string(s[end.dpi]))
			for i := 0; i < end.n; i++ {
				str += s[end.si+1 : index]
			}
			s = strings.Replace(s, s[end.dpi:index+1], str, 1)
			index = end.dpi + len(str) - 1
		}
		index++
	}
	return s
}
