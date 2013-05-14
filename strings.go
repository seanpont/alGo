package strings

import (
	"fmt"
)

//returns the index of pat in the text txt, or -1 if not found
func BruteForceSearch(pat, txt string) int {
	fmt.Printf("Searching for %s in \"%s\"", pat, txt)
	M := len(pat)
	N := len(txt)
	for i := 0; i <= N-M; i++ {
		j := 0
		for j = 0; j < M; j++ {
			if txt[i+j] != pat[j] {
				break
			}
		}
		if j == M {
			return i
		}
	}
	return -1
}

//same as BruteForceSearch, just different pointers
func BruteForceSearch2(pat, txt string) int {
	M, N, i, j := len(pat), len(txt), 0, 0
	for ; i < N && j < M; i++ {
		if txt[i] == pat[j] {
			j++
		} else {
			i, j = i-j, 0
		}
	}
	if j == M {
		return i - M
	}
	return -1
}
