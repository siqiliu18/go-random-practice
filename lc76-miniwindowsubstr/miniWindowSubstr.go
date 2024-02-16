package lc76miniwindowsubstr

import "math"

func minWindow(s string, t string) string {
	res := ""
	letterCnt := make(map[byte]int)

	left, cnt, minLen := 0, 0, math.MaxInt

	for i := 0; i < len(t); i++ {
		letterCnt[t[i]] += 1
	}

	for i := 0; i < len(s); i++ {
		letterCnt[s[i]] -= 1
		if letterCnt[s[i]] >= 0 {
			cnt += 1
		}

		for cnt == len(t) {
			if minLen > i-left+1 {
				res = s[left : i+1]
			}
			letterCnt[s[left]] += 1
			if letterCnt[s[left]] > 0 {
				cnt -= 1
			}
			left += 1
		}
	}
	return res
}
