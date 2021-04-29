package leetcode

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

func isPrime(x int) bool {
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
