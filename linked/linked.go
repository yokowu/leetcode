package linked

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(val int) *ListNode {
	return &ListNode{
		Val: val,
	}
}

func AppendNode(head *ListNode, node *ListNode) *ListNode {
	if head == nil {
		return node
	}

	cur := head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = node
	return head
}

func Append(head *ListNode, val int) *ListNode {
	n := &ListNode{Val: val}
	if head == nil {
		head = n
	}

	cur := head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = n
	return head
}

func Show(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}

func RemoveElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	for tmp := dummy; tmp.Next != nil; {
		if tmp.Next.Val == val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}
	return dummy.Next
}

func MergeTwoLists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	prev := dummy
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}

	if l1 != nil {
		prev.Next = l1
	} else {
		prev.Next = l2
	}
	return dummy.Next
}

func ReverseList(head *ListNode) *ListNode {
	var prev, cur *ListNode = nil, head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

func Center(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]*ListNode)

	for headA != nil {
		m[headA] = headA
		headA = headA.Next
	}

	for headB != nil {
		if n, ok := m[headB]; ok {
			return n
		}
		headB = headB.Next
	}
	return nil
}

func GetIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}

func IsPalindrome(head *ListNode) bool {
	list := make([]int, 0)
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}

	for i, j := 0, len(list)-1; i <= j; i, j = i+1, j-1 {
		if list[i] != list[j] {
			return false
		}
	}
	return true
}

func IsPalindrome2(head *ListNode) bool {
	center := Center(head)
	reverse := ReverseList(center.Next)

	p1 := head
	p2 := reverse
	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	center.Next = ReverseList(reverse)
	return true
}

func DeleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func MiddleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow := head.Next
	fast := head.Next
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	fast, slow := dummy, dummy
	for fast != nil && n > 0 {
		fast = fast.Next
		n--
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

func RotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}

	n := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		n++
	}
	diff := n - k%n
	if diff == n {
		return head
	}
	tail.Next = head
	for diff > 0 {
		tail = tail.Next
		diff--
	}
	ret := tail.Next
	tail.Next = nil
	return ret
}

func RotateRight2(head *ListNode, k int) *ListNode {
	n := 0
	for tmp := head; tmp != nil; tmp = tmp.Next {
		n++
	}
	step := k % n
	if step == 0 {
		return head
	}

	slow, fast := head, head
	for step > 0 && fast.Next != nil {
		fast = fast.Next
		step--
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	newHead := slow.Next
	slow.Next = nil
	fast.Next = head
	return newHead
}

func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	pre := dummy
	cur := head
	for cur != nil {
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		}
		if pre.Next == cur {
			pre = pre.Next
		} else {
			pre.Next = cur.Next
		}
		cur = cur.Next
	}
	return dummy.Next
}

func Partition(head *ListNode, x int) *ListNode {
	small, large := &ListNode{}, &ListNode{}
	smallH, largeH := small, large
	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			large.Next = head
			large = large.Next
		}
		head = head.Next
	}
	small.Next = largeH.Next
	large.Next = nil
	return smallH.Next
}

func ReverseBetween(head *ListNode, left, right int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	end := pre
	for i := 0; i < right-left+1; i++ {
		end = end.Next
	}

	leftHead := pre.Next
	curr := end.Next

	pre.Next = nil
	end.Next = nil

	reverse(leftHead)

	pre.Next = end
	leftHead.Next = curr
	return dummy.Next
}

func reverse(head *ListNode) {
	pre, cur := &ListNode{}, head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
}

func ReverseBetween2(head *ListNode, left, right int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummy.Next
}
