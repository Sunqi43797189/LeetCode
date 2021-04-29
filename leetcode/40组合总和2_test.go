package leetcode

import (
	"fmt"
	"testing"
)

func Test_combinationSum2(t *testing.T) {
	arr := []int{2, 5, 7, 1}
	result := combinationSum2(arr, 6)
	fmt.Println(result)
}
