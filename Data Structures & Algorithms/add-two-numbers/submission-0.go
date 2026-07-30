/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    bal := 0
	prev := &ListNode{}
	head := prev
	for l1!=nil || l2 != nil || bal>0{
		if l1!=nil{ 
			bal += l1.Val
			l1 = l1.Next
		}
		if l2!=nil{ 
			bal += l2.Val
			l2 = l2.Next
		}
		head.Next = &ListNode{Val: bal%10}
		head = head.Next
		bal /= 10
	}
	return prev.Next
}
