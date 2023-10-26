package pass1
/**
 * 剑指 Offer II 057. 值和下标之差都在给定的范围内
 * 给你一个整数数组 nums 和两个整数k 和 t 
 * abs(nums[i] - nums[j]) <= t 
 * abs(i - j) <= k 
 */
	/**
     * 桶排序+滑动窗口
     * 每个桶的容量为t+1
     * - Map<桶编号, nums[i]>
     * - nums[i], nums[j]在同一个桶里, nums[i]-nums[j]一定<=t
     * - nums[i], nums[j]在下一个桶里, 需要比较nums[i]-nums[j]<=t 
     * - nums[i], nums[j]在上一个桶里, 需要比较nums[i]-nums[j]<=t 
     * @param nums
     * @param k
     * @param t
     * @return
     */
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	var getBucketId func(int,int)int
	getBucketId=func(x int,v int)int{
		if x>=0{
			return x/v
		}else{
			// x=-1, 需要放到-1这个桶, 而不是0这个桶
			return ((x+1)/v)-1
		}
	}
	// xcrj golang初始化空map
	var bucket map[int64]int64=map[int64]int64{}
	for i:=0;i<len(nums);i++{
		var bucketId=getBucketId(nums[i],t+1)
		var numi int64=int64(nums[i])
		//
		_,ok:=bucket[int64(bucketId)]
		if ok{
			return true
		}
		numj,ok:=bucket[int64(bucketId-1)]
		if ok{
			var absij int64=numi-numj
			if numi-numj<0{
				absij=numj-numi
			}
			if absij<=int64(t){
				return true
			}
		}
		numj,ok=bucket[int64(bucketId+1)]
		if ok{
			var absij int64=numi-numj
			if numi-numj<0{
				absij=numj-numi
			}
			if absij<=int64(t){
				return true
			}
		}
		//
		bucket[int64(bucketId)]=numi
		//
		if i>=k{
			// xcrj golang map delete
			delete(bucket,int64(getBucketId(nums[i-k],t+1)))
		}
	}
	//
	return false
}

