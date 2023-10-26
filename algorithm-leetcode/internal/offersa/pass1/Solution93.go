package pass1

/**
 * 剑指 Offer II 093. 最长斐波那契数列
 * 找到 arr 中最长的斐波那契式的子序列的长度。如果一个不存在，返回  0 。
 * 斐波那契式子序列
 * - arr[i]>arr[j]>arr[k]且arr[k]+arr[j]=arr[i]
 */

    /**
     * dp[i][j]: 以arr[i],arr[j]为最后两个数字的最长斐波拉契子序列
     * int k=map.getOrDefault(arr[j]-arr[i], -1);
	 * - k>=0: dp[i][j]=Math.max(3,dp[k][i]+1);
     * - k<0: dp[i][j]=dp[i][j]
     * @param arr
     * @return
     */
 func lenLongestFibSubseq(arr []int) int {
	//map<arr[i],idx>
	mp:=map[int]int{}
	for i,a:=range arr{
		mp[a]=i
	}

	//dp中元素值都为0
	dp:=make([][]int,len(arr))
	for i:=0;i<len(dp);i++{
		dp[i]=make([]int,len(arr))
	}
	
	//arr[k],...,arr[i],arr[j]
	maxr:=0
	for j:=0;j<len(arr);j++{ //先定最后1个元素i，再定前1个元素j
		for i:=j;i>=0&&2*arr[i]>arr[j];i--{
			k,ok:=mp[arr[j]-arr[i]]
			if ok{
				dp[i][j]=max3(3,dp[k][i]+1)
			}
			maxr=max3(maxr,dp[i][j])
		}
	}

	//
	return maxr
}

func max3(a,b int)int{
	if a>=b{
		return a
	}else{
		return b
	}
}