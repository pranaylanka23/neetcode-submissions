/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
    curr := head
	grp := 0

	for curr!=nil && grp<k{
		curr = curr.Next
		grp++
	}

	if grp==k{
		curr = reverseKGroup(curr, k)
		for grp>0{
			tmp := head.Next
			head.Next = curr
			curr = head
			head = tmp
			grp--
		}
		head = curr
	}
	return head
}
