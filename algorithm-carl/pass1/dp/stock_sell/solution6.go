package stock_sell

import . "com.xcrj/golang/common"

/**
 * 卖出有手续费
 * https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
 * 只有一支股票
 * 股票只买卖多次
 * 买入到卖出有1次手续费
 */
func maxProfitf(prices []int, fee int) int {
	//dp数组
	//dp[第i天][0/1]=持有股票/现金的最大收益
	n := len(prices)
	dp := Slice2D[int](n, 2)
	//初始状态
	dp[0][0] = -prices[0]
	dp[0][1] = 0 //可以不买入股票直接持有现金
	//状态转移
	for i := 1; i < n; i++ {
		dp[i][0] = Max[int](dp[i-1][0], dp[i-1][1]-prices[i])
		//卖出时有手续费
		dp[i][1] = Max[int](dp[i-1][1], dp[i-1][0]+prices[i]-fee)
	}
	//结果
	return Max(dp[n-1][1], dp[n-1][0]) // -手续费可能导致卖亏损
}
