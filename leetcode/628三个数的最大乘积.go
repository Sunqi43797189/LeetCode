package leetcode

import (
	"math"
	"sort"
)

/*
1. 都是正数 找三个最大的
2. 都是负数，找三个最大的负数
3. 有正有负，找连个小的负数的乘积，然后找一个最大的正数的
4. 取出来后比较三个数的大小
*/
func maximumProductSort(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) <= 3 {
		result := nums[0]
		for i := 1; i < len(nums); i++ {
			result *= nums[i]
		}
		return result
	}
	sort.Ints(nums)
	size := len(nums)
	a := nums[0] * nums[1] * nums[size-1]
	b := nums[size-1] * nums[size-2] * nums[size-3]
	if a > b {
		return a
	} else {
		return b
	}
}

func maximumProduct(nums []int) int {
	min1, min2 := math.MaxInt64, math.MaxInt64 // 第一小，第二小
	max1, max2, max3 := math.MinInt64, math.MinInt64, math.MinInt64 // 第一大 第二大 第三大
	for _, value := range nums {
		if value < min1 {
			min2 = min1
			min1 = value
		} else if value < min2 {
			min2 = value
		}
		if value > max1 {
			max3 = max2
			max2 = max1
			max1 = value
		} else if value > max2 {
			max3 = max2
			max2 = value
		} else if value > max3 {
			max3 = value
		}
	}
	a := min1 * min2 * max1
	b := max1 * max2 * max3
	if a > b {
		return a
	} else {
		return b
	}
}