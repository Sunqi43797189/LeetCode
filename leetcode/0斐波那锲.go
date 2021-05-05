package leetcode

func FibonacciSequence(num int) int {
	hash := make(map[int]int)
	return fibio(num, hash)
}

func fibio(num int, hash map[int]int) int {
	if num <= 1 {
		return num
	}
	if value, ok := hash[num]; ok {
		return value
	}
	hash[num] = fibio(num-1, hash) + fibio(num-2, hash)
	return hash[num]
}

func fibioNormal(num int) int {
	if num <= 1 {
		return num
	}
	return fibioNormal(num-1) + fibioNormal(num-2)
}

func fibioDoublePointer(num int) int {
	if num <= 1 {
		return num
	}
	low , high := 0, 1
	for i := 2; i <=num; i ++{
		sum := low + high
		low = high
		high = sum
	}
	return high
}
