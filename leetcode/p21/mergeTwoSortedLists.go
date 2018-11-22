package main
import "fmt"

// ListNode Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}
 
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	h := ListNode {}
	n := &h

	l, r := l1, l2
	for l != nil && r != nil {
		if l.Val < r.Val {
			n.Next = l
			l = l.Next
		} else {
			n.Next = r
			r = r.Next
		}

		n = n.Next
	}

	for l != nil {
		n.Next = l
		l = l.Next
		n = n.Next
	}

	for r != nil {
		n.Next = r
		r = r.Next
		n = n.Next
	}

	return h.Next
}

func main() {
	a3 := ListNode {Val: 4}
	a2 := ListNode {Val: 2, Next: &a3}
	a1 := ListNode {Val: 1, Next: &a2}

	b3 := ListNode {Val: 4}
	b2 := ListNode {Val: 3, Next: &b3}
	b1 := ListNode {Val: 1, Next: &b2}

	m := mergeTwoLists(&a1, &b1)
	c := m
	for c != nil {
		fmt.Println(c.Val)
		c = c.Next
	}
}