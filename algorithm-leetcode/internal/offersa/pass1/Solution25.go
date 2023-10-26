package pass1

/**
 * 剑指 Offer II 025. 链表中的两数相加
 * 给定两个 非空链表 l1和 l2来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。
 * 可以假设除了数字 0 之外，这两个数字都不会以零开头。
 */

     /**
     * 反转链表
     * 头插法构建结果链表
     * 两数相加可能有进位
     * @param l1
     * @param l2
     * @return
     */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode=nil
	//
	var p1 *ListNode=reverse(l1)
	var p2 *ListNode=reverse(l2)
	// 商和余数
	var quo int=0
	var rem int=0
	// 
	for p1!=nil||p2!=nil||quo!=0{
		//
		var a int=0
		if p1!=nil{
			a=p1.Val
		}
		var b int=0
		if p2!=nil{
			b=p2.Val
		}
		var c =a+b+quo
		quo=c/10
		rem=c%10
		// xcrj golang &ListNode{rem,head}
		var p *ListNode=&ListNode{Val:rem,Next:head}
		head=p
		//
		if(p1!=nil){
			p1=p1.Next
		}
		if(p2!=nil){
			p2=p2.Next
		}
	}
	//
	return head
}

func reverse(head *ListNode) *ListNode{
	if head==nil {
		return nil
	}

	if head.Next==nil {
		return head
	}

	var newHead *ListNode=reverse(head.Next);
	head.Next.Next=head
	head.Next=nil
	return newHead
}
    /**
     * 栈存值
     * 头插法构建结果链表
     * 两数相加可能有进位
     * @param l1
     * @param l2
     * @return
     */
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode=nil
	//
	var stack1 []int=[]int{}
	var stack2 []int=[]int{}
	for l1!=nil{
		stack1=append(stack1,l1.Val)
		l1=l1.Next
	}
	for l2!=nil{
		stack2=append(stack2,l2.Val)
		l2=l2.Next
	}
	//
	var quo int=0
	var rem int=0
	for len(stack1)!=0||len(stack2)!=0||quo!=0{
		//
		var a int=0
		if len(stack1)!=0{
			a=stack1[len(stack1)-1]
			stack1=stack1[:len(stack1)-1]
		}
		var b int=0
		if len(stack2)!=0{
			b=stack2[len(stack2)-1]
			stack2=stack2[:len(stack2)-1]
		}
		var c int =a+b+quo
		quo=c/10
		rem=c%10
		// xcrj golang &取地址符号
		var p *ListNode=&ListNode{rem,head}
		head=p
	}

	//
	return head
}