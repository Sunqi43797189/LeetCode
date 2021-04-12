package helper

import (
	"math/rand"
	"sort"
	"time"
)

func RandomIntArray(length, max int) []int {
	array := make([]int, length)
	rand.Seed(time.Now().Unix())
	for i := 0; i < length; i++ {
		array[i] = rand.Intn(max)
	}
	return array
}

func CopyArray(array []int) []int {
	newArray := make([]int, len(array))
	for index, value := range array {
		newArray[index] = value
	}
	return newArray
}

func ArraySortCompare(dst []int, src []int) bool {
	sort.Ints(src)
	for index, value := range dst {
		if value != dst[index] {
			return false
		}
	}
	return true
}
