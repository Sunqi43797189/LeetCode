package leetcode

//https://leetcode-cn.com/problems/champagne-tower/

func champagneTower(poured int, query_row int, query_glass int) float64 {
	ca := make([][]float64, 100)
	for i := 0; i < 100; i++ {
		ca[i] = make([]float64, 100)
	}
	ca[0][0] = float64(poured)
	for i := 0; i < query_row; i++ {
		for j := 0; j <= i; j++ {
			diff := (ca[i][j] - 1.0) / 2
			if diff > 0 {
				ca[i+1][j] += diff
				ca[i+1][j+1] += diff
			}
		}
	}
	value := ca[query_row][query_glass]
	if value > 1 {
		return 1
	}
	return value

}
