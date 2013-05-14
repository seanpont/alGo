package strings

import (
	"fmt"
	"testing"
)

func TestSqrt(t *testing.T) {
	fmt.Println("testing brute force")
	index := BruteForceSearch("cat", "find the cat in this hat.")
	fmt.Printf("\nsubstring found at index %v\n", index)
	index = BruteForceSearch2("bat", "the ball was hit by the bat")
	fmt.Printf("found %s at %v", "bat", index)
	if index != 23 {
		t.Errorf("should have found 'bat' at 24, found at %v", index)
	}
}
