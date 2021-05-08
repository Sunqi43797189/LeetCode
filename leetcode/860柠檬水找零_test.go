package leetcode

import (
	"fmt"
	"testing"
)

func Test_lemonadeChange(t *testing.T) {
	arr := []int{5, 5, 5, 10, 20}
	result := lemonadeChange(arr)
	fmt.Println(result)
}
