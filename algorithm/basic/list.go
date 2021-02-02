package basic

// List 链表
type List struct {
	head *node
	size int
}

// NewList 创建链表
func NewList() *List {
	return &List{}
}

// Add 增加一个元素
func (l *List) Add(element interface{}) {
	n := &node{data: element, next: nil}
	if l.head == nil {
		l.head = n
	} else {
		n.next = l.head
		l.head = n
	}
	l.size++
}

// Size 获取元素数量
func (l *List) Size() int {
	return l.size
}

// Iterator 获取迭代器
func (l *List) Iterator() *ListIterator {
	return &ListIterator{
		cur: l.head,
	}
}

//--------------------------------------------//

// ListIterator 链表迭代器
type ListIterator struct {
	cur *node
}

// HasNext 有下一个元素
func (i *ListIterator) HasNext() bool {
	return i.cur != nil
}

// Next 访问下一个元素
func (i *ListIterator) Next() interface{} {
	element := i.cur.data
	i.cur = i.cur.next
	return element
}
