package main

import "strings"

func main() {
	_ = validStrings(2)
}

func validStrings(n int) []string {
	result := []string{}
	generateString(n, "", &result)
	return result
}

func generateString(n int, str string, result *[]string) {
	if len(str) == n {
		*result = append(*result, str)
		return
	}

	generateString(n, str+"1", result)

	if strings.HasSuffix(str, "1") || len(str) < 1 {
		generateString(n, str+"0", result)
	}
}
