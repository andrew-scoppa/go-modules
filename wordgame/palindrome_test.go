package wordgame

import (
	"testing"
)

/*
This test function tests the IsPalindrome function in the wordgame package.
details:
	- The function should return true for the words "kayak" and "bob" as they are palindromes.
	- The function should return false for the words "apple", "orange" and "pear" as they are not palindromes.
assertions:
	- The test should error if the function returns false for the words "kayak" and "bob".
*/
func TestPalindome(t *testing.T) {
	var words = []string{"kayak", "bob"}
	for _, word := range words {
		if !IsPalindrome(word) {
			t.Errorf("'%s' incorrectly not identifed as palindrome", word)
		}
	}

}


/*
This function tests for non-palindromes.
details:
	- The function should return false for the words "apple", "orange" and "pear" as they are not palindromes.
assertions:
	- The test should error if the function returns true for the words "apple", "orange" and "pear".
*/
func TestNonPalindome(t *testing.T) {
	var words = []string{"apple", "orange", "pear"}
	for _, word := range words {
		if IsPalindrome(word) {
			t.Errorf("'%s' incorrectly identifed as palindrome", word)
		}
	}
}
