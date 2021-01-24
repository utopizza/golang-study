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

// Delete 删除指定key的元素
func (b *BinarySearchTree) Delete(key string) {
	delete(b.root, key)
}

func delete(n *node, key string) *node {
	if n == nil {
		return nil
	} else if n.key < key {
		n.right = delete(n.right, key)
		return n
	} else if n.key > key {
		n.left = delete(n.left, key)
		return n
	} else {
		// 当前节点为要删除的结点
		if n.left == nil {
			// 若无左孩，用右孩替代本结点
			return n.right
		} else if n.right == nil {
			// 若无右孩，用左孩替代本结点
			return n.left
		} else {
			// 若既有左孩又有右孩，则需用[前驱结点]，或者[后继结点]替代本结点
			// [前驱结点]即为左子树的最大结点
			// [后继结点即]为右子树的最小结点
			x := min(n.right)
			n = deleteMin(n)
			x.left = n.left
			x.right = n.right
			return x
		}
	}
}

// Min 返回最小key以及value
func (b *BinarySearchTree) Min() (string, interface{}) {
	n := min(b.root)
	return n.key, n.value
}

func min(n *node) *node {
	if n == nil {
		return nil
	}
	for n.left != nil {
		n = n.left
	}
	return n
}

func deleteMin(n *node) *node {
	if n == nil {
		return nil
	} else if n.left == nil {
		return n.right
	} else {
		n.left = deleteMin(n.left)
		return n
	}
}

// Max 返回最大key以及value
func (b *BinarySearchTree) Max() (string, interface{}) {
	n := max(b.root)
	return n.key, n.value
}

func max(n *node) *node {
	if n == nil {
		return nil
	}
	for n.right != nil {
		n = n.right
	}
	return n
}

func deleteMax(n *node) *node {
	if n == nil {
		return nil
	} else if n.right == nil {
		return n.left
	} else {
		n.right = deleteMax(n.right)
		return n
	}
}

// Keys 获取全部key，按大小顺序
func (b *BinarySearchTree) Keys() []string {
	var keys []string
	inorder(b.root, keys)
	return keys
}

func inorder(n *node, keys []string) {
	if n == nil {
		return
	}
	// left tree
	if n.left != nil {
		inorder(n.left, keys)
	}
	// current node
	keys = append(keys, n.key)
	// right tree
	if n.right != nil {
		inorder(n.right, keys)
	}
}
