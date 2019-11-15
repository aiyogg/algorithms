package main

import "fmt"

/**
反转字符串中的元音字符
Given s = "leetcode", return "leotcede".
*/
type vowel []byte

func (v vowel) contains(c byte) bool {
	for _, char := range v {
		if c == char {
			return true
		}
	}
	return false
}

var vowels = vowel{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}

func reverseVowels(s string) string {
	j := len(s) - 1
	result := make([]byte, int(len(s)))
	for i := 0; i <= j; {
		ci := s[i]
		cj := s[j]
		if !vowels.contains(ci) {
			result[i] = ci
			i++
		} else if !vowels.contains(cj) {
			result[j] = cj
			j--
		} else {
			result[i] = cj
			result[j] = ci
			i++
			j--
		}
	}
	return string(result)
}

func main() {
	fmt.Printf("reverseVowels: %v", reverseVowels("Leetcode"))
}
