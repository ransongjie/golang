package pass1
import(
	"math"
)
/**
 * 剑指 Offer II 094. 最少回文分割
 * 给定一个字符串 s，请将 s 分割成一些子串，使每个子串都是回文串。
 * 返回符合要求的 最少分割次数 。
 * 
 * 输入，字符串s
 * 输出，将s分割成多个回文子串的最少分割次数
 */

 /**
     * 动态规划回文串
     * @param s
     * @return
     */
func minCut(s string) int {
	//xcrj 是否回文串 i~j是回文串，取决于前一个阶段i+1~j-1是回文串，单个字符也是回文串
	var pss [][]bool=make([][]bool,len(s))
	for i:=0;i<len(s);i++{
		pss[i]=make([]bool,len(s))
	}
	//xcrk 最小回文分割, 0~i的最小分割次数，取决于0~j的最少分割次数+1&&j~i是否回文串
	var fs []int=make([]int,len(s))
	
	//默认都是回文串
	for i:=0;i<len(s);i++{
		for j:=0;j<len(s);j++{
			pss[i][j]=true
		}
	}
	//求最小回文分割次数
	for i:=0;i<len(s);i++{
		fs[i]=math.MaxInt
	}

	var cs []rune=[]rune(s)
	//判断回文串
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

	//0~i的最小回文分割，取决于0~j的最小回文分割+1&&j~i是否回文串
	for i:=0;i<len(s);i++{
		//不需要分割
		if pss[0][i]{
			fs[i]=0
		}else{
		//需要分割
			for j:=0;j<i;j++{
				if pss[j+1][i]{
					if fs[j]+1<fs[i] {
						fs[i]=fs[j]+1
					}
				}
			}		
		}
	}
	
	//
	return fs[len(s)-1]
}