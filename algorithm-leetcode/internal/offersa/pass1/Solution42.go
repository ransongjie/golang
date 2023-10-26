package pass1
//xcrj 使用slice实现队列
type RecentCounter struct {
	queue []int
}

//静态方法 构造结构体对象
func Constructor42() RecentCounter {
	var que []int=[]int{}
	return RecentCounter{que}
}


func (this *RecentCounter) Ping(t int) int {
	//入队
	this.queue=append(this.queue,t)
	//对头值
	for this.queue[0]<t-3000{
		//出队
		this.queue=this.queue[1:]
	}
	return len(this.queue)
}