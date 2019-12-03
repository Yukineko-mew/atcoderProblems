package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)

	for 2 <= len(s) {
		n := len(s) - 2
		s = s[:n]

		mid := n / 2
		// fmt.Println(n, s, mid, s[:mid], s[mid:])
		if s[:mid] == s[mid:] {
			fmt.Println(n)
			break
		}
	}
}
