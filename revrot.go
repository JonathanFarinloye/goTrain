package main

import (
	"fmt"
	"strconv"
)

func revRot(str string, sz int) string {
	totalStr := ""
	if sz <= 0 || str == "" {
		return totalStr
	}
	for i := 0; i < len(str)/sz; i++ {
		newStr := ""
		strInt := 0
		for j := i * sz; j < sz*(i+1); j++ {
			newStr += string(str[j])
			x, err := strconv.Atoi(string(str[j]))
			if err == nil {
				strInt += (x * x * x)
			}
		}
		if strInt%2 == 0 {
			totalStr += reverseString(newStr)
			fmt.Println(newStr)
		} else {
			stringSlice := []rune(newStr)
			one := stringSlice[0]
			for w := 0; w < len(stringSlice); w++ {
				if w+1 == len(stringSlice) {
					stringSlice[w] = one
				} else {
					stringSlice[w] = stringSlice[w+1]
				}
			}
			totalStr += string(stringSlice)
		}
	}
	return totalStr
}

func reverseString(s string) string {
	stringSlice := []rune(s)
	for i, j := 0, len(stringSlice)-1; i < j; i, j = i+1, j-1 {
		stringSlice[i], stringSlice[j] = stringSlice[j], stringSlice[i]
	}

	return string(stringSlice)
}