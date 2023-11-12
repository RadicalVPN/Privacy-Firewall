package util

import "strings"

func FilterDuplicates(input []string) []string {
	seen := make(map[string]struct{}, len(input))
	j := 0
	for _, v := range input {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		input[j] = v
		j++
	}
	return input[:j]
}

func FilterComments(input string) string {
	if strings.HasPrefix(input, "#") {
		return ""
	}

	if strings.HasPrefix(input, "!") {
		return ""
	}

	return input
}
