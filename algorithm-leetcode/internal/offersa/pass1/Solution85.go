package pass1

/**
 * 剑指 Offer II 085. 生成匹配的括号
 * 分析：
 * - n=左括号数量=右括号数量
 * - 只有括号，没有数字
 */
// xcrj 直接把返回值声明到函数后面
func generateParenthesis(n int) (rs []string) {
	var l []rune=[]rune{}
	
	var dfs func(int,int)
	// open左括号数量, close右括号数量
	dfs=func(open int,close int){
		//回退
		if 2*n==open+close{
			rs=append(rs,string(l))
			return
		}

		//回溯
		if open<n{
			l=append(l,'(')
			dfs(open+1,close)
			l=l[:len(l)-1]
		}

		if close<open{
			l=append(l,')')
			dfs(open,close+1)
			l=l[:len(l)-1]
		}
	}
	dfs(0,0)

	return
}