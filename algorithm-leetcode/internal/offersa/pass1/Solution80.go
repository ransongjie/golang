package pass1

/**
 * 组合数，不含有重复数字，数字不能重复选择&&结果集中组合不能重复
 * 剑指 Offer II 080. 含有 k 个元素的组合
 * 给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
 */
func combine(n int, k int) [][]int {
	var lss [][]int=[][]int{}
	var ls []int=[]int{}

	var dfs func(int,int,int)
	dfs=func(n int,k int,i int){
		//回退
		if len(ls)==k{
			var ls2 []int=make([]int,len(ls))
			copy(ls2,ls)
			lss=append(lss,ls2)
			return
		}
		//已经选择的和未选择的元素个数小于k
		if len(ls)+(n-i+1)<k{
			return
		}
		//
		if i>n {
			return
		}

		//先添加再深度
		//深度回溯
		//选择当前元素，处理下一个元素
		//不选择当前元素，处理下一个元素
		ls=append(ls,i)
		dfs(n,k,i+1)
		ls=ls[0:len(ls)-1]
		dfs(n,k,i+1)
	}
	// 1 ... n
	dfs(n,k,1)
	return lss
}