package arrays_and_strings

import (
	"testing"
)

func TestUniqueStringCharacters(t *testing.T) {
  got := uniqueString("abcdxyz")
	if got == false {
	  t.Fatalf("uniqueString failed")
	}
}

func TestDuplicateStringCharacters(t *testing.T) {
  got := uniqueString("aabcdxyz")
	if got == true {
	  t.Fatalf("uniqueString failed")
	}
}

func TestReverseString(t *testing.T) {
  got := reverseString("abcde")
	if got != "edcba" {
		t.Fatalf("reverseString fails")
	}
}

func TestRemoveDuplicateChars(t *testing.T) {
  got := removeDuplicateChars("abbcde")
	if got != "abcde" {
		t.Fatalf("removeDuplicateChars fails")
	}
	got = removeDuplicateChars("abcabcde")
	if got != "abcde" {
		t.Fatalf("removeDuplicateChars fails")
	}
}

func TestIsAnagram(t *testing.T) {
  got := isAnagram("abba")
	if got != true {
		t.Fatalf("isAnagram fails")
	}
}
