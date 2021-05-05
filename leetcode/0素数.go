package leetcode

// 暴力算法
func PrimeCount(n int) int {
	if n <= 1 {
		return 0
	}
	count := 0
	for x := 2; x < n; x++ {
		if isPrime(x) {
			count += 1
		}
	}
	return count
}

// 埃氏筛选法
func Eratosthenes(n int) int {
	if n <= 1 {
		return 0
	}
	arr := make([]bool, n) // 初始默认所有都为素数
	count := 0
	for i := 2; i < n; i++ {
		if arr[i] == false {
			count++
			for j := i * i; j < n; j += i {
				arr[j] = true
			}
		}
	}
	return count
}

func isPrime(x int) bool {
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
