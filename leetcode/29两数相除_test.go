package leetcode

import (
	"fmt"
	"testing"
)

func Test_divide(t *testing.T) {
	//result := divide(10, 3)
	//fmt.Println(result)
	for i := 0; i <= 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
}
