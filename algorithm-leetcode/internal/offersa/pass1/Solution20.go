package pass1

func countSubstrings(s string) int {
	var cnt int=0
	var cs []rune=[]rune(s)
	for i:=0;i<len(s)*2-1;i++{
		var l int=i/2
		var r int=l+i%2
		for l>=0&&r<len(s)&&cs[l]==cs[r]{
			cnt++
			l--
			r++
		}
	}
	return cnt
}

	/**
     * 在基础之上进行中心拓展
	 * 利用历史最大回文半径进行拓展, xcrj 历史基础
     * 上面的方法，每次从回文中心都是从0开始，没有基础
     * 下面的方法，每次中回文中心都是从f[i]开始，有基础
     * 
     * 利用历史中心拓展信息
     * - 历史最大回文半径的回文中心的右端点 maxr
     * - 历史最大回文半径的回文中心 maxc
     * 
     * 使用动态规划数组f[i]=center_i_max_radius
     * - i>maxr：无法利用之前的所有中心拓展的信息。f[i]=1
     * - i<=maxr：利用i的关于maxc的对称点j f[j]。利用i到maxr的距离 rmax-i+1。f[i]=min{f[j],rmax-i+1}
     * 
     * Manacher可以在线性时间内求解 最长回文子串
     * 
     * @param s
     * @return
     */
func countSubstrings2(s string) int {
	var cnt int=0
	// xcrj 此种求解方法 要求 字符串序列为奇数，s.len*2+1+1+1
    // 处理s，$#c1#,c2#,c3#!
	var cs []rune=[]rune(s)
	var sb []rune=make([]rune,0)
	sb=append(sb,'$')
	sb=append(sb,'#')
	for i:=0;i<len(cs);i++{
		sb=append(sb,cs[i])
		sb=append(sb,'#')
	}
	sb=append(sb,'!')
	// 动态规划数组, 使用slice替代，golang数组必须是常量
	var f []int=make([]int,len(sb)-1)
	// 初始化
	var maxc int=0
	var maxr int=0
	// 开始的$和末尾的!不需要处理
	for i:=1;i<len(sb)-1;i++{
		// xcrj 找基础
		if i>maxr{
			f[i]=1
		}else{
			// min
			if f[2*maxc-i]>(maxr-i+1){
				f[i]=maxr-i+1
			}else{
				f[i]=f[2*maxc-i]
			}
		}
		// 在基础之上进行中心拓展
		for sb[i+f[i]]==sb[i-f[i]]{
			f[i]++
		}
		//更新 maxr maxc
		if i+f[i]-1>maxr{
			maxr=i+f[i]-1
			maxc=i
		}
		//不考虑添加的特殊字符 /2
		cnt+=f[i]/2
	}
	return cnt
}