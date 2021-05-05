package leetcode

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	i, j := 0, 1
	for j < len(nums) {
		if nums[j] != nums[i] { // 不相等，先修改i的下一个元素，然后i++，j++
			nums[i+1] = nums[j]
			j++
			i++
		} else { // 相等j继续走
			j++
		}
	}
	return i + 1
}
