package pass1
/**
 * 剑指 Offer II 052. 二叉搜索树展平为右斜树
 * - 二叉搜索树就是二叉排序树
 * - 中序遍历二叉搜索数，将得到递增有序的序列
 * 
 * 按中序遍历将其重新排列为一棵递增顺序搜索树
 * - 树中最左边的节点成为树的根节点
 * - 每个节点没有左子节点，只有一个右子节点
 */

	/**
     * 中序遍历二叉搜索结果放入链表中
     * 从链表中构建只有右结点的二叉树
     * @param root
     * @return
     */
func increasingBST(root *TreeNode) *TreeNode {
	//
	var dummy *TreeNode=&TreeNode{Val:0}
	var p *TreeNode=dummy
	//
	var list []int=[]int{}
	// xcrj golang中方法调用 必须指明引用传递，否则都是值传递
	// xcrj golang传递slice
	dfs52(root,&list)
	//
	for _,val:=range list{
		var node *TreeNode =&TreeNode{Val:val}
		p.Right=node
		p=node
	}

	//
	return dummy.Right
}

func dfs52(root *TreeNode,list *[]int){
	if root==nil{
		return
	}
	//
	dfs52(root.Left,list)
	// xcrj *list 值, list是指针是地址值
	*list=append(*list,root.Val)
	dfs52(root.Right,list)
}

    /**
     * 在中序遍历的过程中构建树，类变量记录上一个结点
     * @param root
     * @return
     */
var p *TreeNode
func increasingBST1(root *TreeNode) *TreeNode {
	//
	var dummy *TreeNode=&TreeNode{Val:0}
	p=dummy
	//
	dfs52b(root)
	//
	return dummy.Right
}

func dfs52b(root *TreeNode){
	if root==nil{
		return
	}
	//
	dfs52b(root.Left)
	p.Right=root
	root.Left=nil
	p=root
	dfs52b(root.Right)
}