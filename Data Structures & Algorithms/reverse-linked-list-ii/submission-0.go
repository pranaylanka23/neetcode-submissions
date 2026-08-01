/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Val:0, Next: head}
	leftPrev := dummy
	curr := head

	for i:=0; i<left-1; i++{
		leftPrev = curr
		curr = curr.Next
	}

	var prev *ListNode
	for i:=0; i<right-left+1; i++{
		tmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
	}

	leftPrev.Next.Next = curr
	leftPrev.Next = prev
	
	return dummy.Next
}


