package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	sc.Split(bufio.ScanWords)

	n, k := nextInt(), nextInt()
	l := []int{}

	for i := 0; i < n; i++ {
		l = append(l, nextInt())
	}

	sort.Sort(sort.Reverse(sort.IntSlice(l)))

	var r float64 = 0
	for i := k - 1; 0 <= i; i-- {
		r = (r + float64(l[i])) / 2.0
	}

	fmt.Printf("%.10f\n", r)
}
