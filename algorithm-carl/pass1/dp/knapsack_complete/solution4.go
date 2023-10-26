package knapsack_complete

/**
 * https://leetcode.cn/problems/climbing-stairs/
 * 爬楼梯。需要 m 阶你才能到达楼顶
 * 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
 *
 * n是背包总承重
 * 1,2 是物品重量，每个物品可以选择多次
 */
func climbStairs(m int) int {
	//问题转换
	var bagSize int = m
	var weight [2]int = [2]int{1, 2}
	// var value [2]int = [2]int{1, 2}
	var n int = len(weight)
	//dp数组
	var dp []int = make([]int, bagSize+1)
	//初始状态
	dp[0] = 1
	//状态转移
	//排列，先走1步再走2步和先走2步再走1步是不同的排列
	for i := 0; i < bagSize+1; i++ {
		for j := 0; j < n; j++ {
			if i >= weight[j] {
				dp[i] += dp[i-weight[j]]
			}
		}
	}
	//结果
	return dp[bagSize]
}
