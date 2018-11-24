package main
import "fmt"

// ListNode Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	remain := head
	tail := remain
	i := 1
	for i <= k && remain != nil {
		tail = remain
		remain = remain.Next
		i++
	}

	if i < k + 1 {
		return head
	}

	tail.Next = nil

	h, t := reverse(head)
	t.Next = reverseKGroup(remain, k)
	
	return h
}

func reverse(head *ListNode) (*ListNode, *ListNode) {
	e := head
	var p *ListNode
	c := head
	
	for c != nil {
		t := c.Next
		c.Next = p
		p = c
		c = t		
	}

	return p, e 
}

func main() {
	a5 := ListNode {Val: 5}
	a4 := ListNode {Val: 4, Next: &a5}
	a3 := ListNode {Val: 3, Next: &a4}
	a2 := ListNode {Val: 2, Next: &a3}
	a1 := ListNode {Val: 1, Next: &a2}

	m := reverseKGroup(&a1, 4)
	c := m
	for c != nil {
		fmt.Println(c.Val)
		c = c.Next
	}
}