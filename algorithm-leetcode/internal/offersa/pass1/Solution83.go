package pass1

/**
 * 全排列，不含重复元素
 */
func permute(nums []int) [][]int {
 	//后面的元素逐个放到头部，第i个元素与first元素交换不交换
 	var lss [][]int=[][]int{}
 	var ls []int=[]int{}
	for _,num:=range nums{
		ls=append(ls,num)
	}

	var dfs func(int,int)
	dfs=func(limit int,first int){
		//回退
		if first==limit{
			var ls2 []int=make([]int,len(ls))
			copy(ls2,ls)
			lss=append(lss,ls2)
			return
		}

		//递归回溯
		for i:=first;i<limit;i++{
			ls[first],ls[i]=ls[i],ls[first]
			dfs(limit,first+1)
			ls[first],ls[i]=ls[i],ls[first]
		}
	}
	dfs(len(nums),0)

 	return lss
}