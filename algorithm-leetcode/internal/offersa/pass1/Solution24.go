package pass1

/***
 * 剑指 Offer II 024. 反转链表
 * 给定单链表的头节点 head ，请反转链表，并返回反转后的链表的头节点。
 */

 /**
  * 正序反转
  * - pre current next三指针
  */
 func reverseList(head *ListNode) *ListNode {
	var pre *ListNode=nil
	var cur *ListNode=head
	var next *ListNode=nil
	for cur!=nil{
		//记录下次出发点
		next=cur.Next
		//反转
		cur.Next=pre
		pre=cur
		//进行下一次反转
		cur=next
	}

	return pre
}

 /**
  * 倒序反转
  * - 递归从最后一个结点开始反转
  */
func reverseList1(head *ListNode) *ListNode {
	if head==nil{
		return nil
	}

	if head.Next==nil{
		return head
	}
	// 找到反转之后 链表的头结点
	var newHead *ListNode=reverseList1(head.Next)
	head.Next.Next=head
	head.Next=nil
	// 返回新头结点
	return newHead
}