package pass1

/**
 * 剑指 Offer II 104. 排列的数目，和等于target的排列的数目
 */
/**
 * 和等于target的排列的数目，和等于0的排列的数目，和等于1的排列额数目
 * dp[target]=和等于target的排列的数目，方案数
 * dp[0]=1
 * i>=nums[j]: dp[i]+=dp[i-nums[j]]
 * @param nums
 * @param target
 * @return
 */
func combinationSum4(nums []int, target int) int {
	//初始状态
	m:=target+1//从0~target
	dp:=make([]int,m)
	dp[0]=1;

	//状态转移
	for i:=1;i<m;i++{
		for j:=0;j<len(nums);j++{
			if i>=nums[j]{
				dp[i]+=dp[i-nums[j]]
			}
		}
	}

	//
	return dp[m-1];
}