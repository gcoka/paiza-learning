package main

import "fmt"
import "sort"

func main() {
	// 何個数字を扱うか？
	var n int
	fmt.Scan(&n)
	// 配列を用意する
	var s []int

	// forで配列に数字を入れる
	for i := 0; i < n; i++ {
		var v int
		fmt.Scan(&v)
		s = append(s, v)
	}

	// ソートする
	sort.Sort(sort.IntSlice(s))

	for i := 0; i < n; i++ {
		fmt.Println(s[i])
	}
}
