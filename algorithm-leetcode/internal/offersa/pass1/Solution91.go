package pass1

/**
 * 剑指 Offer II 091. 粉刷房子
 * cost[i][j] j in [0,2] 三种不同颜色
 * 相邻房子的颜色不能一样
 */
	/**
     * 分情况讨论
	   最后一间房被刷成红色的最小成本，被刷成黄色的成本, 被刷成蓝色的成本
     * dpnew[j]=costs[i][j]+Math.min(dp[(j+1)%3],dp[(j+2)%3]);
     * 第i间房刷成0色的成本+min(第i-1间房刷成1色的最小总成本+第i-1间房刷成2色的最小总成本)
     */
	 func minCost(costs [][]int) int {
		//初始化
		dp:=make([]int,3)
		for i:=0;i<3;i++{//第0间房被刷成各个颜色的成本
			dp[i]=costs[0][i];
		}
	
		//
		
		for i:=1;i<len(costs);i++{
			dpnew:=make([]int,3)
			for j:=0;j<3;j++{
				dpnew[j]=costs[i][j]+min(dp[(j+1)%3],dp[(j+2)%3])
			}
			dp=dpnew
		}
	
		//
		return min(min(dp[0],dp[1]),dp[2])
	}
	
	func min(a int,b int)int{
		if a<=b{
			return a
		}else{
			return b
		}
	}