package main

import (
	"fmt"
	"testing"
)

func perm(s []rune) [][]rune {
	if len(s) == 2 {
		return [][]rune{{s[0], s[1]}, {s[1], s[0]}}
	}
	var result [][]rune
	for idx, r := range s {
		var sub_runes []rune
		sub_runes = append(sub_runes, s[:idx]...)
		sub_runes = append(sub_runes, s[idx+1:]...)
		sub_result := perm(sub_runes)
		for i, sub := range sub_result {
			sub_result[i] = append([]rune{r}, sub...)
		}
		result = append(result, sub_result...)
	}
	return result
}

func fact(n int) int {
	result := 1
	for i := n; i > 1; i-- {
		result *= i
	}
	return result
}

func TestPerm(t *testing.T) {
	s := "abc"
	result := perm([]rune(s))
	for _, r := range result {
		fmt.Printf("%c\n", r)
	}
}
