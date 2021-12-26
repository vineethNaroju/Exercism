package linkedlist

import "fmt"

// Define the List and Element types here.

type Element struct {
	value int
	next *Element
}


type List struct {
	size int
	root *Element
}

func New(a []int) *List {

	l := &List{}

	for i:=0; i < len(a); i++ {
		l.Push(a[i])
	}

	return l
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Push(element int) {
	l.size++
	temp := &Element{value: element, next: nil}

	if l.size == 1 {
		l.root = temp
	} else {
		temp.next = l.root
		l.root = temp
	}
}

func (l *List) Pop() (int, error) {
	
	if l.size == 0 {
		return 0, fmt.Errorf("empty list")
	}

	val := l.root.value
	l.root = l.root.next

	l.size--

	return val, nil
}

func (l *List) Array() []int {
	res := make([]int, l.size)
	curr := l.root

	idx := l.size - 1

	for curr != nil {
		res[idx] = curr.value
		idx--
		curr = curr.next
	}
	
	return res
}


// a->b->c->d
// a<-b<-c<-d
func (l *List) Reverse() *List {
	if l.size <= 1 {
		return l
	}

	curr := l.root
	l.root = nil
	prev := l.root

	for curr != nil {
		x := curr.next
		if prev == nil {
			prev = curr
			prev.next = nil
		} else {
			curr.next = prev
			prev = curr
		}
		curr = x
	}

	l.root = prev

	return l
}
