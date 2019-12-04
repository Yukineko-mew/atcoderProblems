package main

import (
	"fmt"
	"time"
)

func main() {
	var s string
	fmt.Scan(&s)

	layout := "2006/01/02"
	t, _ := time.Parse(layout, s)

	for !(t.Year()%int(t.Month()) == 0 && t.Year()%t.Day() == 0) {
		t = t.AddDate(0, 0, 1)
	}
	fmt.Println(t.Format(layout))
}
