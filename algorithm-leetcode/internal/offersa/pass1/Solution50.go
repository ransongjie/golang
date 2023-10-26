package pass1
/**
 * 剑指 Offer II 050. 向下的路径节点之和
 * 给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
 * 路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
 */

 /**
     * 对前缀和进行回溯
     * 前缀和+回溯
     */
func pathSum(root *TreeNode, targetSum int) int {
	var sumMap map[int64]int=map[int64]int{}
	sumMap[0]=1
	// xcrj golang中必须指明引用传递，否则都是值传递
	var dfs func(*TreeNode, int64,*map[int64]int,int64)int
	dfs=func(p *TreeNode,targetSum int64,sumMap *map[int64]int,sum int64)int{
		//退出
		if p==nil{
			return 0
		}
		//
		sum+=int64(p.Val)
		// xcrj golang (*sumMap)取值
		r,_:=(*sumMap)[sum-targetSum]
		// map的key不存在将返回0
		// 深度
		(*sumMap)[sum]+=1
		r+=dfs(p.Left,targetSum,sumMap,sum)
		r+=dfs(p.Right,targetSum,sumMap,sum)
		// 回溯
		(*sumMap)[sum]-=1
		return r
	}
	return dfs(root,int64(targetSum),&sumMap,0);
}