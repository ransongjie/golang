package pass1

/**
 * 剑指 Offer II 022. 链表中环的入口节点
 */

 type void22 struct{}
 var vv void22 
     /**
     * 有环 问题转换 访问了以前被访问过的结点
     * - 使用hash表记录之前访问过的结点
     * 
     * @param head
     * @return
     */
func detectCycle(head *ListNode) *ListNode {
    //特殊情况
	if head==nil {
		return nil
	}
	//
	var p *ListNode=head
	var set map[*ListNode]void22=make(map[*ListNode]void22)
	for p!=nil{
		_,ok:=set[p]
		if ok {
			return p
		}
		set[p]=vv
		p=p.Next
	}
	return nil
}

    /**
     * 快慢指针
     * - fast指针是slow指针的2倍速
     * - slow指针和fast指针相遇
     * - a是head到环开始结点的距离
     * - b是环开始结点到相遇结点的距离
     * - c是fast指针从相遇结点再走到相遇结点多走的距离
     * - 2(a+b)=(a+b+c+b)》a=c
     * - fast指针和slow指针相遇之后，p指针从头结点出发到与slow指针相遇的结点即环的入口结点
     */
func detectCycle1(head *ListNode) *ListNode {
	//
	if head==nil{
		return nil
	}
	//
	var fast *ListNode=head
	var slow *ListNode=head
	for fast != nil{
		// fast.Next.Next 能够为nil
		if fast.Next!=nil{
			fast=fast.Next.Next
		}else{ // 不是含有环的链表
			return nil
		}
		slow=slow.Next
		//
		if fast==slow{
			var p *ListNode=head
			// 不等继续往下走
			for p!=slow{
				p=p.Next
				slow=slow.Next
			}
			return p
		}
	}
	return nil
}