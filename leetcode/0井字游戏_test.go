package leetcode

import (
	"fmt"
	"testing"
)

func Test_tictactoe(t *testing.T) {
	board := []string{"OOX", "XXO", "OX "}
	result := tictactoe(board)
	fmt.Println(result)

}
