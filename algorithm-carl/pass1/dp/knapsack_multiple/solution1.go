package knapsack_multiple

import (
	"fmt"

	"com.xcrj/golang/algorithm-carl/common"
)

// 完全背包问题，每个物品的重量为优先个
func RunKM() {
	bagSize := 10
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	nums := []int{2, 3, 2}

	ans := test1(bagSize, weight, value, nums)
	fmt.Println(ans)
	ans = test2(bagSize, weight, value, nums)
	fmt.Println(ans)
}

// 转换为01背包问题
func test1(bagSize int, weight []int, value []int, nums []int) int {
	//问题转换
	for i := 0; i < len(nums); i++ {
		for j := 0; j < nums[i]-1; j++ { //-1 总共3件，原weight中已经有1件了
			weight = append(weight, weight[i])
			value = append(value, value[i])
		}
	}
	n := len(weight)
	//dp数组
	dp := make([]int, bagSize+1)
	//初始状态
	dp[0] = 0
	//状态转移
	//先物品再背包。01背包是倒序背包，完全背包是正序背包
	for i := 0; i < n; i++ {
		for j := bagSize; j >= weight[i]; j-- {
			dp[j] = common.Max[int](dp[j], value[i]+dp[j-weight[i]])
		}
	}

	//结果
	return dp[bagSize]
}

func test2(bagSize int, weight []int, value []int, nums []int) int {
	n := len(weight)
	//dp数组
	dp := make([]int, bagSize+1)
	//初始状态
	dp[0] = 0
	//状态转移
	//先物品再背包。01背包是倒序背包，完全背包是正序背包
	for i := 0; i < n; i++ {
		for j := bagSize; j >= weight[i]; j-- {
			for k := 1; k <= nums[i] && j >= k*weight[i]; k++ {
				// k*value[i]
				dp[j] = common.Max[int](dp[j], k*value[i]+dp[j-k*weight[i]])
			}
		}
	}

	//结果
	return dp[bagSize]
}
