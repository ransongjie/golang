package stock_sell

import . "com.xcrj/golang/algorithm-carl/common"

/**
 * 至多k次
 * https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv/
 * 买卖股票至多k次
 */
func maxProfitd(k int, prices []int) int {
	//dp数组
	n := len(prices)
	//dp[第i天][第j次][0持有现金]
	dp := Slice3D[int](n, k+1, 2) // !!! k+1，第0次什么也不做
	//初始状态
	for i := 1; i < k+1; i++ {
		//dp[第0天][第i次][1持有股票]
		dp[0][i][1] = -prices[0]
	}
	//状态转移
	for i := 1; i < n; i++ {
		for j := 1; j <= k; j++ {
			// 卖出先买入; dp[i][第j次][0持有现金]=, dp[i-1][第j次][1持有股票]
			dp[i][j][0] = Max[int](dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			// 买入先卖出；dp[i][第j次][1持有股票]=, dp[i-1][第j-1次][0持有现金] 第j-1次卖，第j次买
			dp[i][j][1] = Max[int](dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
	}
	//结果
	return dp[n-1][k][0] //0持有现金
}
