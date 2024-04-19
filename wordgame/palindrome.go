package wordgame


/*
This function takes a string and returns a boolean value indicating whether the string is a palindrome.
A palindrome is a word that is written the same forwards and backwards.
Example:
	"racecar" is a palindrome
	"hello" is not a palindrome
*/
func IsPalindrome(word string) bool{

	for i := range word{
		l := len(word)
		if word[i] != word[l - i - 1]{
			return false
		}
	}
	return true
}