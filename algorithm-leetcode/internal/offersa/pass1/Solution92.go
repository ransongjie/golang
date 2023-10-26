package pass1

/**
 * 剑指 Offer II 092. 翻转字符，保持单调递增的最小翻转次数
 */

/**
 * dp0, dp1 第i项目翻转为0 第i项目翻转为1
   第i项目是1翻转为0的仍然递增的最小翻转次数：dp[i][0]=dp[i-1][0]+1
 * 第i项目是0翻转为1的仍然是递增的最小翻转次数：dp[i][1]=min(dp[i-1][0],dp[i-1][1])
 * @param s
 * @return
 */
func minFlipsMonoIncr(s string) int {
	//初始化 从第-1个字符开始
	dp0:=0
	dp1:=0

	//
	for i:=0;i<len(s);i++{
		dpnew0:=dp0
		dpnew1:=min2(dp0,dp1)//将翻转为1，前面为0或1都递增，求最小值，最小翻转次数
		if '1'==s[i]{//将翻转为0,前一项只能翻转为0
			dpnew0=dp0+1
		}

		if '0'==s[i]{//将翻转为1,前一项翻转为0或1都可以
			dpnew1=min2(dp0,dp1)+1
		}
		dp0=dpnew0
		dp1=dpnew1
	}

	//
	return min2(dp0,dp1)
}

func min2(a int,b int)int{
	if a<=b{
		return a
	}else{
		return b
	}
}