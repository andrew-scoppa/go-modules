/*
This file contains the main function that will be executed when the program is run.
*/

package main

/*
import the fmt package to print to the console
import the wordgame package in this module to use the IsPalindrome function
*/
import (
    "fmt"
	"andrew-scoppa/words/wordgame"
)

/*
Write the main function that will be executed when the program is run.

Details:
	- Prompt the user to enter a word.
	- Read the word from the console.
	- If the input is 'q', exit the program.
	- Check if the word is a palindrome.
	- Print the result to the console.
	- Repeat the process.

*/
func main() {
	for {
		fmt.Println("Enter a word:")
		var word string
		fmt.Scanln(&word)
		if word == "q" {
			break
		}
		if wordgame.IsPalindrome(word) {
			fmt.Println("The word is a palindrome.")
		} else {
			fmt.Println("The word is not a palindrome.")
		}
	}
}


