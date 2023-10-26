package knapsack_complete

/**
 * https://leetcode.cn/problems/combination-sum-iv/submissions/
 * 给定一个由正整数组成且不存在重复数字的数组，找出和为给定目标正整数的组合的个数。
 * nums = [1, 2, 3]
 * target = 4
 * 所有可能的组合为： (1, 1, 1, 1) (1, 1, 2) (1, 2, 1) (1, 3) (2, 1, 1) (2, 2) (3, 1)
 * 请注意，顺序不同的序列被视作不同的组合。
 * 因此输出为 7
 * <p>
 * 实质是排列数
 * 每个元素可以使用任意多次
 */
func combinationSum4(nums []int, target int) int {
	//问题转换
	var bagSize int = target
	var weight []int = nums
	// var value []int = nums
	var n int = len(weight)
	//dp数组
	var dp []int = make([]int, bagSize+1)
	//初始状态
	dp[0] = 1 //背包承重为0,1个物品也不放也是一种方案
	//状态转移
	//排列，先背包再物品
	//若先物品再背包，物品1一定在物品2之前，固定顺序，只是一种排列
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
