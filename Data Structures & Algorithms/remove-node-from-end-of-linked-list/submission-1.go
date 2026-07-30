/**
 * Definition for singly-linked list.
  * type ListNode struct {
   *     Val int
    *     Next *ListNode
	 * }
	  */

	  func removeNthFromEnd(head *ListNode, n int) *ListNode {
	  	l:=0
			curr := head
				for curr!=nil{ 
						l++ 
								curr = curr.Next
									}
										l = l-n
											if l==0 { return head.Next}
												curr=head
													for i:=0; i<l-1; i++{
															curr = curr.Next
																}
																	curr.Next = curr.Next.Next
																		return head
																		}
																		