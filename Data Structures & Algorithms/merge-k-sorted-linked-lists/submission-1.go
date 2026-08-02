/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists)==0 { return nil}
	head := lists[0]
	for i:=1; i<len(lists);i++{
		head = merge2Lists(lists[i], head)
	}
	return head
}

func merge2Lists(head1, head2 *ListNode) *ListNode{
	if head1==nil{ return head2}
	if head2==nil { return head1}
	if head1.Val>head2.Val{
		head2.Next = merge2Lists(head1, head2.Next)
		return head2
	}
	head1.Next = merge2Lists(head1.Next, head2)
	return head1
}