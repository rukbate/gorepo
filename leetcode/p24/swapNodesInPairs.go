package main
import "fmt"

// ListNode Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}
 
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}	

	h := ListNode {}
	h.Next = head
	p := &h
	m := head

	for m != nil {
		if m.Next != nil {
			n := m.Next
		
			m.Next = n.Next
			p.Next = n
			n.Next = m

			p = m
			m = m.Next
		} else {
			break
		}
	}

	return h.Next
}

func main() {
	a4 := ListNode {Val: 4}
	a3 := ListNode {Val: 3, Next: &a4}
	a2 := ListNode {Val: 2, Next: &a3}
	a1 := ListNode {Val: 1, Next: &a2}

	m := swapPairs(&a1)
	c := m
	for c != nil {
		fmt.Println(c.Val)
		c = c.Next
	}
}