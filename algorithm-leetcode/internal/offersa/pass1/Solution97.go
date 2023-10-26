package pass1

/**
 * 剑指 Offer II 097. 子序列的数目
 * t在s中出现的次数，t可不连续
 */

    /**
     * s中t出现的次数，s的子串中t的子串出现的次数
     * dp[i][j]=s[i:]中t[j:]出现的次数
     * - s[i]==s[j]: dp[i][j]=dp[i+1][j+1]+dp[i+1][j]
     * - s[i]!=s[j]: dp[i][j]=dp[i+1][j]
     */
	 func numDistinct(s string, t string) int {
		//
		if len(t)>len(s){
			return 0
		}
	
		//初始化，s的任意字串都包含t的空串
		dp:=make([][]int,len(s)+1)//字符个数 从0到n
		for i:=0;i<len(dp);i++{
			dp[i]=make([]int,len(t)+1)
		}
		for i:=0;i<len(dp);i++{
			dp[i][len(t)]=1
		}
	
		//状态转移
		for i:=len(s)-1;i>=0;i--{
			for j:=len(t)-1;j>=0;j--{
				if s[i]==t[j]{
					// 取用相等字符或者不去用相等字符
					dp[i][j]=dp[i+1][j+1]+dp[i+1][j]
				}else{
					dp[i][j]=dp[i+1][j];
				}
			}
		}
	
		//
		return dp[0][0];
	}