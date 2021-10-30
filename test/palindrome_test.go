package test

import (
	"github.com/WahidinAji/fizz-buzz-with-golang/function"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPalindrome(t *testing.T)  {
	assert.Equal(t, true, function.Palindrome("a"))
	assert.Equal(t, true, function.Palindrome("aba"))
	assert.Equal(t, true, function.Palindrome("kodok"))
	assert.Equal(t, true, function.Palindrome(""))
	assert.Equal(t, false, function.Palindrome("ab"))
	assert.Equal(t, false, function.Palindrome("abab"))
	assert.Equal(t, false, function.Palindrome("kodcok"))
	assert.Equal(t, false, function.Palindrome("aji"))
}

func TestPalindromeComparePerChar(t *testing.T)  {
	assert.Equal(t, true, function.PalindromeComparePerChar("a"))
	assert.Equal(t, true, function.PalindromeComparePerChar("aba"))
	assert.Equal(t, true, function.PalindromeComparePerChar("kodok"))
	assert.Equal(t, true, function.PalindromeComparePerChar(""))
	assert.Equal(t, false, function.PalindromeComparePerChar("ab"))
	assert.Equal(t, false, function.PalindromeComparePerChar("abab"))
	assert.Equal(t, false, function.PalindromeComparePerChar("kodcok"))
	assert.Equal(t, false, function.PalindromeComparePerChar("aji"))
}

func TestPalindromeRecursive(t *testing.T)  {
	assert.Equal(t, true, function.PalindromeRecursive("a"))
	assert.Equal(t, true, function.PalindromeRecursive("aba"))
	assert.Equal(t, true, function.PalindromeRecursive("kodok"))
	assert.Equal(t, true, function.PalindromeRecursive(""))
	assert.Equal(t, false, function.PalindromeRecursive("ab"))
	assert.Equal(t, false, function.PalindromeRecursive("abab"))
	assert.Equal(t, false, function.PalindromeRecursive("kodcok"))
	assert.Equal(t, false, function.PalindromeRecursive("aji"))
}

func TestPalindromeSimple(t *testing.T)  {
	function.Palindrome("kodok")
}
func TestPalindromeCompare(t *testing.T)  {
	function.PalindromeComparePerChar("katak")
}
func TestPalindromeRecursiveSimple(t *testing.T)  {
	function.PalindromeRecursive("katak")
}
