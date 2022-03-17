package hashset

import "fmt"

type HashSet struct {
	base      int
	hashTable []*node
}

type node struct {
	val  int
	next *node
}

func Constructor() HashSet {
	return HashSet{
		base:      769,
		hashTable: make([]*node, 769),
	}
}

func (h *HashSet) hash(key int) int {
	return key % h.base
}

func (h *HashSet) Add(key int) {
	if h.contains(key) {
		return
	}

	n := &node{val: key}

	head := h.hashTable[h.hash(key)]
	if head == nil {
		h.hashTable[h.hash(key)] = n
		return
	}

	cur := head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = n
	h.hashTable[h.hash(key)] = head
}

func (h *HashSet) Show() {
	for k, head := range h.hashTable {
		cur := head
		for cur != nil {
			fmt.Printf("hash: %d value: %d\n", k, cur.val)
			cur = cur.next
		}
	}
}

func (h *HashSet) Remove(key int) {
	head := h.hashTable[h.hash(key)]
	dummy := &node{next: head}
	for tmp := dummy; tmp.next != nil; {
		if tmp.next.val == key {
			tmp.next = tmp.next.next
		} else {
			tmp = tmp.next
		}
	}
	h.hashTable[h.hash(key)] = dummy.next
}

func (h *HashSet) contains(key int) bool {
	cur := h.hashTable[h.hash(key)]
	for cur != nil {
		if cur.val == key {
			return true
		}
		cur = cur.next
	}
	return false
}
