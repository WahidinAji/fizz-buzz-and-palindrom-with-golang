package function

import (
	"fmt"
)

/*
Simple palindrome, reverse string manual
 */

func Palindrome(value string) bool {
	var temp = ""

	for i := len(value) -1; i >= 0; i-- {
		temp = temp + string(value[i])
	}
	fmt.Println(temp)
	result := temp == value
	return result
}

/*
Palindrome with no temporary variable.
Shared 2 `(/2)`, it means don't compare all characters when begin from first char to last char,
just comparing a half of all characters,
example `katak`
k : k
a : a
t : t enough and then return true, because a half after that, its always same
 */

func PalindromeComparePerChar(value string) bool {
	for i := 0; i < len(value) / 2; i++ {
		var firstIndex = i
		var lastIndex = len(value) - i -1
		//compare per index
		fmt.Println(string(value[firstIndex]) , " : " , string(value[lastIndex]))

		if string(value[firstIndex]) != string(value[lastIndex]) {
			fmt.Println("Palindrome false")
			return false
		}
	}
	fmt.Println("Palindrome true")
	return true
}

/*
Palindrome without looping.
The solution is use recursive.
 */

func isPalindromeRecursive(value string, i int) bool {
	if i < len(value) /2 {
		var firstIndex = i
		var lastIndex = len(value) - i -1
		//compare per index
		fmt.Println(string(value[firstIndex]) , " : " , string(value[lastIndex]))

		if string(value[firstIndex]) != string(value[lastIndex]) {
			fmt.Println("Palindrome recursive false")
			return false
		} else {
			//if true call again the recursive, and add 1 for i var
			return isPalindromeRecursive(value, i+1)
		}
	} else {
		fmt.Println("Palindrome recursive true")
		return true
	}
}

func PalindromeRecursive(value string) bool {
	return isPalindromeRecursive(value,0)
}
