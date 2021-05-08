package leetcode

// https://leetcode-cn.com/problems/lemonade-change/
func lemonadeChange(bills []int) bool {
	if len(bills) == 0 {
		return true
	}

	five, ten := 0, 0
	for _, value := range bills {
		if value == 5 {
			five++
		} else if value == 10 {
			if five == 0 {
				return false
			} else {
				five--
				ten++
			}
		} else {
			if five >= 1 && ten >= 1 {
				five--
				ten--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}
