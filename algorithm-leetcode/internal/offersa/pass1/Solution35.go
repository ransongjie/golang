package pass1

import(
	"math"
	"sort"
)
/**
 * 剑指 Offer II 035. 24小时制最小时间差
 * 24小时制（"HH:MM"）的时间列表
 * 以分钟数表示的 最小时间差
 */

    /**
     * 鸽巢原理：
     * - 如果要把n+1个物体放进n个盒子，那么至少有1个盒子包含2个或更多的物体。
     * - 共有24*60=1440种不同的分钟数，如果timePoints.size()>1440, 则一定存在相同的分钟数，相同的时间
     * 
     * 排序后
     * - 从左到右两两比较，求最小时间差
     * - 首尾比较，求最小时间差
     * @param timePoints
     * @return
     */
func findMinDifference(timePoints []string) int {
	var minDiff int=math.MaxInt
	//
	if len(timePoints)>24*60{
		return 0
	}
	//
	sort.Strings(timePoints)
	//
	var m0 int=getMinute(timePoints[0])
	var ma int=m0
	for i:=1;i<len(timePoints);i++{
		var mb int=getMinute(timePoints[i])
		//abs
		var diff int =mb-ma
		if mb-ma<0{
			diff=ma-mb
		}
		//min
		if minDiff>diff{
			minDiff=diff
		}
		ma=mb
	}
	// xcrj 00:12 = 24:12
	var diff int=m0+24*60-ma
	//min
	if minDiff>diff {
		minDiff=diff
	}
	//
	return minDiff
}

    /**
     * 获取时间的分钟表示方式
     * @param timePoint
     * @return
     */
func getMinute(timePoint string)int{
	// xcrj 12:23, 中间的:需要跳过
	var cs []rune=[]rune(timePoint)
	var mh int =(int(cs[0]-'0')*10+int(cs[1]-'0'))*60
	var mm int =int(cs[3]-'0')*10+int(cs[4]-'0')
	return mh+mm
}