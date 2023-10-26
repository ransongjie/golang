package pass1

func subsets(nums []int) [][]int {
	var lss [][]int=[][]int{}
	var ls []int=[]int{}
	//
	var dfs func([]int,int)
	// xcrj 函数内部不会对nums进行修改，直接值传递
	// xcrj 如果为了节省空间仍然可以使用引用传递
	dfs=func(nums []int,i int){
		//出口
		if i==len(nums){
			var ls2 []int=make([]int,len(ls))
			copy(ls2,ls)
			lss=append(lss,ls2)
			return
		}

		//深度, 选择当前元素
		ls=append(ls,nums[i])
		dfs(nums,i+1)

		//回溯
		ls=ls[0:len(ls)-1]
		dfs(nums,i+1)
	}
	dfs(nums,0)
	//
	return lss
}