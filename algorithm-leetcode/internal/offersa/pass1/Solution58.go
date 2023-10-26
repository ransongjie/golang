package pass1
/**
 * 剑指 Offer II 058. 日程表
 * 时间是半开区间，即 [start, end), 实数 x 的范围为，start <= x < end。
 * 当两个日程安排有一些时间上的交叉时（例如两个日程安排都在同一时间内），就会产生重复预订。
 * 调用book()不产生重复预订返回true，产生重复预定返回false
 */

	/**
     * 存储到list<int[]>后再遍历
     */
type MyCalendar struct {
	list [][]int
}


func Constructor58() MyCalendar {
	// xcrj golang 二维slice创建 与 使用
	return MyCalendar{list:[][]int{}}
}


func (this *MyCalendar) Book(start int, end int) bool {
	for _,se:=range this.list{
		var s int=se[0]
		var e int=se[1]
		//[s1,e1), [s2,e2)
		//不冲突, s1>=e2||s2>=e1
		//冲突, s1<e2&&s2<e1
		if start<e&&s<end{
			return false
		}
	}

	// xcrj golang 二维slice创建 与 使用
	var e []int=[]int{start,end}
	this.list=append(this.list,e)
	return true
}

/**
 *动态线段树，把start,end小线段以二分查找的方式插入到0, 1000000000 大线段中合适的位置
 *合适的位置，start,end小线段包含子线段l,r 
 */
type void58 struct{}
var vv58 void58

type MyCalendar2 struct {
	//线段结点编号
	tree map[int]void58
	//快速查找
	lazy map[int]void58
}


func Constructor58b() MyCalendar2 {
	return MyCalendar2{map[int]void58{},map[int]void58{}}
}


func (this *MyCalendar2) Book(start int, end int) bool {
	//从结点1开始查找，start,end的子线段是否已经在线段树中
	//xcrj 注意 end-1
	if this.query(start,end-1,0,1000000000,1){
		return false
	}
	this.update(start,end-1,0,1000000000,1)
	return true
}

/**
 * 二分查找start,end的子线段l,r是否在线段树中
 * xcrj l,mid;l,r;mid,r，父亲线段包含子线段
 */
func (this *MyCalendar2) query(start int, end int, l int,r int, idx int) bool{
	//查找
	//start,end,l,r||l,r,start,end
	if end<l||start>r{
		return false
	}
	//lazy快速查找
	_,ok:=this.lazy[idx]
	if ok{
		return true
	}
	//start,l,r,end
	if start<=l&&end>=r{
		_,ok=this.tree[idx]
		return ok
	}

	//二分
	var mid int=(l+r)>>1
	//左侧，l,start,end,mid
	if(end<=mid) {
		return this.query(start,end,l,mid,idx*2)
	}
	//右侧，mid,start,end,r
	if(start>mid) {
		return this.query(start,end,mid+1,r,idx*2+1)
	}

	// l,start,mid,end,r
	var lb bool=this.query(start,end,l,mid,idx*2)
	var rb bool=this.query(start,end,mid+1,r,idx*2+1)
	return lb||rb
}

/**
 * 建立线段树，叶子节点是start,end不包含的子区间和包含的子区间
 */
func (this *MyCalendar2) update(start int, end int, l int, r int, idx int){
	//查找
	//start,end,l,r||l,r,start,end
	if end<l||start>r{
		return
	}
	//start,l,r,end
	if start<=l&&end>=r{
		this.tree[idx]=vv58
		this.lazy[idx]=vv58
		return
	}

	//二分
	var mid int=(l+r)>>1
	//l,start,end,mid
	this.update(start,end,l,mid,idx*2)
	//mid,start,end,r
	this.update(start,end,mid+1,r,idx*2+1)
	//
	this.tree[idx]=vv58
	//l,mid;l,r;mid,r，父亲线段包含子线段
	_,okidx2:=this.lazy[2*idx]
	_,okidx21:=this.lazy[2*idx+1]
	if okidx2&&okidx21{
		this.lazy[idx]=vv58
	}
}