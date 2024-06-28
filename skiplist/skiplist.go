package skiplist

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	value int
	next  []*Node
}

type SkipList struct {
	rand     *rand.Rand
	head     *Node
	maxLevel int
}

func NewSkipList(maxLevel int) *SkipList {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	return &SkipList{
		rand: rand,
		head: &Node{
			next: make([]*Node, maxLevel),
		},
		maxLevel: maxLevel,
	}
}

func (sl SkipList) randomLevel() int {
	level := 1
	for ; level < sl.maxLevel; level++ {
		if sl.rand.Float32() > 0.5 {
			break
		}
	}
	return level
}

func (sl *SkipList) Insert(value int) {
	update := make([]*Node, sl.maxLevel)
	prev := sl.head
	next := sl.head
	for i := sl.maxLevel - 1; i >= 0; i-- {
		next = prev.next[i]
		for next != nil && next.value < value {
			prev = next
			next = next.next[i]
		}
		update[i] = prev
	}
	level := sl.randomLevel()

	xnode := &Node{
		value: value,
		next:  make([]*Node, sl.maxLevel),
	}

	//只用更新底层的
	for i := 0; i < level; i++ {
		xnode.next[i] = update[i].next[i]
		update[i].next[i] = xnode
	}
}

func (sl *SkipList) Delete(value int) {
	update := make([]*Node, sl.maxLevel)
	prev, next := sl.head, sl.head

	for i := sl.maxLevel - 1; i >= 0; i-- {
		next = prev.next[i]
		for next != nil && next.value < value {
			prev = next
			next = next.next[i]
		}
		update[i] = prev
	}
	current := next
	if current == nil || current.value != value {
		return
	}
	for i := 0; i < sl.maxLevel && update[i].next[i] != nil; i++ {
		update[i].next[i] = current.next[i]
	}
}

func (sl *SkipList) Search(value int) *Node {
	//update := make([]*Node, sl.maxLevel)
	//current := sl.head
	//for i := sl.maxLevel - 1; i >= 0; i-- {
	//	if current.next[i] == nil {
	//		continue
	//	}
	//	current = current.next[i]
	//	for current != nil && current.value < value {
	//		current = current.next[i]
	//	}
	//}

	prev, next := sl.head, sl.head

	for i := sl.maxLevel - 1; i >= 0; i-- {
		next = prev.next[i]
		for next != nil && next.value < value {
			prev = next
			next = next.next[i]
		}
	}
	current := next

	if current == nil || current.value != value {
		fmt.Println("not found!!!", value)
		return nil
	}
	fmt.Println("found ", value, " !!!")

	return current
}

func (sl *SkipList) String() string {
	values := make([]int, 0)
	current := sl.head
	for ; current.next[0] != nil; current = current.next[0] {
		values = append(values, current.next[0].value)
	}
	return fmt.Sprintf("%v\n", values)
}
