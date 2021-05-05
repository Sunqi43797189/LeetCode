package leetcode

import (
	"fmt"
	"testing"
)

func Test_largestPerimeter(t *testing.T) {
	arr := []int{3, 6, 2, 3}
	result := largestPerimeter(arr)
	fmt.Println(result)
}
