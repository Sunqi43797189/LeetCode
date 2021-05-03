package leetcode

// 暴力算法
func mySqrtV1(x int) int {
	for i := 0; i < x; i++ {
		if i*i == x {
			return i
		}
		if i*i < x && (i+1)*(i+1) > x {
			return i
		}
		if i*i < x && (i+1)*(i+1) == x {
			return i + 1
		}
	}
	return 0
}

// 二分查找法
func mySqrt(x int) int {
	left, right, result := 0, x, -1 // left right default
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid <= x {
			result = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}
