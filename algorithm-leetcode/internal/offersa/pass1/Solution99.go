package pass1

/**
 * 剑指 Offer II 099. 机器人行走最小路径之和
 * 注意，机器人只能往下或右走
 * 输入，grid[i][j]=权重
 * 输出，从左上角 到 右下角 的最小路径权重和
 */

    /**
     * dp[i][j]=从(0,0)走到(i,j)的最小路径和
     * dp[0][0]=grid[0][0];
     * dp[0][j]=dp[0][j-1]+grid[0][j];//往下走
     * dp[i][0]=dp[i-1][0]+grid[i][0];//往右走
     * dp[i][j]=min(dp[i-1][j],dp[i][j-1])+grid[i][j];//从上走到(i,j) 从左走到(i,j)的最小值
     * @param grid
     * @return
     */

func minPathSum(grid [][]int) int {
	//特殊情况
	if grid==nil{
		return 0
	}
	if len(grid)==0{
		return 0
	}
	if len(grid[0])==0{
		return 0
	}

	//初始状态
	m:=len(grid)
	n:=len(grid[0])
	dp:=make([][]int,m)
	for i:=0;i<m;i++{
		dp[i]=make([]int,n)
	}
	dp[0][0]=grid[0][0]// 起点
	for i:=1;i<m;i++{// 一直往下走
		dp[i][0]=dp[i-1][0]+grid[i][0]
	}
	for j:=1;j<n;j++{// 一直往右走
		dp[0][j]=dp[0][j-1]+grid[0][j]
	}

	//状态转移
	for i:=1;i<m;i++{
		for j:=1;j<n;j++{
			dp[i][j]=min99(dp[i-1][j],dp[i][j-1])+grid[i][j]
		}
	}

	//
	return dp[m-1][n-1]
}

func min99(a,b int)int{
	if a<=b{
		return a
	}else{
		return b
	}
}