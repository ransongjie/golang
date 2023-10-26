package pass1
import(
	"math"
)
/**
 * 剑指 Offer II 039. 直方图最大矩形面积
 * heights[]，数组中的数字用来表示柱状图中各个柱子的高度，每个柱子彼此相邻且宽度为1
 */

     /**
     * 历史单调栈
     * - 单增栈, 放入比栈顶大的值, 可求最近更小的值
     * - 单减栈, 放入比栈顶小的值, 可求最近更大的值
     * @param heights
     * @return
     */
 func largestRectangleArea(heights []int) int {
	//
	var maxa int=math.MinInt
	//ls[i]=l, 第i个直方图的左边界l
	//rs[i]=r, 第i个直方图的有边界r
	var ls []int=make([]int,len(heights))
	var rs []int=make([]int,len(heights))
	//
	var top int =-1
	var stack []int=make([]int,len(heights))
	//
	for i:=0;i<len(heights);i++{
		for top!=-1&&heights[stack[top]]>=heights[i] {
			top--
		}
		if top==-1{
			ls[i]=-1
		}else{
			ls[i]=stack[top]
		}
		//
		top++
		stack[top]=i
	}
	//
	top=-1
	stack=make([]int,len(heights))
	//
	for i:=len(heights)-1;i>=0;i--{
		for top!=-1&&heights[stack[top]]>=heights[i]{
			top--
		}
		if top==-1{
			rs[i]=len(heights)
		}else{
			rs[i]=stack[top]
		}
		//
		top++
		stack[top]=i
	}

	//
	for i:=0;i<len(heights);i++{
		var a int=(rs[i]-ls[i]-1)*heights[i]
		if maxa<a{
			maxa=a
		}
	}
	return maxa
}


    /**
     * 出栈t，确定i的左边界。反之，被出栈t的伪右边界 是i
     * 
     * 比如所有柱子都一样高
     * - 和i同高的最右侧柱子可以得到真右边界
     * - 最终求的是maxArea，有1个真右边界即可
     * - 因此初始时右边界默认为len
     * @param heights
     * @return
     */
func largestRectangleArea1(heights []int) int {
		//
		var maxa int=math.MinInt
		//ls[i]=l, 第i个直方图的左边界l
		//rs[i]=r, 第i个直方图的有边界r
		var ls []int=make([]int,len(heights))
		var rs []int=make([]int,len(heights))
		//因此初始时右边界默认为len
		for i:=0;i<len(rs);i++{
			rs[i]=len(heights)
		}
		//
		var top int =-1
		var stack []int=make([]int,len(heights))
		//
		for i:=0;i<len(heights);i++{
			for top!=-1&&heights[stack[top]]>=heights[i] {
				rs[stack[top]]=i
				top--
			}
			if top==-1{
				ls[i]=-1
			}else{
				ls[i]=stack[top]
			}
			//
			top++
			stack[top]=i
		}
		//
		for i:=0;i<len(heights);i++{
			var a int=(rs[i]-ls[i]-1)*heights[i]
			if maxa<a{
				maxa=a
			}
		}
		return maxa
}