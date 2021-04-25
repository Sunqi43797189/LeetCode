package main

import "fmt"

func Dp(x int) int {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}

	return Dp(x-1) + Dp(x-2)
}

func DpV2(x int) int {
	arr := make([]int, x+1)
	arr[1] = 1
	arr[2] = 1
	for i := 3; i < x+1; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[x]
}

func main() {
	fmt.Println(Dp(7))
	fmt.Println(DpV2(7))
}
