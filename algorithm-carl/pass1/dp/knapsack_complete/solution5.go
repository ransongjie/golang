package knapsack_complete

import (
	"math"

	"com.xcrj/golang/algorithm-carl/common"
)

/**
 * https://leetcode.cn/problems/coin-change/
 * 给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。
 * 如果没有任何一种硬币组合能组成总金额，返回 -1。
 *
 * 背包最大承重 amount
 * 物品重量 coins
 * 每个物品可以选多次
 */
func coinChange(coins []int, amount int) int {
	//问题转换
	var bagSize int = amount
	var weight []int = coins
	// var value []int=coins
	var n int = len(weight)
	//dp数组
	var dp []int = make([]int, bagSize+1)
	for i := 0; i < len(dp); i++ { // 填充
		dp[i] = math.MaxInt
	}
	//初始状态
	dp[0] = 0 //dp[金额为0]=min硬币个数就是0个
	//状态转移
	//组合，最少硬币个数与顺序无关
	for i := 0; i < n; i++ {
		for j := weight[i]; j < bagSize+1; j++ {
			if dp[j-weight[i]] != math.MaxInt {
				dp[j] = common.Min[int](dp[j], 1+dp[j-weight[i]]) // 泛型函数
			}
		}
	}
	//结果
	if dp[bagSize] == math.MaxInt {
		return -1
	} else {
		return dp[bagSize]
	}
}
