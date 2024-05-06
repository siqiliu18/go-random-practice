package inttowords

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Assisted by WCA@IBM
// Latest GenAI contribution: ibm/granite-20b-code-instruct-v2
// TestNumberToWords_LessThan20 tests the function numberToWords() with input values less than 20.
func TestNumberToWords_LessThan20(t *testing.T) {
	assert.Equal(t, "One", numberToWords(1))
	assert.Equal(t, "Ten", numberToWords(10))
	assert.Equal(t, "Twenty", numberToWords(20))
	assert.Equal(t, "Twenty One", numberToWords(21))
	assert.Equal(t, "Thirty Nine", numberToWords(39))
	assert.Equal(t, "Nineteen", numberToWords(19))
}

// TestNumberToWords_Hundreds tests the function numberToWords() with input values between 100 and 999.
func TestNumberToWords_Hundreds(t *testing.T) {
	assert.Equal(t, "One Hundred", numberToWords(100))
	assert.Equal(t, "One Hundred Twenty", numberToWords(120))
	assert.Equal(t, "One Hundred Twenty Three", numberToWords(123))
	assert.Equal(t, "Two Hundred Thirty Four", numberToWords(234))
	assert.Equal(t, "Nine Hundred Ninety Nine", numberToWords(999))
}

// TestNumberToWords_Thousands tests the function numberToWords() with input values between 1,000 and 999,999.
func TestNumberToWords_Thousands(t *testing.T) {
	assert.Equal(t, "One Thousand", numberToWords(1000))
	assert.Equal(t, "One Thousand Two Hundred Thirty Four", numberToWords(1234))
	assert.Equal(t, "Two Thousand Three Hundred Forty Five", numberToWords(2345))
	assert.Equal(t, "Nine Thousand Nine Hundred Ninety Nine", numberToWords(9999))
}

// TestNumberToWords_Millions tests the function numberToWords() with input values between 1,000,000 and 999,999,999.
func TestNumberToWords_Millions(t *testing.T) {
	assert.Equal(t, "One Million", numberToWords(1000000))
	assert.Equal(t, "One Million Two Hundred Thirty Four Thousand", numberToWords(1234000))
	assert.Equal(t, "Two Million Three Hundred Forty Five Thousand Six Hundred Seventy Eight", numberToWords(2345678))
	assert.Equal(t, "Nine Million Nine Hundred Ninety Nine Thousand Nine Hundred Ninety Nine", numberToWords(999999999))
}

// TestNumberToWords_Zero tests the function numberToWords() with input value 0.
func TestNumberToWords_Zero(t *testing.T) {
	assert.Equal(t, "Zero", numberToWords(0))
}
