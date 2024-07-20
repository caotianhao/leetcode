package main

import (
	"fmt"
	"math"
)

func minimumTotal2100(triangle [][]int) int {
	//相邻的结点在这里指的是下标与上一层结点下标相同或者等于上一层结点下标 +1 的两个结点
	//也就是说，如果正位于当前行的下标 i，那么下一步可以移动到下一行的下标 i 或 i+1
	l, res := len(triangle), math.MaxInt64
	dp := make([][]int, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]int, i+1)
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < l; i++ {
		// 这两行为边界条件，需要特殊处理
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
		for j := 1; j < i; j++ {
			dp[i][j] = min2100(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
		}
	}
	for i := 0; i < l; i++ {
		if dp[l-1][i] < res {
			res = dp[l-1][i]
		}
	}
	return res
}

func min2100(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(minimumTotal2100([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}))
}