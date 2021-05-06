package leetcode

import "fmt"

// https://leetcode-cn.com/problems/house-robber/

func rob(nums []int) int {
	//return robRecursion(nums, 0)
	return robIterate(nums)
}

// 递归版本 最优解 第一个选不选 取决于 第二个与第一个加第三个和的谁大
func robRecursion(nums []int, index int) int {
	n := len(nums)
	if n == 0 || index > n-1 {
		return 0
	}
	if index == n-1 {
		return nums[index]
	}
	a := robRecursion(nums, index+1)
	b := robRecursion(nums, index+2) + nums[index]
	if a > b {
		return a
	} else {
		return b
	}
}

func robIterate(nums []int) int {
	size := len(nums)
	if nums == nil || size == 0 {
		return 0
	}
	if size == 1 {
		return nums[0]
	}
	dpArr := make([]int, size)
	dpArr[0] = nums[0]
	dpArr[1] = maxInt(nums[0], nums[1])
	for i := 2; i < size; i++ {
		dpArr[i] = maxInt(dpArr[i-2]+nums[i], dpArr[i-1])
	}
	fmt.Println(dpArr)
	return dpArr[size-1]
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
