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
	for i, j := 0, len(c_string)-1; i < j; i, j = i+1, j-1 {
		c_string[i], c_string[j] = c_string[j], c_string[i]
	}
	return strings.Join(c_string, "")
}

//1.3

func removeDuplicateChars(s string) string {
	var c_string []string = strings.Split(s, "")
	var (
		result string = ""
	)
	result = result + c_string[0]
	for i := 0; i < len(c_string)-1; i++ {
		for j := i + 1; j < len(c_string); j++ {
			if !strings.Contains(result, c_string[j]) {
				result = result + c_string[j]
			}
		}
	}
	return result
}

//1.4

func isAnagram(s string) bool {
	var (
		reverse_s string   = ""
		c_string  []string = strings.Split(s, "")
		flag      bool     = false
	)
	for i := len(c_string) - 1; i >= 0; i-- {
		reverse_s = reverse_s + c_string[i]
	}
	if reverse_s == s {
		flag = true
	}
	return flag
}

//1.5

func replaceSpaces(s string) string {
	var (
		replaceStr = ""
	)
	for _, r := range s {
		if string(r) == " " {
			replaceStr = replaceStr + "%20"
		} else {
			replaceStr = replaceStr + string(r)
		}
	}
	return replaceStr
}

//1.7
type Matrix [][]int

func zeroReplaceMatrix(mat Matrix) Matrix {
	var (
		zero_row int = 0
		zero_col int = 0
	)
	for row, i := range mat {
		for col, elem := range i {
			if elem == 0 {
				zero_row = row
				zero_col = col
			}
		}
		for r, w := range mat {
			for s, _ := range w {
				if r == zero_row {
					mat[r][s] = 0
				}
				if s == zero_col {
					mat[r][s] = 0
				}
			}
		}
	}
	return mat
}
