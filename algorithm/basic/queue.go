package basic

// Queue 队列
type Queue struct {
	head, tail *node
	size       int
}

// NewQueue 创建队列
func NewQueue() *Queue {
	return &Queue{}
}

// Enqueue 队尾入队
func (q *Queue) Enqueue(element interface{}) {
	n := &node{data: element, next: nil}
	if q.tail == nil {
		q.head = n
		q.tail = n
	} else {
		q.tail.next = n
		q.tail = n
	}
	q.size++
}

// Dequeue 队头出队
func (q *Queue) Dequeue() interface{} {
	if q.head == nil {
		return nil
	}
	n := q.head
	if q.size == 1 {
		q.head, q.tail = nil, nil
	} else {
		q.head = n.next
	}
	n.next = nil
	q.size--
	return n.data
}

// Size 获取元素数量
func (q *Queue) Size() int {
	return q.size
}

// Iterator 获取迭代器
func (q *Queue) Iterator() *ListIterator {
	return &ListIterator{
		cur: q.head,
	}
}

//--------------------------------------------//

// QueueIterator 队列迭代器
type QueueIterator struct {
	cur *node
}

// HasNext 有下一个元素
func (i *QueueIterator) HasNext() bool {
	return i.cur != nil
}

// Next 访问下一个元素
func (i *QueueIterator) Next() interface{} {
	element := i.cur.data
	i.cur = i.cur.next
	return element
}
