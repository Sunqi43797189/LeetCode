package helper

import (
	"math/rand"
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
