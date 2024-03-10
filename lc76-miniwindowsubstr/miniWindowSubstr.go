package lc76miniwindowsubstr

import (
	"math"
)

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

func minWindow2(s string, t string) string {
	res := s
	lens, lent := len(s), len(t)
	count := 0
	li, ri := 0, 1

	if lens < lent {
		return ""
	}

	tm := make(map[byte]int)
	for _, c := range t {
		tm[byte(c)] = 0
	}

	for ri < lens {
		licc, ok := tm[s[li]]
		if ok { // found char
			if licc == 0 { // not yet count
				tm[s[li]] = 1
				count += 1
				// ri = li + 1
			} // else: found before - is it possible?
		} else { // not found
			li += 1
			// ri = li + 1
		}
		ricc, ok := tm[s[ri]]
		if ok { // found right char
			if ricc == 0 { // not yet count
				tm[s[ri]] = 1
				count += 1
				if count == lent { // found all t char
					if (ri - li) < len(res) {
						res = s[li : ri+1]
					}
					tm[s[li]] = 0
					count -= 1
					li += 1
				} else { // not all t char found
					ri += 1
				}
			}
		} else { // not found right char
			ri += 1
		}
	}

	return res
}

// https://leetcode.com/problems/minimum-window-substring/
func minWindowAns(s string, t string) string {
	res := ""
	tm := make(map[byte]int)
	li, count := 0, 0
	minLen := math.MaxInt

	for _, c := range t {
		tm[byte(c)] += 1
	}

	for i := 0; i < len(s); i++ {
		tm[s[i]] -= 1
		if tm[s[i]] >= 0 {
			count += 1
		}
		for count == len(t) {
			if minLen > i-li {
				minLen = i - li
				res = s[li : i+1]
			}
			tm[s[li]] += 1
			if tm[s[li]] > 0 {
				count -= 1
			}
			li++
		}
	}

	return res
}
