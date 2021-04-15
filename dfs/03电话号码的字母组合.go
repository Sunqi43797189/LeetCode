package main

import (
	"fmt"
	"strings"
)

var data = [][]string{
	{},
	{},
	{"a", "b", "c"},
	{"d", "e", "f"},
	{"g", "h", "i"},
	{"j", "k", "l"},
	{"m", "n", "o"},
	{"p", "q", "r", "s"},
	{"t", "u", "v"},
	{"w", "x", "y", "z"},
}

var res = []string{}

func DfsPhoneNum(str string, index int, res []string, tempArr []string) {
	if index == len(str) {
		res = append(res, strings.Join(tempArr, ""))
		fmt.Println(res)
		return
	}
	for _, value := range data[index] {
		fmt.Println(value)
		tempArr = append(tempArr, value)
		DfsPhoneNum(str, index+1, res, tempArr)
		tempArr = tempArr[:len(tempArr)-1]
	}

}

func main() {
	str := "2,3"
	temp := []string{}
	DfsPhoneNum(str, 0, res, temp)
	fmt.Println(res)
}
