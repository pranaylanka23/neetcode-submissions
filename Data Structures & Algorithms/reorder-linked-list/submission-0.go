/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	if head==nil || head.Next ==nil {return}
    t1,t2 := head, head.Next
	for t2!=nil && t2.Next!=nil{
		t1=t1.Next
		t2=t2.Next.Next
	}
	t2= t1.Next
	t1.Next=nil
	var prev *ListNode
	for t2!=nil{
		temp := t2.Next
		t2.Next=prev
		prev=t2
		t2=temp
	}
	fir,sec := head, prev
	for sec!=nil{
		t1,t2 = fir.Next, sec.Next
		fir.Next = sec
		sec.Next = t1
		fir,sec = t1,t2
	}
}
