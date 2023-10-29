package stock_sell

import "com.xcrj/golang/algorithm-carl/common"

/**
 * 买卖多次，不限次
 * https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/
 * 只有一支股票
 * 股票只买卖多次
 */
func maxProfitb(prices []int) int {
	//dp数组
	n := len(prices)
	//dp[第i天][0]=持有股票的最大收益 dp[第i天][1]=持有现金的最大收益
	dp := common.Slice2D[int](n, 2)
	//初始状态
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	//状态转移
	for i := 1; i < n; i++ {
		//第i-1天持有股票，第i-1天持有现金-今天买入股票
		dp[i][0] = common.Max[int](dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = common.Max[int](dp[i-1][1], dp[i-1][0]+prices[i])
	}
	//结果
	return dp[n-1][1]
}
