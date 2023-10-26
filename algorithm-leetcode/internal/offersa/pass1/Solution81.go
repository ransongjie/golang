package pass1
/**
 * 组合数，不含有重复数字，数字能够重复选择&&结果集中组合不能重复
 * 剑指 Offer II 081. 允许重复选择元素的组合
 * - target 选择的组合数之和
 * - candidates 不含有重复数字，选择相同元素的数量不同也认为是不同的组合数
 */

func combinationSum(candidates []int, target int) [][]int {
	var lss [][]int=[][]int{}
	var ls []int=[]int{}

	var dfs func(*[]int,int,int)
	dfs=func(candidates *[]int,target int,i int){
		//回退
		if target==0{
			var ls2 []int=make([]int,len(ls))
			copy(ls2,ls)
			lss=append(lss,ls2)
			return
		}
		if i==len((*candidates)){
			return
		}

		//深度回溯
		dfs(candidates,target,i+1)
		if target-(*candidates)[i]>=0{
			ls=append(ls,(*candidates)[i])
			dfs(candidates,target-(*candidates)[i],i)// 可以重复选择
			ls=ls[0:len(ls)-1]
		}
	}
	dfs(&candidates,target,0)

	return lss
}