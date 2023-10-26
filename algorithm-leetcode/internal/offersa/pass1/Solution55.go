package pass1

func findBottomLeftValue55(root *TreeNode) (curVal int) {
    curHeight := 0
    var dfs func(*TreeNode, int)
    dfs = func(node *TreeNode, height int) {
        if node == nil {
            return
        }
        height++
        dfs(node.Left, height)
        dfs(node.Right, height)
        if height > curHeight {
            curHeight = height
            curVal = node.Val
        }
    }
    dfs(root, 0)
    return
}