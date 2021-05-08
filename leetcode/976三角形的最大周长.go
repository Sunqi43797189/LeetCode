package leetcode

import "sort"

//https://leetcode-cn.com/problems/largest-perimeter-triangle/

func largestPerimeter(nums []int) int {
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 2; i-- {
		if nums[i-1]+nums[i-2] > nums[i] {
			return nums[i-1] + nums[i-2] + nums[i]
		}
	}
	return 0
}
