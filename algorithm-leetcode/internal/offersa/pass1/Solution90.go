package pass1

/**
 * 剑指 Offer II 089. 环形房屋偷盗最高金额
 * 相邻的房间都被偷盗，会触发报警
 */

	/**
     * 分情况讨论  
	 max(0~len-2, 0~len-1)都分开考虑了
     * dp[i]=max(dp[i-2]+nums[i],dp[i-1])
     * 偷第i-2间房和第i间房，偷第i-1间房
     * @param nums
     * @return
     */
	 func rob90(nums []int) int {
		//特殊情况
		if nums==nil{
			return 0
		}
		if len(nums)==1{
			return nums[0]
		}
		if len(nums)==2{
			return max2(nums[0],nums[1])
		}
	
		//nums[0]紧邻nums[len-1]
		return max2(robRange(nums,0,len(nums)-2),robRange(nums,1,len(nums)-1))
	}
	
	func robRange(nums []int, start int, end int) int {
		var pre int=nums[start]
		var cur int=max2(nums[start],nums[start+1])
	
		for i:=start+2;i<=end;i++{
			nxt:=max2(cur,pre+nums[i])
			pre=cur
			cur=nxt
		}
	
		return cur
	}
	
	func max2(a int,b int) int{
		if(a>=b){
			return a
		}else{
			return b
		}
	}