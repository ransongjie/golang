package pass1
/**
 * 剑指 Offer II 103. 最少的硬币数目
 * 输入，coins[]硬币面值，硬币面值至少为1，amount 总金额，硬币数量无限
 * 输出，组成总金额的最少硬币个数; 没有任何一种硬币组合能组成总金额，返回 -1。
 */
    /**
     * 组成amount所需最少硬币数量，组成子amount的最少硬币数量
     * dp[i]=组成i值的最小硬币数量
     * dp[0]=0
     * i>=coins[j]: dp[i]=min((dp[i-coints[j]])+1,dp[i]), 把每种币值都尝试下
     */
	 func coinChange(coins []int, amount int) int {
		//初始状态
		m:=amount+1// 从0~amount
		dp:=make([]int,m)
		for i:=1;i<m;i++{// 求组成i的最小值
			dp[i]=amount+1;
		}
		dp[0]=0//amount=0 选取0个硬币
	
		//状态转移
		for i:=1;i<m;i++{
			for j:=0;j<len(coins);j++{//把每种币值都尝试下
				if i>=coins[j]{//当前amount i >= 当前币值
					dp[i]=min103(dp[i-coins[j]]+1,dp[i])
				}
			}
		}
	
		//
		if dp[m-1]==amount+1{
			return -1
		}else{
			return dp[m-1]
		}
	}
	
	func min103(a,b int)int{
		if(a<=b){
			return a
		}else{
			return b
		}
	}