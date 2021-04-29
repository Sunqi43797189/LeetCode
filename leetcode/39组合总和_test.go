package leetcode

import (
	"fmt"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	arr := []int{2, 5, 7, 1}
	result := combinationSum(arr, 6)
	fmt.Println(result)
}
