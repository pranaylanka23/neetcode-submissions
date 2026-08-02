/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists)==0 { return nil}
	for len(lists)>1{
		var mergedLists []*ListNode
		for i:=0; i<len(lists); i+=2{
			head1 := lists[i]
			var head2 *ListNode
			if i+1<len(lists){
				head2= lists[i+1]
			}
			mergedLists = append(mergedLists, merge2Lists(head1,head2))
		}
		lists = mergedLists
	}
	return lists[0]
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