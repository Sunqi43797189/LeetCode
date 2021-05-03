package leetcode

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	for index, value := range nums {
		if pos, ok := hash[target - value];ok {
			return []int{index, pos}
		}
		hash[value] = index
	}
	return []int{}
}

func twoSumDfs(nums []int, target int) []int {
	result := []int{}
	boolArr := make([]bool, len(nums))
	dfsTwoSum(&result, nums, target, []int{}, boolArr)
	return result
}

func dfsTwoSum(resArr *[]int, nums []int, target int, tempArr []int, boolArr []bool) {
	if len(*resArr) == 2 {
		return
	}
	if len(tempArr) == 2 {
		if nums[tempArr[0]]+nums[tempArr[1]] == target && len(*resArr) < 2 {
			*resArr = append(*resArr, tempArr...)
		}
		return
	}
	for i := 0; i < len(nums); i++ {
		if boolArr[i] == false {
			boolArr[i] = true
			tempArr = append(tempArr, i)
			dfsTwoSum(resArr, nums, target, tempArr, boolArr)
			boolArr[i] = false
			tempArr = tempArr[:len(tempArr)-1]
		}
	}
}
