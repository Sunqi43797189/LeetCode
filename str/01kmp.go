package main

import (
	"fmt"
)

func GetNextArray(str string) []int {
	if len(str) == 0 {
		return []int{}
	}
	if len(str) == 1 {
		return []int{-1}
	}

	if len(str) == 2 {
		return []int{-1, 0}
	}
	next := make([]int, len(str))
	next[0] = -1
	next[1] = 0
	for i := 2; i < len(str); i++ {
		maxLen := -1
		for j := i - 1; j > 0; j-- {
			if str[0:i-j] == str[j:i] && maxLen < (i-j) { // 前缀等于后缀 取
				maxLen = i - j
			}
		}
		next[i] = maxLen
	}
	return next
}

func KMP(str, subStr string) int {
	if len(str) == 0 {
		return 0
	}
	if len(subStr) > len(str) {
		return -1
	}

	if len(subStr) == 0 {
		return 0
	}
	i, j := 0, 0
	nextArray := GetNextArray(subStr)
	for i < len(str) && j < len(subStr) {
		if str[i] == subStr[j] { // 相等同时加一
			i++
			j++
		} else if nextArray[j] == -1 { // 负一了，子串nextarray第一个元素，此时主串加一
			i++
		} else { // 辅助串跳转到前缀的下一个位置
			j = nextArray[j]
		}
	}
	if j == len(subStr) {
		return i - j
	} else {
		return -1
	}
}

func main() {
	subStr := "abcdabe"
	//fmt.Println(GetNextArray(subStr))

	str := "abcedabcdabeaploabcda"
	fmt.Println(KMP(str, subStr))
}
