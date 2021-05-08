package leetcode

import (
	"fmt"
	"testing"
)

func Test_removeElement(t *testing.T) {
	arr := []int{0, 1, 2, 2, 3, 0, 4, 2}
	result := removeElement(arr, 2)
	fmt.Println(result)
}
