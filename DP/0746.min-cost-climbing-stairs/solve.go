func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < len(cost); i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}
	return min(dp[len(dp)-1], dp[len(dp)-2])
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}