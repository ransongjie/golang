package pass1
/**
 * 剑指 Offer II 098. 机器人行走路径的数目
 * 注意，机器人只能往下或右走
 * 输入，方格大小m,n
 * 输出，从方格左上角走到方格右下角的路径数量
 */

    /**
     * dp[i][j], 从(0,0)走到(i,j)的路径数目
     * dp[i][0]=1;//往下走
     * dp[0][j]=1;//往右走
     * dp[i][j]=dp[i-1][j]+dp[i][j-1];//从上走到(i,j) 从左走到(i,j)
     * @param m
     * @param n
     * @return
     */
func uniquePaths(m int, n int) int {
	//初始化
	dp:=make([][]int,m)
	for i:=0;i<m;i++{
		dp[i]=make([]int,n)
	}
	for i:=0;i<m;i++{//只有1条路往下走
		dp[i][0]=1
	}
	for j:=0;j<n;j++{
		dp[0][j]=1
	}

	//状态转移
	for i:=1;i<m;i++{
		for j:=1;j<n;j++{
			dp[i][j]=dp[i-1][j]+dp[i][j-1]
		}
	}

	//
	return dp[m-1][n-1]
}