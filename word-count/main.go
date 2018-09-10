package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	// スペース区切りで配列に入れる
	words := strings.Fields(sc.Text())

	uniqueWords := []string{}
	wordCount := map[string]int{}

	/*
		{
			"blue": 1,
			"red": 2,
			"green":3
		}
	*/

	for _, v := range words {

		// 未知の単語か？
		_, ok := wordCount[v]
		if !ok {
			// 未知の単語を単語リストに追加
			uniqueWords = append(uniqueWords, v)
		}
		// 単語の出現回数をカウント
		wordCount[v]++
	}

	// 最後の結果出力
	for _, w := range uniqueWords {
		fmt.Println(w, wordCount[w])
	}
}
