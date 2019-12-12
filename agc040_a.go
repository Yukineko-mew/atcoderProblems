package main

import (
	"fmt"
)

//  < > > > < < > < < < < < > > > <

//  <       < <   < < < < <       <
//   1       1 2   1 2 3 4 5       1

//    > > >     >           > > >
//   3 2 1     1           3 2 1

// 0<3>2>1>0<1<2>0<1<2<3<4<5>2>1>0<1

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var s string
	fmt.Scan(&s)

	resArr := make([]int, len(s)+1)

	// まず純粋に<の数を数えて素直に値を入れていく
	for i, v := range s {
		if v == '<' {
			resArr[i+1] = resArr[i] + 1
		}
	}

	// for _, v := range resArr {
	// 	fmt.Printf(" %d", v)
	// }
	// fmt.Println("")

	// 今度は逆から>の数を数えて値を入れていく
	// すでに入っている値より大きいなら上書きする
	for i := range s {
		j := len(s) - i - 1
		if s[j] == '>' {
			resArr[j] = max(resArr[j], resArr[j+1]+1)
		}
	}

	// for _, v := range resArr {
	// 	fmt.Printf(" %d", v)
	// }
	// fmt.Println("")

	var sum int64 = 0
	for _, v := range resArr {
		sum += int64(v)
	}
	fmt.Println(sum)
}
