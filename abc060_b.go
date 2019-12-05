package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)

	for i := 1; i <= 100; i++ {
		x := a * i
		if x%b == c {
			fmt.Println("YES")
			return
		}
	}
	fmt.Println("NO")
}
