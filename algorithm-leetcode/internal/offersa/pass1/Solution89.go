package pass1
/**
 * 剑指 Offer II 089. 房屋偷盗最高金额
 * 相邻的房间都被偷盗，会触发报警
 */
	/**
     * dp[i]=偷到第i间房能够获得最大收益
	 * dp[i]=偷盗0~i间房能获取的最大收益
     * dp[i]=Math.max(dp[i-2]+nums[i],dp[i-1]);
     * 偷到了第i-2间房和第i间房
     * 偷到了第i-1间房
     * @param nums
     * @return
     */
func rob(nums []int) int {
	//特殊情况
	if nums==nil{
		return 0
	}
	if len(nums)==1{
		return nums[0]
	}
	if len(nums)==2{
		return max(nums[0],nums[1])
	}

	//初始化
	var dp []int=make([]int,len(nums))
	dp[0]=nums[0]
	dp[1]=max(nums[0],nums[1])

	//动态规划
	for i:=2;i<len(nums);i++{
		dp[i]=max(dp[i-2]+nums[i],dp[i-1])
	}

	return dp[len(nums)-1]
}

func max(a int,b int) int{
	if(a>=b){
		return a
	}else{
		return b
	}
}

/*
 滚动数组
 */
func rob2(nums []int) int {
	//特殊情况
	if nums==nil{
		return 0
	}
	if len(nums)==1{
		return nums[0]
	}
	if len(nums)==2{
		return max(nums[0],nums[1])
	}

	//初始化
	var pre int=nums[0]
	var cur int=max(nums[0],nums[1])

	//动态规划
	for i:=2;i<len(nums);i++ {
		var nxt int=max(cur,pre+nums[i])
		pre=cur
		cur=nxt
	}

	return cur
}