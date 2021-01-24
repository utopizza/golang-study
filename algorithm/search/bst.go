package search

// BinarySearchTree 二叉搜索树
type BinarySearchTree struct {
	root *node
}

type node struct {
	left, right *node
	key         string
	value       interface{}
}

// NewBinarySearchTree 创建二叉搜索树
func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

// Put 插入元素
func (b *BinarySearchTree) Put(key string, value interface{}) {
	b.root = put(b.root, key, value)
}

func put(n *node, key string, value interface{}) *node {
	if n == nil {
		return &node{key: key, value: value}
	} else if n.key == key {
		n.value = value
		return n
	} else if n.key < key {
		n.right = put(n.right, key, value)
		return n
	} else {
		n.left = put(n.left, key, value)
		return n
	}
}

// Get 查找元素
func (b *BinarySearchTree) Get(key string) interface{} {
	return get(b.root, key)
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
