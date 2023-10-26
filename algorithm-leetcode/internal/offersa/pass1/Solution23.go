package pass1

    /**
     * 散列表
     * - headA链表中所有结点放入散列表中
     * - headB链表中每个结点是否在散列表中
     * 
     * @param headA
     * @param headB
     * @return
     */
type void23 struct{}
var vv23 void23 
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA==nil{
		return nil
	}
	if headB==nil{
		return nil
	}
	//
	var set map[*ListNode]void23=make(map[*ListNode]void23)
	var ap *ListNode=headA
	var bp *ListNode=headB
	for ap!=nil{
		set[ap]=vv23
		ap=ap.Next
	}
	//
	for bp!=nil{
		_,ok:=set[bp]
		if ok {
			return bp
		}
		bp=bp.Next
	}
	//
	return nil
}
    /**
     * 同速双指针，交叉遍历
     * - 先遍历自己，再遍历他人，相遇即找到 两个链表的第一个重合节点
     * 
     * A链表：a+c
     * B链表：b+c
     * 链表长度一致
     * - A链表：走a
     * - B链表：走b
     * 链表长度不一致
     * - A链表：走a+c+b
     * - B链表：走b+c+a
     * 
     * @param headA
     * @param headB
     * @return
     */
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	if headA==nil{
		return nil
	}
	if headB==nil{
		return nil
	}
	//
	var ap *ListNode=headA
	var bp *ListNode=headB
	for ap!=bp{
		if ap==nil{
			ap=headB
		}else{
			ap=ap.Next
		}
		if bp==nil{
			bp=headA
		}else{
			bp=bp.Next
		}
	}

	return ap
}