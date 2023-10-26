package pass1
import(
	"math"
)
/**
 * 剑指 Offer II 101. 分割等和子集，选取数字和=数组和一半
  * 这道题与传统的「0−1背包问题」的区别：
 * - weight sum <= bag availabe weight:传统的「0−1背包问题」要求选取的物品的重量之和“不超过”背包的总容量，
 * - weight sum = bag availabe weight/2:这道题则要求选取的数字的和“恰好等于”整个数组的元素和的一半
 */

    /**
     * 从数组中选取的数之和=数组和的一半=目标值。从数组的子序列中选取的数之和=目标值的子值
     * dp[i][j]= nums[0]~nums[i]选取的数=j
     * dp[i][0]=true;//从数组子序列中1个都不选择之和=0
     * dp[0][nums[0]]=true;//nums[0:0]选择了之和=nums[0]
     * dp[i][j]
     * - j<nums[i]: 不能选nums[i], dp[i][j]=dp[i-1][j]
     * - j>=nums[i]: 可选可不选nums[i], dp[i][j]=dp[i-1][j-nums[i]]|dp[i-1][j]
     */
func canPartition(nums []int) bool {
	// 特殊情况
	if nums==nil{
		return false
	}
	if len(nums)<2{
		return false
	}
	sum:=sum101(nums)
	if sum%2==1{//奇数
		return false
	}
	max:=max101(nums)
	if max>sum/2{//最大值
		return false
	}
	
	//初始状态
	m:=len(nums)
	n:=sum/2+1//dp[m][sum/2] 从0~m中选取的数之和等于sum/2 存在与否
	dp:=make([][]bool,m)
	for i:=0;i<m;i++{
		dp[i]=make([]bool,n)
	}
	for i:=0;i<m;i++{//从0~i中选取的数之和等于0 存在与否 一个元素都不选
		dp[i][0]=true
	}
	dp[0][nums[0]]=true//从0~0中选取的数之和等于nums[0] 存在与否 就选择nums[0]

	//状态转移
	for i:=1;i<m;i++{
		for j:=1;j<n;j++{
			if j<nums[i]{
				dp[i][j]=dp[i-1][j]//不能选第i个元素，往前看
			}else{
				dp[i][j]=dp[i-1][j-nums[i]]||dp[i-1][j]//可选可不选第i个元素
			}
		}
	}

	//
	return dp[m-1][n-1]
}

func sum101(nums []int) (sum int){
	for _,num:=range nums{
		sum+=num
	}
	return
}

func max101(nums []int) (max int){
	max=math.MinInt
	for _,num:=range nums{
		if max<num{
			max=num
		}
	}
	return
}