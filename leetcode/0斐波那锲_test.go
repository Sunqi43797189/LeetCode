package leetcode

import (
	"fmt"
	"testing"
	"time"
)

func TestFibonacciSequence(t *testing.T) {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()
	//result := FibonacciSequence(100)
	//result := fibioNormal(100)
	result := fibioDoublePointer(100)
	fmt.Println(result)
}
