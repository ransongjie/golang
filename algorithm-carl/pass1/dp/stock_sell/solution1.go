package stock_sell

import (
	"math"

	"com.xcrj/golang/algorithm-carl/common"
)

/**
 * 买卖1次
 * https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/
 * 给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
 * 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
 * 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
 * <p>
 * 只有1只股票 and 股票只买卖1次
 */
func maxProfit(prices []int) int {
	//dp数组
	//dp[第i天][0/1]=持有股票/现金的最大收益
	n := len(prices)
	dp := common.Slice2D[int](n, 2)
	//初始状态
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	//状态转移
	for i := 1; i < n; i++ {
		//第i天持有股票的最大收益=前1天持有股票的最大收益，第i天持有股票(只能买入1次)
		dp[i][0] = common.Max[int](dp[i-1][0], -prices[i])
		//！！！第i-1天持有股票是一种状态，不是说第i-1天才买入股票，之前的某1天可能就买入了股票
		dp[i][1] = common.Max[int](dp[i-1][1], dp[i-1][0]+prices[i])
	}
	//结果
	return dp[n-1][1] //只能买卖1只股票1次，最大收益自然是最后1天持有现金
}

// 贪心
func maxProfit1(prices []int) int {
	low := math.MaxInt
	profix := 0

	for _, price := range prices {
		low = common.Min[int](low, price)
		//在low基础上求最大利润
		profix = common.Max[int](profix, price-low)
	}

	return profix
}
