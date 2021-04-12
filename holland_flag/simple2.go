package main

import (
	"fmt"
)

// 小于num的数放右边，等于的放中间，大于的放左边
func main() {
	arr := []int{10, 8, 5, 3, 1, 10, 11, 9, 4, 19, 17, 20}
	num := 10
	j, k, i := 0, len(arr)-1, 0
	for ; i < len(arr)-1 && i <= k; {
		if arr[i] < num {
			arr[j], arr[i] = arr[i], arr[j]
			j++
			i++
		} else if arr[i] == num {
			i++
		} else {
			arr[k], arr[i] = arr[i], arr[k]
			k--
		}
	}
	fmt.Println(arr)
}
