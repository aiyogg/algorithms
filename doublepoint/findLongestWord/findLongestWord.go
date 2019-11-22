package findlongestword

import "strings"

/**
Longest Word in Dictionary through Deleting
Input:
s = "abpcplea", d = ["ale","apple","monkey","plea"]

Output:
"apple"
*/

func findLongestWord(s string, d []string) string {
	longestWord := ""
	for _, target := range d {
		l1, l2 := len(longestWord), len(target)
		if l1 > l2 || (l1 == l2 && strings.Compare(longestWord, target) < 0) {
			continue
		}
		if isSubstr(s, target) {
			longestWord = target
		}
	}
	return longestWord
}

func isSubstr(s, target string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(target) {
		if s[i] == target[j] {
			j++
		}
		i++
	}
	return j == len(target)
}
