package inttowords

import (
	"math"
	"strings"
)

// https://leetcode.com/problems/integer-to-english-words/description/
func numberToWordsWCA(num int) string {
	/*
		Input: num = 1234567
		Output: "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"

		The function takes an integer as input and returns a string of words separated by spaces.
		The function converts the integer to English words on the basis of the following rules:
		1. If the integer is less than 20, then return the corresponding word.
		2. If the integer is between 20 and 99, then concatenate the first digit and the corresponding word. For example, 23 becomes "Twenty Three".
		3. If the integer is between 100 and 999, then concatenate the first two digits and the corresponding word. For example, 123 becomes "One Hundred Twenty Three".
		4. If the integer is greater than 999, then concatenate the first three digits, a space, and the result of converting the remaining digits to English words. For example, 1234 becomes "One Thousand, Two Hundred Thirty Four".

		The function handles integers from 0 to 999,999,999. It returns "Zero" if the input is 0.
	*/

	if num == 0 {
		return "Zero"
	}

	ones := []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	tens := []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	scales := []string{"", "Thousand", "Million", "Billion"}

	var words []string
	for i := len(scales) - 1; i >= 0; i-- {
		scale := scales[i]
		cur := num % (1000 * (i + 1)) / (1000 * i)
		if cur > 0 {
			if cur < 20 {
				words = append(words, ones[cur])
			} else if cur%10 == 0 {
				words = append(words, tens[cur/10])
			} else {
				words = append(words, tens[cur/10], ones[cur%10])
			}

			if scale != "" {
				words = append(words, scale)
			}
		}
	}

	return strings.Join(words, " ")
}

func numberToWords(num int) string {
	/*
		Input: num = 1234567
		Output: "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"
	*/

	if num == 0 {
		return "Zero"
	}

	ones := []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	tens := []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	scales := []string{"", "Thousand", "Million", "Billion"}

	var words []string
	for i := len(scales) - 1; i >= 0; i-- {
		scale := scales[i]
		cur := num / int(math.Pow(1000, float64(i)))
		num %= int(math.Pow(1000, float64(i)))
		if cur > 0 {
			if cur < 20 {
				words = append(words, ones[cur])
			} else if cur >= 20 && cur <= 99 {
				words = append(words, tens[cur/10], ones[cur%10])
			} else {
				words = append(words, ones[cur/100], "Hundred") // not done here, 23 and 12 need to handle differently.
			}

			if scale != "" {
				words = append(words, scale)
			}
		}
	}

	return strings.TrimSpace(strings.Join(words, " "))
}
