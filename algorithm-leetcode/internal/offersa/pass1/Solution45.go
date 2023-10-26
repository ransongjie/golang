package pass1

/**
 * 剑指 Offer II 045. 二叉树最底层最左边的值
 */
    /**
     * 深度优先+最大层数 后序遍历 左右根
     * @param root
     * @return
     */
func findBottomLeftValue(root *TreeNode) int {
	var maxHeight int;
	var maxLeftVal int;

	var dfs func(*TreeNode,int)
	dfs=func(p *TreeNode,height int){
		if p==nil{
			return
		}
		height++
		//
		dfs(p.Left,height)
		dfs(p.Right,height)
		//
		if maxHeight<height{
			maxHeight=height
			maxLeftVal=p.Val
		}
	}
	//
	dfs(root,0)

	return maxLeftVal
}


    /**
     * 广度优先，先入队右子结点，如此最后出队的结点是最左下的结点
     * @param root
     * @return
     */
func findBottomLeftValue1(root *TreeNode) int {
	var r int
	//
	if root==nil{
		return 0
	}
	//
	var que []*TreeNode=[]*TreeNode{}
	que=append(que,root)
	for len(que)>0 {
		var p *TreeNode=que[0]
		r = p.Val
		// xcrj 队列操作
		que=que[1:]
		//
		if p.Right!=nil{
			que=append(que,p.Right)
		}
		if p.Left!=nil{
			que=append(que,p.Left)
		}
	}
	//
	return r
}