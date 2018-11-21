package main
import "fmt"


// ListNode Definition for singly-linked list.
type ListNode struct {
   Val int
   Next *ListNode
}
 
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	ordMap := make(map[*ListNode]int)

	var p *ListNode
	c := head
	for c != nil {
		ord := getOrdinal(c, ordMap)
		if ord == n {
			if p != nil {
				p.Next = c.Next			
			} else {
				head = c.Next
			}
			break
		} else {
			p = c
			c = c.Next
		}
	}

	return head
}

func getOrdinal(head *ListNode, ordMap map[*ListNode]int) int {
	cur := head

	if od, found := ordMap[cur]; found {
		return od
	}

	if cur == nil {
		return 0		
	} else if cur.Next == nil {
		return 1
	} else {
		return getOrdinal(head.Next, ordMap) + 1
	}
}

func main() {
	n1 := new(ListNode)
	n1.Val = 1
	
	n2 := new(ListNode)
	n2.Val = 2

	n3 := new(ListNode)
	n3.Val = 3

	n4 := new(ListNode)
	n4.Val = 4

	n5 := new(ListNode)
	n5.Val = 5

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5

	n1.Next = nil

	n := removeNthFromEnd(n1, 1)

	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
}