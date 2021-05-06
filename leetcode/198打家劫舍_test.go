package leetcode

import (
	"fmt"
	"testing"
)

func Test_rob(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	result := rob(nums)
	fmt.Println(result)
}
