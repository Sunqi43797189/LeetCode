package main

func DfsAllNum(arr []int, res *[][]int, level int, boolArr []bool, tempArr []int) {
	if level >= len(arr)+1 {
		a := make([]int, len(tempArr))
		copy(a, tempArr)
		*res = append(*res, a)
		return
	}
	for i := 0; i < len(arr); i++ {
		if boolArr[i] {
			tempArr = append(tempArr, arr[i])
			boolArr[i] = false
			DfsAllNum(arr, res, level+1, boolArr, tempArr)
			boolArr[i] = true
			tempArr = tempArr[:len(tempArr)-1]
		}
	}

}

// func main() {
// 	arr := []int{1, 2, 3}
// 	var res [][]int
// 	boolArr := make([]bool, len(arr))
// 	for i := 0; i < len(arr); i ++ {
// 		boolArr[i] = true
// 	}
// 	DfsAllNum(arr, &res, 1, boolArr, []int{})
// 	fmt.Println(res)
// }
