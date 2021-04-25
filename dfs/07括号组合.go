package main

import (
	"fmt"
)

func dfs(size int, res *[]string, str string, tempMap map[string]int) {
	if len(str) == size {
		*res = append(*res, str)
		return
	}
	if tempMap["("] > 0 {
		tempMap["("] --
		dfs(size, res, str + "(", tempMap)
		tempMap["("] ++
	}
	if tempMap[")"] > 0 && tempMap["("] != tempMap[")"] {
		tempMap[")"] --
		dfs(size, res, str + ")", tempMap)
		tempMap[")"] ++
	}
}

func main() {
	n := 3
	var res []string
	tempMap := map[string]int{
		"(": n,
		")": n,
	}
	dfs(n*2, &res, "", tempMap)
	fmt.Println(res)
}
