package main

import "fmt"

//https://leetcode-cn.com/problems/count-of-matches-in-tournament/
func numberOfMatches(n int) int {
	res := []int{}
	dfs(n, &res)
	return SumArray(res)
}

func dfs(n int, res *[]int, ) {
	if n == 1 {
		return
	}
	fmt.Println(n)
	if n%2 == 0 { // 偶数队伍数
		n = n / 2 // 配对次数
		*res = append(*res, n)
		dfs(n, res)
	} else { // 奇数队伍数
		// 配对次数 (n-1) / 2
		*res = append(*res, (n-1)/2)
		n = (n-1)/2 + 1
		dfs(n, res)
	}
}

func SumArray(arr []int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}

func main() {
	n := 14
	res := []int{}
	dfs(n, &res)
	fmt.Println(SumArray(res))
}


