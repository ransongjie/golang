package pass1

/**
 * 剑指 Offer II 096. 字符串交织
 * 给定三个字符串 s1、s2、s3，请判断 s3 能不能由 s1 和 s2 交织（交错） 组成。
 * 两个字符串 s 和 t 交织 的定义与过程如下，其中每个字符串都会被分割成若干 非空 子字符串：
 * - s = s1 + s2 + ... + sn
 * - t = t1 + t2 + ... + tm
 * - |n - m| <= 1
 * - 交织 是 s1 + t1 + s2 + t2 + s3 + t3 + ... 或者 t1 + s1 + t2 + s2 + t3 + s3 + ...
 */

    /**
     * dp[i][j], dp[s1的前i个元素][s2的前j个元素]=能否交织为s3的前i+j个元素
     * i>0, dp[i][j]=dp[i][j]||s1.charAt(i-1)==s3.charAt(i+j-1)&&dp[i-1][j]
     * j>0, dp[i][j]=dp[i][j]||s1.charAt(j-1)==s3.charAt(i+j-1)&&dp[i][j-1]
     * @param s1
     * @param s2
     * @param s3
     * @return
     */
	 func isInterleave(s1 string, s2 string, s3 string) bool {
		//特殊情况
		if len(s1)+len(s2)!=len(s3){
			return false
		}
	
		//xcrj 动态规划数组什么时候需要+1, 下标代表元素个数时
		dp:=make([][]bool,len(s1)+1)
		for i:=0;i<len(dp);i++{
			dp[i]=make([]bool,len(s2)+1)
		}
		dp[0][0]=true
	
		//
		for i:=0;i<=len(s1);i++{
			for j:=0;j<=len(s2);j++{
				if i>0{
					dp[i][j]=dp[i][j]||s1[i-1]==s3[i+j-1]&&dp[i-1][j]
				}
	
				if j>0{
					dp[i][j]=dp[i][j]||s2[j-1]==s3[i+j-1]&&dp[i][j-1]
				}
			}
		}
	
		return dp[len(s1)][len(s2)]
	}

    /**
     * 滚动数组优化
     * - 数组f的第i行只和第i−1行相关,去掉i这一维度,使用滚动方式代替行
     * - 滚动，第i行取决于第i-1行，等式右侧的f[j]是代表f[s1的前i-1个元素][s2的前j个元素]，左侧f[j]代表f[s1的前i个元素][s2的前j个元素]
     */
func isInterleave2(s1 string, s2 string, s3 string) bool {
	//特殊情况
	if len(s1)+len(s2)!=len(s3){
		return false
	}

	//xcrj 动态规划数组什么时候需要+1, 下标代表元素个数时
	dp:=make([]bool,len(s2)+1)
	dp[0]=true

	//
	for i:=0;i<=len(s1);i++{
		for j:=0;j<=len(s2);j++{
			if i>0{
				// dp[i][j]=dp[i][j]||s1[i-1]==s3[i+j-1]&&dp[i-1][j]
				dp[j]=s1[i-1]==s3[i+j-1]&&dp[j]
			}

			if j>0{
				// dp[i][j]=dp[i][j]||s2[j-1]==s3[i+j-1]&&dp[i][j-1]
				dp[j]=dp[j]||s2[j-1]==s3[i+j-1]&&dp[j-1]
			}
		}
	}

	return dp[len(s2)]
}