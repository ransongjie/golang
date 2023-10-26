package knapsack_complete

import "math"

/**
 * https://leetcode.cn/problems/perfect-squares/
 * 给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。
 * 给你一个整数 n ，返回和为 n 的完全平方数的 最少数量。
 *
 * 背包：n
 * 物品重量：i*i
 * 顺序：求组合
 * 求最小方法数
 */
func numSquares(n int) int {
	//问题转换
	var bagSize int = n
	var weight []int = []int{}
	for i := 1; i*i <= n; i++ { //题目要求从1开始
		weight = append(weight, i*i)
	}
	// var value []int=weight
	var n1 = len(weight)
	//dp数组
	var dp []int = make([]int, bagSize+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = math.MaxInt
	}
	//初始状态
	dp[0] = 0
	//状态转换
	//组合，求最少数量与顺序无关
	for i := 0; i < n1; i++ {
		for j := weight[i]; j < bagSize+1; j++ {
			// 不需要if条件，n一定可以由n个1组成
			dp[j] = Min[int](dp[j], 1+dp[j-weight[i]])
		}
	}
	//结果
	return dp[bagSize]
}
