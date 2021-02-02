package basic

// Stack 栈
type Stack struct {
	top  *node
	size int
}

// NewStack 创建stack
func NewStack() *Stack {
	return &Stack{}
}

// Push 压入
func (s *Stack) Push(element interface{}) {
	n := &node{data: element, next: nil}
	if s.top == nil {
		s.top = n
	} else {
		n.next = s.top
		s.top = n
	}
	s.size++
}

// Pop 弹出
func (s *Stack) Pop() interface{} {
	if s.top == nil {
		return nil
	}
	n := s.top
	s.top = n.next
	n.next = nil
	s.size--
	return n.data
}

// Top 获取栈顶元素
func (s *Stack) Top() interface{} {
	if s.top == nil {
		return nil
	}
	return s.top.data
}

// Size 获取元素数量
func (s *Stack) Size() int {
	return s.size
}

// Iterator 获取迭代器
func (s *Stack) Iterator() *StackIterator {
	return &StackIterator{
		cur: s.top,
	}
}

//---------------------------------------------//

// StackIterator 栈迭代器
type StackIterator struct {
	cur *node
}

// HasNext 有下一个元素
func (i *StackIterator) HasNext() bool {
	return i.cur != nil
}

// Next 访问下一个元素
func (i *StackIterator) Next() interface{} {
	element := i.cur.data
	i.cur = i.cur.next
	return element
}
