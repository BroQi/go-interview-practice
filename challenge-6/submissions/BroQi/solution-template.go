// Package challenge6 contains the solution for Challenge 6.
package challenge6

import (
	// Add any necessary imports here
	"fmt"
	"regexp"
	"strings"
)

// CountWordFrequency takes a string containing multiple words and returns
// a map where each key is a word and the value is the number of times that
// word appears in the string. The comparison is case-insensitive.
//
// Words are defined as sequences of letters and digits.
// All words are converted to lowercase before counting.
// All punctuation, spaces, and other non-alphanumeric characters are ignored.
//
// For example:
// Input: "The quick brown fox jumps over the lazy dog."
// Output: map[string]int{"the": 2, "quick": 1, "brown": 1, "fox": 1, "jumps": 1, "over": 1, "lazy": 1, "dog": 1}
func CountWordFrequency(text string) map[string]int {
	// 结果map
	result := make(map[string]int)
	if len(text) == 0 || len(strings.TrimSpace(text)) == 0 {
		return result
	}
	re := regexp.MustCompile(`[,\s\r\n\t-]+`)
	replaceRe := regexp.MustCompile(`[^a-z0-9,\s\r\n\t-]+`)
	text = replaceRe.ReplaceAllString(strings.ToLower(strings.TrimSpace(text)), "")
	strArr := re.Split(text, -1)
	fmt.Println(strArr)
	
	for _, word := range strArr {
		count := result[word]
		if count > 0 {
			result[word] = result[word] + 1
		} else {
			result[word] = 1
		}
	}
	return result
} 