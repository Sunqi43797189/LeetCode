package leetcode

import (
	"fmt"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	arr := []int{0,1,2,2,3,3,4}
	result := removeDuplicates(arr)
	fmt.Println(result)
}
