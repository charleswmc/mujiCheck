package utils

import "strings"

func SortASC(a []string) []string {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if strings.Compare(a[i], a[j]) > 0 {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	return a
}
