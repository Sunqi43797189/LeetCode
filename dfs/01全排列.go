package main

import (
	"fmt"
	"strings"
)

func Dfs(arr []string, enable []bool, res []string) {
	if len(res) == len(arr) {
		fmt.Println(strings.Join(res, ""))
		return
	}
	for index, value := range arr {
		if enable[index] {
			enable[index] = false
			res = append(res, value)
			Dfs(arr, enable, res)
			res = res[:len(res)-1]
			enable[index] = true
		}
	}
}

func DfsV2(arr []string, res []string) {
	if len(res) == len(arr) {
		fmt.Println(strings.Join(res, ""))
		return
	}
	for index, value := range arr {
		if value != "" {
			res = append(res, value)
			arr[index] = ""
			DfsV2(arr, res)
			res = res[:len(res)-1]
			arr[index] = value
		}
	}
}

func main() {
	arr := []string{"A", "B", "C"}
	//enable := []bool{true, true, true}
	res := []string{}
	//Dfs(arr, enable, res)
	DfsV2(arr, res)
}
