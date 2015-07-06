//1.1

package arrays_and_strings

import "strings"

func uniqueString(s string) bool {
	var (
		flag bool = true
	)
	for _, r := range s {
		if strings.Count(s, string(r)) > 1 {
		  flag = false
		}
	}
	return flag
}

//1.2

func reverseString(s string) string {
  //in Go, strings are immutable
	var c_string []string = strings.Split(s, "")
	for i, j :=0, len(c_string)-1; i<j; i, j = i+1, j-1 {
    c_string[i], c_string[j] = c_string[j], c_string[i] 
	}
	return strings.Join(c_string,"")
}

//1.3

func removeDuplicateChars(s string) string {
  
}
