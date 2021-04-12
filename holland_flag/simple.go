package main

import "fmt"

// 小于num的数放右边，大于的放左边
func main() {
	arr := []int{19, 20, 10, 9, 10, 8, 2, 3, 7, 10, 16, 16, 18}
	num := 10
	j, i := -1, 0
	for ; i < len(arr)-1; i++ {
		if arr[i] <= num {
	 		arr[j+1], arr[i] = arr[i], arr[j+1]
	 		j++
		}
	}
	fmt.Println(arr)
}
