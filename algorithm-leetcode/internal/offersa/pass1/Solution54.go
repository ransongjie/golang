package pass1
/**
 * 剑指 Offer II 054. 二叉搜索树此结点的值=sum(大于此结点值)
 */
func convertBST(root *TreeNode) *TreeNode {
	var sum int=0
	// xcrj golang 函数变量
	var dfs func(root *TreeNode)
	dfs=func(root *TreeNode){
		if root==nil{
			return
		}
		//
		dfs(root.Right)
		sum+=root.Val
		root.Val=sum
		dfs(root.Left)
	}
	dfs(root)
	//
	return root
}

    /**
     * 动态线索二叉树，在构建线索二叉树的过程中，操作sum
     * 使用node的 空left 指向 node的后继（反中序遍历前驱）
     * Morris动态线索二叉树
     * - O(n)：没有右子树的结点只达到1次，有右子树的结点到达2次
     * - O(1)：利用数的空闲指针，实现空间开销的缩减
     * @param root
     * @return
     */
func convertBST1(root *TreeNode) *TreeNode {
	var p *TreeNode=root
	var sum int=0
	for p!=nil{
		// p结点无前驱, 求和
		if p.Right==nil{
			sum+=p.Val
			p.Val=sum
			p=p.Left
		}else{// p结点有前驱
			// 寻找p结点的前驱
			var pre *TreeNode=getPre(p)
			// 前驱结点左孩子为空，构建线索二叉树，pre结点的左孩子指向p
			if pre.Left==nil{
				// xcrj 从左指针到后继结点
				pre.Left=p
				// xcrj 继续处理前驱结点
				p=p.Right
			}else{ 
				// 前驱结点左孩子不为空，pre.Left等于p，p前驱结点已经被处理，处理p，继续处理p的后继结点
				pre.Left=nil
				sum+=p.Val
				p.Val=sum
				p=p.Left
			}
		}
	}
	//
	return root
}

/**
 *获取当前结点的直接前驱结点
 *右孩子的最左子结点
 */
func getPre(node *TreeNode) *TreeNode{
	var r *TreeNode=node.Right
	// r.Left!=node防止环
	for r.Left!=nil&&r.Left!=node {
		r=r.Left
	}
	return r
}