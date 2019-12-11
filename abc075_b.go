package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanLines)
	line := strings.Split(nextLine(), " ")
	h, _ := strconv.Atoi(line[0])
	w, _ := strconv.Atoi(line[1])

	//  i-1,j-1 i-1,j i-1,j+1
	//  i  ,j-1   i,j i  ,j+1
	//  i+1,j-1 i+1,j i+1,j+1
	roundI := []int{-1, -1, -1, 0, 0, 0, 1, 1, 1}
	roundJ := []int{-1, 0, 1, -1, 0, 1, -1, 0, 1}

	arr := make([][]string, h)
	for i := 0; i < h; i++ {
		arr[i] = strings.Split(nextLine(), "")
	}
	ansArr := make([]string, h)
	for i := 0; i < h; i++ {
		var ansStr bytes.Buffer
		for j := 0; j < w; j++ {
			// 自身が#ならスキップ
			if arr[i][j] == "#" {
				ansStr.WriteString("#")
				continue
			} else {
				cnt := 0
				// fmt.Printf("%d %d\n", i, j)
				for n := 0; n < 9; n++ {
					//
					moveI := i + roundI[n]
					moveJ := j + roundJ[n]
					// 配列の要素をはみ出るようならスキップ
					if moveI < 0 || moveJ < 0 || h <= moveI || w <= moveJ {
						// fmt.Printf("x")
						continue
					}
					if arr[moveI][moveJ] == "#" {
						cnt++
						// fmt.Printf("1")
					} else {
						// fmt.Printf("0")
					}
				}
				// fmt.Printf("\n")
				ansStr.WriteString(strconv.Itoa(cnt))
			}
		}
		ansArr = append(ansArr, ansStr.String())
	}

	for i := 0; i < len(ansArr); i++ {
		fmt.Println(ansArr[i])
	}
}
