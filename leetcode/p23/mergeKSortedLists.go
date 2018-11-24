package main
import "fmt"

// ListNode Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if(length == 0) {
		return nil
	}

	h := ListNode {}
	n := &h

	for anyNotNil(lists) {
		m := min(lists)

		n.Next = lists[m]
		lists[m] = lists[m].Next

		n = n.Next
	}
	
	return h.Next
}

func anyNotNil(lists []*ListNode) bool {
	length := len(lists)
	for i := 0; i < length; i++ {
		if lists[i] != nil {
			return true
		}
	}

	return false
}

func min(lists []*ListNode) int {
	minimum := -1
	length := len(lists)
	for i := 0; i < length; i++ {		
		if lists[i] != nil && (minimum == -1 || lists[i].Val < lists[minimum].Val) {
			minimum = i		
		}
	}

	return minimum
}

func main() {
	a3 := ListNode {Val: 5}
	a2 := ListNode {Val: 4, Next: &a3}
	a1 := ListNode {Val: 1, Next: &a2}

	b3 := ListNode {Val: 4}
	b2 := ListNode {Val: 3, Next: &b3}
	b1 := ListNode {Val: 1, Next: &b2}

	c2 := ListNode {Val: 6}
	c1 := ListNode {Val: 2, Next: &c2}

	m := mergeKLists([]*ListNode {&a1, &b1, &c1})
	c := m
	for c != nil {
		fmt.Println(c.Val)
		c = c.Next
	}
}