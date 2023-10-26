package knapsack_complete

/**
 * https://leetcode.cn/problems/word-break/submissions/
 * 给定一个非空字符串 s 和一个包含非空单词的列表 wordDict，判定 s 是否可以被拆分为一个或多个在字典中出现的单词。
 * <p>
 * bagSize: s.len
 * 物品重量：wordDict中的单词
 * 顺序：求排列
 */
func wordBreak(s string, wordDict []string) bool {
	//问题转换
	var bagSize int = len(s)
	var weight *Set[string] = &Set[string]{} // 是每个单词
	for _, str := range wordDict {
		weight.Add(str)
	}
	// var value = weight
	// var n int = len(weight)
	//dp数组
	//dp[s[0~i-1]即i个字母]=能否被wordDict中单词组成=能否承重weight中物品
	var dp []bool = make([]bool, bagSize+1)
	//初始状态
	dp[0] = true //dp[0个字母]=可以不选物品
	//状态转移
	//排列，s由wordDict中单词组成，与顺序相关
	for i := 1; i < bagSize+1; i++ {
		//!dp[i]已经为true不用再处理
		for j := 0; j < i && !dp[i]; j++ {
			//s[0~i-1]=s[0~j-1]+s[j~i-1]
			_, ok := (*weight)[s[j:i]]
			dp[i] = dp[j] && ok
		}
	}
	//结果
	return dp[bagSize]
}
