package main

func DfsNonReapeated(size int, res *[][]int, pb []int, arr []int, tempArr []int) {
	if len(tempArr) == size {
		a := make([]int, len(tempArr))
		copy(a, tempArr)
		*res = append(*res, a)
		return
	}
	for i := 0; i < len(arr); i++ {
		if pb[i] > 0 {
			pb[i]--
			tempArr = append(tempArr, arr[i])
			DfsNonReapeated(size, res, pb, arr, tempArr)
			tempArr = tempArr[:len(tempArr)-1]
			pb[i]++
		}
	}
}

// func main() {
// 	arr := []int{1, 1, 2}
// 	var res [][]int
// 	var countMap = make(map[int]int)

// 	for _, value := range arr {
// 		if _, ok := countMap[value]; ok {
// 			countMap[value] ++
// 		} else {
// 			countMap[value] = 1
// 		}
// 	}
// 	var nonRepeatedArr = make([]int, len(countMap))
// 	var pb = make([]int, len(countMap))
// 	index := 0
// 	for value, count := range countMap {
// 		nonRepeatedArr[index] = value
// 		pb[index] = count
// 		index++
// 	}
// 	DfsNonReapeated(len(arr), &res, pb, nonRepeatedArr, []int{})
// 	fmt.Println(res)
// }
