/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNodes(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	stack := []*ListNode{}
	for head != nil {
		for len(stack) > 0 && stack[len(stack)-1].Val < head.Val {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, head)
		head = head.Next
	}
	var curr *ListNode
	for _, val := range stack {
		if head == nil {
			head = val
			curr = head
		} else {
			curr.Next = val
			curr = curr.Next
		}
	}
	return head
}
