/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    if head==nil || head.Next==nil {return false}
	t1, t2 := head, head.Next
	for t2!=nil && t2.Next!=nil{
		if t1==t2{return true}
		t1=t1.Next
		t2=t2.Next.Next
	}
	return false
}
