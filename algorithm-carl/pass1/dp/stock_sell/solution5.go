package stock_sell

import . "com.xcrj/golang/common"

/**
 * https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown/
 * 只有一支股票
 * 股票只买卖多次
 * 有冷冻期
 * [买入，卖出，冷冻，买入。卖出]
 * 两次买入卖出之间有冷冻期
 */
func maxProfite(prices []int) int {
	//dp数组
	n := len(prices)
	dp := Slice2D[int](n, 4)
	//初始状态
	//0持有股票
	dp[0][0] = -prices[0]
	//1持有现金（无操作）通过无操作持有现金 ！！！只前一定没有卖出股票的操作
	dp[0][1] = 0
	//2持有现金（卖出股票）通过卖出股票持有现金 ！！！只前一定有卖出股票的操作
	dp[0][2] = 0
	//3冷冻期
	dp[0][3] = 0
	//状态转移
	for i := 1; i < n; i++ {
		//左边是要达到的状态
		dp[i][0] = Max[int](dp[i-1][0], Max[int](dp[i-1][1]-prices[i], dp[i-1][3]-prices[i]))
		dp[i][1] = Max[int](dp[i-1][1], dp[i-1][3]) //dp[i-1][3]通过无操作持有现金
		dp[i][2] = dp[i-1][0] + prices[i]           //只前一定有卖出股票的操作
		dp[i][3] = dp[i-1][2]
	}
	//结果
	return Max[int](dp[n-1][1], Max[int](dp[n-1][2], dp[n-1][3]))
}
