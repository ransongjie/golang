package pass1
import(
	"math"
)
/**
 * 剑指 Offer II 100. 三角形行走最小路径之和
 * 三角形形状
 * [2]
 * [3,4]
 * [6,5,7]
 * [4,1,8,3]
 * 分析：
 * - 第1行有1个元素，第2行2个元素，...，第n行n个元素
 * - 移动，往下 或 往右下
 * 注意：
 * - Solution99是正方形，到达同一个终点
 * - Solution100是三角形，到达底边，有多个终点
 */
    /**
     * 从(0,0)到底边的最小路径和
     * dp[0][0]=triangle[0][0]
     * dp[i][0]=dp[i-1][0]+triangle[i][j]
     * dp[i][i]=dp[i-1][i-1]+triangle[i][i]
     * dp[i][j]=min(dp[i-1][j],dp[i-1][j-1])+triangle[i][j]
     */
func minimumTotal(triangle [][]int) int {
	//初始状态
	m:=len(triangle)
	dp:=make([][]int,m)
	for i:=0;i<m;i++{
		dp[i]=make([]int,m)
	}
	dp[0][0]=triangle[0][0]

	//状态转移
	for i:=1;i<m;i++{
		//一直往下走
		dp[i][0]=dp[i-1][0]+triangle[i][0]
		for j:=1;j<i;j++{
			dp[i][j]=min100(dp[i-1][j],dp[i-1][j-1])+triangle[i][j]
		}
		//走斜边
		dp[i][i]=dp[i-1][i-1]+triangle[i][i]
	}

	//
	minTotal:=math.MaxInt
	for i:=0;i<m;i++{
		minTotal=min100(minTotal,dp[m-1][i])
	}
	return minTotal
}
func min100(a,b int)int{
	if(a<=b){
		return a
	}else{
		return b
	}
}