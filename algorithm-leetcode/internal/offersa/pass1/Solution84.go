package pass1

import(
	"sort"
)
/**
 * 全排列, 含有重复元素，排列不能重复
 * - 先排序nums，将相等元素放到一起
 * - 插空法，将第j个元素插入第i个空中，相等元素插到一起
 */
//将相同元素插入到紧邻的空，打包法
func permuteUnique(nums []int) [][]int {
	var lss [][]int=[][]int{}
	var ls []int=[]int{}
	var added []bool=make([]bool,len(nums))

	sort.Ints(nums)

	var dfs func(int)
	//函数内部不会修改nums，可以不传入指针
	dfs=func(i int){
		//回退
		if i==len(nums){
			// append方法一，不会生成临时slice
			// var ls2 []int=make([]int,len(ls))
			// copy(ls2,ls)
			// lss=append(lss,ls2)
			
			// append方法二，会生成临时slice
			// xcrj golang 中的链式操作
			lss=append(lss,append([]int{},ls...))
			return
		}

		//深度回溯
		//把第j个元素插入到第i个空，相等元素插入到相邻空
		for j:=0;j<len(nums);j++{
			if added[j]||j>0&&nums[j-1]==nums[j]&&!added[j-1]{
				continue
			}

			//xcrj slice 插入到指定位置
			ls=append(ls,0)// xcrj 多append1个元素 为了将 ls[i:]整体位移到 ls[i+1:]
			copy(ls[i+1:],ls[i:])
			ls[i]=nums[j]
			// ls=append(ls,nums[j])
			added[j]=true
			dfs(i+1)

			//xcrj slice 删除指定位置
			ls=append(ls[:i],ls[i+1:]...)
			// ls=ls[:len(ls)-1]
			added[j]=false
		}
	}
	dfs(0)

	return lss
}