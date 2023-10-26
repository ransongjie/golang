package pass1
import(
	"sort"
)
/**
 * 组合数，含有重复数字，数字不能重复选择&&结果集中组合不能重复
 * 剑指 Offer II 081. 允许重复选择元素的组合
 * - target 选择的组合数之和
 * - candidates 可能含有重复数字。每个数字只能使用一次&&结果集中组合不能重复
 */
func combinationSum2(candidates []int, target int) [][]int {
	var lss [][]int=[][]int{}
	var ls []int=[]int{}
	var nts [][]int=[][]int{}//List<int[]{num,times}>

	sort.Ints(candidates)//golang 排序
	for _,c:=range candidates{
		if len(nts)==0{
			var nt []int=[]int{c,1}
			nts=append(nts,nt)
			continue
		}
		if c==nts[len(nts)-1][0]{
			nts[len(nts)-1][1]++
			continue
		}
		var nt []int=[]int{c,1}
		nts=append(nts,nt)
	}

	var dfs func(int,int)
	dfs=func(target int,i int){
		//回退
		if target==0{
			var ls2 []int=make([]int,len(ls))
			copy(ls2,ls)
			lss=append(lss,ls2)
			return
		}
		if i==len(nts){
			return
		}
		if target<nts[i][0]{
			return
		}

		//深度回溯
		dfs(target,i+1)
		//重复元素一起处理
		var k int=nts[i][1]
		if target/nts[i][0]<nts[i][1]{
			k=target/nts[i][0]
		}
		for j:=1;j<=k;j++{
			ls=append(ls,nts[i][0])
			dfs(target-j*nts[i][0],i+1)
		}
		for j:=1;j<=k;j++{
			ls=ls[0:len(ls)-1]
		}
	}
	dfs(target,0)

	return lss
}