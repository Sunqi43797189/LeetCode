package main

import (
	"fmt"
	"strconv"
	"strings"
)

func DfsIP(str string, index int, level int, res []string) {
	if level == 5 || index == len(str) {
		if level == 5 && index == len(str) {
			fmt.Println(strings.Join(res, "."))
		}
		return
	}
	for i := 1; i < 4; i++ {
		var subStr string
		if index+1+i >= len(str) {
			subStr = str[index+1:]
		} else {
			subStr = str[index+1 : index+1+i]
		}
		num, _ := strconv.Atoi(subStr)
		if num <= 255 && (subStr == "0" || !strings.HasPrefix(subStr, "0")) {
			res = append(res, subStr)
			DfsIP(str, index+i, level+1, res)
			res = res[:len(res)-1]
		}
	}
}

// func main() {
// 	str := "19216801"
// 	res := []string{}
// 	DfsIP(str, -1, 1, res)
// }
