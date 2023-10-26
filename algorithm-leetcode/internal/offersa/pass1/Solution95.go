package pass1

/**
 * 剑指 Offer II 095. 最长公共子序列
 */

    /**
     * 两字符串的最长公共子序列，两字符串的子串的最长公共子序列
        dp[i][j], 子串长度为i, 子串长度为j, 两子串的最长公共子序列
        i=0或j=0, 空串和任何字符串的最长公共子序列都是0
        i>0或j>0, 
        - s1[i-1]=s2[j-1], dp[i][j]=dp[i−1][j−1]+1
        - s1[i-1]!=s2[j-1], dp[i][j]=max(dp[i−1][j],dp[i][j−1])
     */
func longestCommonSubsequence(text1 string, text2 string) int {
	//初始化 dp[i][0]=0 dp[0][j]=0 空串长度为0
	dp:=make([][]int,len(text1)+1)
	for i:=0;i<len(dp);i++{
		dp[i]=make([]int,len(text2)+1)
	}

	// i和j代表子串长度
	for i:=1;i<=len(text1);i++{
		for j:=1;j<=len(text2);j++{
			if text1[i-1]==text2[j-1]{
				dp[i][j]=dp[i-1][j-1]+1
			}else{
				dp[i][j]=max4(dp[i-1][j],dp[i][j-1])
			}
		}
	}

	//
	return dp[len(text1)][len(text2)]
}

func max4(a,b int)int{
	if a>=b{
		return a
	}else{
		return b
	}
}