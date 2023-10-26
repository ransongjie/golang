package pass1

/**
 * 剑指 Offer II 086. 分割回文子字符串
 * 给定一个字符串 s ，请将 s 分割成一些子串，使每个子串都是 回文串 ，返回 s 所有可能的分割方案。
 */

     /**
     * 动态规划+回溯法
     * 动态规划：阶段化，动态规划数组
     * 回溯法：深度优先，回溯
     */
 func partition(s string) [][]string {
	//fss[i][j] i~j是否为回文串
	var pss [][]bool=make([][]bool,len(s))
	var lss [][]string=[][]string{}
	var ls []string =[]string{}
	for i:=0;i<len(s);i++{
		pss[i]=make([]bool,len(s))
	}
	for i:=0;i<len(s);i++{
		for j:=0;j<len(s);j++{
			pss[i][j]=true
		}
	}

	//动态规划，判断回文串 
	var cs []rune=[]rune(s)
	//i从len~0, j从i~len
	for i:=len(s)-1;i>=0;i--{
		for j:=i+1;j<len(s);j++{
			if cs[i]==cs[j]&&pss[i+1][j-1] {
				pss[i][j]=true
			}else{
				pss[i][j]=false
			}
		}
	}

	//回溯法，分割回文串
	var backTrack func(string,int)
	backTrack=func(s string,i int){
		//终止递归
		if i==len(s){
			var ls2 []string=make([]string,len(ls))
			copy(ls2,ls)
			lss=append(lss,ls2);
			return
		}

		//回溯
		/**
         * 0~n-1: 全是1个字母的回文串，有1个2个字母的回文串其它全是1个字母的回文串，有2个2个字母的回文串其它全是1个字母的回文串
		 * 0~n-1: abcd, ab cd, a bc d, ab c d, a bcd, abc d, abcd
         * list是一种分割方案，不同的list中子串可能相同
         * 例如：[["g","o","o","g","l","e"],["g","oo","g","l","e"],["goog","l","e"]]
         */
		 for j:=i;j<len(s);j++{
			if pss[i][j]{
				ls=append(ls,string(cs[i:j+1]))
				//深度
				backTrack(s,j+1)
				//回溯
				ls=ls[0:len(ls)-1]
			}
		 }
	}
	backTrack(s,0)

	//结果	
	return lss	
}

    /**
     * 回溯法+记忆搜索
     * 记忆搜索优化了 partition1() 中的判断回文串的两重for循环
     * xcrj 记忆搜索优势，记录了搜索的中间结果，例如，判断s[i]~s[j]是否为回文串的中间s[i+1]~s[j-1]是否为回文串
     */
func partition2(s string) [][]string {
	//fss[i][j] i~j是否为回文串, 0 表示未搜索，1 表示是回文串，-1 表示不是回文串
	var mem [][]int=make([][]int,len(s))
	var lss [][]string=[][]string{}
	var ls []string =[]string{}
	var cs []rune=[]rune(s)
	for i:=0;i<len(s);i++{
		mem[i]=make([]int,len(s))
	}

	//记忆搜索i~j是否回文串
	var memSearch func(string,int,int) int
	memSearch=func(s string, i int,j int) int{
		//已经处理过
		if mem[i][j]!=0{
			return mem[i][j]
		}

		//中心拓展，1个字母
		if i>=j{
			mem[i][j]=1
			return mem[i][j]
		}

		//中心拓展
		if cs[i]==cs[j]{
			mem[i][j]=memSearch(s,i+1,j-1)
			return mem[i][j]
		}

		//
		mem[i][j]=-1
		return mem[i][j]
	}

	//回溯法 模板
	var backTrack func(string,int)
	backTrack=func(s string, i int){
		//递归出口
		if i==len(s){
			var ls2 []string=make([]string,len(ls))
			copy(ls2,ls)
			lss=append(lss,ls2)
			return
		}

		//回溯
		for j:=i;j<len(s);j++{
			if memSearch(s,i,j)==1{
				ls=append(ls,string(cs[i:j+1]))
				//深度
				backTrack(s,j+1)
				//回溯
				ls=ls[0:len(ls)-1]
			}
		}
	}
	backTrack(s,0)

	//
	return lss
}