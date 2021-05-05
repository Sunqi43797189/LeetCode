package leetcode

import (
	"fmt"
	"testing"
)

func Test_pivotIndex(t *testing.T) {
	arr := []int{-1,-1,-1,-1,-1,0}
	result := pivotIndex(arr)
	fmt.Println(result)
}
