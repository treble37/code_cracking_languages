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


