package helper

import (
	"math/rand"
	"sort"
	"strconv"
	"strings"
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

func SumArray(arr []int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}

func IsIntArrayContain(arr []int, target int) bool {
	for _, value := range arr{
		if value == target {
			return true
		}
	}
	return false
}

func IntArrayJoinString(arr []int, seq string) string  {
	var temp []string
	for _, value := range arr {
		temp = append(temp, strconv.Itoa(value))
	}
	return strings.Join(temp, seq)
}

func IsStringsArrayContain(array []string, item string) bool {
	for _, str := range array {
		if str == item {
			return true
		}
	}
	return false
}
