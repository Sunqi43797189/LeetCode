package main

import (
	"fmt"
	"strings"
)

var data = map[string][]string{
	"0": {},
	"1": {},
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

var res = []string{}

func DfsPhoneNum(str string, index int, res *[]string, tempArr []string) {
	array := strings.Split(str, "")
	if index == len(str) {
		*res = append(*res, strings.Join(tempArr, ""))
		return
	}
	for _, value := range data[array[index]] {
		tempArr = append(tempArr, value)
		DfsPhoneNum(str, index+1, res, tempArr)
		tempArr = tempArr[:len(tempArr)-1]
	}

}

func main() {
	str := "239"
	temp := []string{}
	DfsPhoneNum(str, 0, &res, temp)
	fmt.Println(res)
}
