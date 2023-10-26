package pass1
/**
 * 剑指 Offer II 102. 加减的目标值
 * 输入，一个正整数数组 nums 和一个整数 target 
 * 向数组中的每个整数前添加 '+' 或 '-' 构造一个 表达式
 * 输出，正数和+负数和= target 的不同表达式的数目
 */
 //+nums[i] 或 -nums[i]
func findTargetSumWays(nums []int, target int) int {
	var cnt int=0

	var dfs func(int,int)
	dfs =func(sum int,idx int){
		//回退
		if idx==len(nums){
			if sum==target{
				cnt++
			}
			return
		}

		//回溯
		dfs(sum+nums[idx],idx+1)
		dfs(sum-nums[idx],idx+1)
	}
	dfs(0,0)

	return cnt
}