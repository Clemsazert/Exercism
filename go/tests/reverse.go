package main

import (
	"fmt"
)

func reverse(s []int) {
	for i := 0; i < int(len(s)/2); i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

func test_reverse() {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(s)
	reverse(s)
	fmt.Println(s)
}
