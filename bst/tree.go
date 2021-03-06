package bst

import (
	"github.com/arthurh0812/datastruct/types"
	"sync"
)

type Tree struct {
	root *Node
	size int64

	duplicates bool

	mu sync.Mutex
}

func (t *Tree) Size() int64 {
	return t.size
}

func (t *Tree) setSize(s int64) {
	t.mu.Lock()
	t.size = s
	t.mu.Unlock()
}

func (t *Tree) increaseSize() {
	t.mu.Lock()
	t.size++
	t.mu.Unlock()
}

func (t *Tree) decreaseSize() {
	t.mu.Lock()
	t.size--
	t.mu.Unlock()
}

func (t *Tree) isEmpty() bool {
	return t == nil || t.size == 0 || t.root == nil
}

func (t *Tree) IsEmpty() bool {
	return t.isEmpty()
}

func (t *Tree) setRoot(n *Node) {
	t.mu.Lock()
	t.root = n
	t.mu.Unlock()
}

func (t *Tree) allowDuplicate() {
	t.mu.Lock()
	t.duplicates = true
	t.mu.Unlock()
}

func (t *Tree) forbidDuplicate() {
	t.mu.Lock()
	t.duplicates = false
	t.mu.Unlock()
}

func (t *Tree) allowsDuplicates() bool {
	return t.duplicates
}

func (t *Tree) traverse(n *Node) (pre *Node) {
	trav := t.root
	for trav != nil {
		pre := trav
		trav = t.chooseNext(trav, n)
		if trav == nil {
			return pre
		}
	}
	return nil // should never happen as traverse is only called if tree is not empty
}

// O(log(n)) average time complexity
func (t *Tree) find(n *Node) *Node {
	trav := t.root
	for trav != nil {
		if trav.isEqualTo(n) {
			return trav
		}
		trav = t.chooseNext(trav, n)
	}
	return nil
}

func (t *Tree) Contains(val types.Value) bool {
	n := &Node{val: val}
	found := t.find(n)
	return found != nil
}

// O(log(n)) average time complexity
func (t *Tree) findPrev(n *Node) (prev, found *Node) {
	trav, prev := t.root, nil
	for trav != nil {
		if trav.isEqualTo(n) {
			return prev, trav
		}
		prev = trav
		trav = t.chooseNext(trav, n)
	}
	return nil, nil
}

func (t *Tree) insert(n *Node) {
	curr := t.root
	for curr != nil {
		prev := curr
		curr = t.chooseNext(curr, n)
		if curr == nil { // last one found, insert after previous
			prev.insert(n)
			t.increaseSize()
			break
		}
	}
}

func (t *Tree) chooseNext(curr, toCompare *Node) *Node {
	if t.allowsDuplicates() && curr.isEqualTo(toCompare) || curr.isGreaterThan(toCompare) {
		return curr.left
	} else if curr.isLessThan(toCompare) {
		return curr.right
	}
	return nil
}

func (t *Tree) Insert(val types.Value) {
	n := &Node{val: val}
	if t.isEmpty() {
		t.init(n)
		return
	}
	t.insert(n)
}

func (t *Tree) init(n *Node) {
	t.setRoot(n)
	t.setSize(1)
}

func (t *Tree) clear() {
	t.mu.Lock()
	t.root = nil
	t.size = 0
	t.mu.Unlock()
}

func (t *Tree) Clear() {
	t.clear()
}