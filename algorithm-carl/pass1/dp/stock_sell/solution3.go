package stock_sell

import "com.xcrj/golang/algorithm-carl/common"

/**
 * 至多2次
 * https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/
 * 买卖股票至多2次
 */
func maxProfitc(prices []int) int {
	//dp数组
	n := len(prices)
	//dp[i][0] 无操作
	//dp[i][1] 第1次买入
	//dp[i][2] 第1次卖出
	//dp[i][3] 第2次买入
	//dp[i][4] 第2次卖出
	dp := common.Slice2D[int](n, 5)
	//初始状态
	dp[0][0] = 0 //!!!
	dp[0][1] = -prices[0]
	dp[0][2] = 0
	dp[0][3] = -prices[0]
	dp[0][4] = 0
	//状态转移
	for i := 1; i < n; i++ {
		//dp[i-1][0]=第i-1天什么也没干
		dp[i][1] = common.Max[int](dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = common.Max[int](dp[i-1][2], dp[i-1][1]+prices[i])
		dp[i][3] = common.Max[int](dp[i-1][3], dp[i-1][2]-prices[i])
		dp[i][4] = common.Max[int](dp[i-1][4], dp[i-1][3]+prices[i])
	}
	//结果
	// 第2次卖出包含了第1次卖出，所以第2次卖出（持有现金）一定是最大值
	return dp[n-1][4]
}
