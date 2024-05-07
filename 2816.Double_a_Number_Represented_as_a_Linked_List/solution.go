/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverse(head *ListNode) *ListNode {
	var node *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = node
		node = head
		head = tmp
	}
	return node
}

func doubleIt(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	head = reverse(head)
	curr := head
	var carry int
	for curr != nil {
		n := curr.Val * 2
		curr.Val = n%10 + carry
		carry = n / 10
		if curr.Next == nil && carry != 0 {
			curr.Next = &ListNode{Val: carry}
			break
		}
		curr = curr.Next
	}
	return reverse(head)
}
