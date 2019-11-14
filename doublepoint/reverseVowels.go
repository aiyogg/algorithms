package main

/**
反转字符串中的元音字符
Given s = "leetcode", return "leotcede".
*/
type chars []byte

func (vowels chars) contains(c byte) bool {
	for _, char := range vowels {
		if c == char {
			return true
		}
	}
	return false
}

var vowels = chars{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}

func reverseVowels(s string) string {
	j := len(s)
	result := make([]byte, j)
	for i := 0; i < j; {
		ci := s[i]
		cj := s[j]
		if !vowels.contains(ci) {
			_ = append(result, ci)
			i++
		} else if !vowels.contains(cj) {
			//_ = append(re)
		}
	}
}
