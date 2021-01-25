package search

const (
	red   = true
	black = false
)

// RBTree 红黑树
type RBTree struct {
	root *node
}

type node struct {
	left, right *node
	key         string
	value       interface{}
	color       bool
	size        int
}

// NewRBTree 创建红黑树
func NewRBTree() *RBTree {
	return &RBTree{}
}

func (n *node) isRed() bool {
	if n == nil {
		return false
	}
	return n.color
}

func rotateLeft(n *node) *node {
	x := n.right
	n.right = x.left
	x.left = n

	x.color = n.color
	n.color = red

	x.size = n.size
	n.size = size(n.left) + 1 + size(n.right)

	return x
}

func rotateRight(n *node) *node {
	x := n.left
	n.left = x.right
	x.right = n

	x.color = n.color
	n.color = red

	x.size = n.size
	n.size = size(n.left) + 1 + size(n.right)

	return x
}

func size(n *node) int {
	if n == nil {
		return 0
	}
	return n.size
}

func flipColors(n *node) {
	n.color = red
	n.left.color = black
	n.right.color = black
}

// Put 插入元素
func (t *RBTree) Put(key string, val interface{}) {
	t.root = put(t.root, key, val)
}

func put(n *node, key string, val interface{}) *node {
	if n == nil {
		return &node{
			key:   key,
			value: val,
			color: red,
			size:  1,
		}
	}

	if n.key == key {
		n.value = val
	} else if n.key > key {
		n.left = put(n.left, key, val)
	} else {
		n.right = put(n.right, key, val)
	}

	if n.right.isRed() && !n.left.isRed() {
		n = rotateLeft(n)
	}
	if n.left.isRed() && n.left.left.isRed() {
		n = rotateRight(n)
	}
	if n.left.isRed() && n.right.isRed() {
		flipColors(n)
	}

	n.size = size(n.left) + 1 + size(n.right)

	return n
}

// Get 查找元素
func (t *RBTree) Get(key string) interface{} {
	return get(t.root, key)
}

func get(n *node, key string) interface{} {
	if n == nil {
		return nil
	} else if n.key == key {
		return n.value
	} else if n.key < key {
		return get(n.right, key)
	} else {
		return get(n.left, key)
	}
}
