package pass1
import(
	"strconv"
	"strings"
)
/**
 * 剑指 Offer II 087. 复原IP
 * 从s中复原获取IP
 * - 每个整数位于0到255之间组成，且不能含有前导0
 * - 有前导0，这一段就是0
 */

 func restoreIpAddresses(s string) (rs []string) {
	const SEG_NUM int=4
	var segs [4]int=[4]int{}//数组

	var dfs func(int,int)
	dfs=func(i int,k int){
		//回退
		if i==len(s)&&k==SEG_NUM{
			var ts []string=[]string{}
			for _,seg:=range segs{
				var st string=strconv.Itoa(seg)
				ts=append(ts,st)
			}
			// 字符串 join
			var str string=strings.Join(ts,".")
			rs=append(rs,str)
			return
		}
		if k==SEG_NUM&&i<len(s){
			return 
		}
		if k<SEG_NUM&&i==len(s){
			return
		}

		//回溯
		//是0则这一段就是0,处理下一段IP
		if s[i]=='0'{
			segs[k]=0
			dfs(i+1,k+1);
		}

		//不是0
		var sp int=0;
		for j:=i;j<len(s);j++{
			sp=sp*10+int(s[j]-'0')
			if sp>0&&sp<=0xFF{
				segs[k]=sp
				// xcrj 注意 j+1				
				dfs(j+1,k+1)
			}else{
				break
			}
		}
	}
	// 第0个字符，第0段IP开始
	dfs(0,0)

	return
}