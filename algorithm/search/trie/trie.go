package search

// Trie 字典树
type Trie struct {
	root *node
}

type node struct {
	val    interface{}
	childs [256]*node // 256长度对应256个ascii符号
}

// NewTrie 创建字典树
func NewTrie() *Trie {
	return &Trie{}
}

// Put 插入元素
func (t *Trie) Put(key string, value interface{}) {
	t.root = put(t.root, []rune(key), value, 0)
}

func put(n *node, key []rune, value interface{}, i int) *node {
	if n == nil {
		n = &node{}
	}
	// 所有字符都已经挂载后，将val挂载到叶子结点
	if i == len(key) {
		n.val = value
		return n
	}
	// 否则挂载当前字符，剩余字符继续往下挂载
	c := key[i]
	n.childs[c] = put(n.childs[c], key, value, i+1)
	return n
}

// Get 查找元素
func (t *Trie) Get(key string) interface{} {
	n := get(t.root, []rune(key), 0)
	if n != nil {
		return n.val
	}
	return nil
}

func get(n *node, key []rune, i int) *node {
	if n == nil {
		return nil
	}
	if i == len(key) {
		return n
	}
	c := key[i]
	return get(n.childs[c], key, i+1)
}

// Keys 获取所有key
func (t *Trie) Keys() []string {
	return t.KeysWithPrefix("")
}

// KeysWithPrefix 获取指定前缀的所有key
func (t *Trie) KeysWithPrefix(keyPrefix string) []string {
	// 先找到前缀对应结点
	n := get(t.root, []rune(keyPrefix), 0)
	// 遍历收集该结点下所有val
	var keys []string
	collectKeys(n, keyPrefix, &keys)
	return keys
}

func collectKeys(n *node, keyPrefix string, keys *[]string) {
	if n == nil {
		return
	}
	if n.val != nil {
		*keys = append(*keys, keyPrefix)
	}
	for i := range n.childs {
		if n.childs[i] != nil {
			str := string([]rune{rune(i)})
			collectKeys(n.childs[i], keyPrefix+str, keys)
		}
	}
}

// Delete 删除元素
func (t *Trie) Delete(key string) {
	t.root = delete(t.root, []rune(key), 0)
}

func delete(n *node, key []rune, i int) *node {
	if n == nil {
		return nil
	}

	// 递归找到key并删除val
	if i == len(key) {
		n.val = nil
	} else {
		c := key[i]
		n.childs[c] = delete(n.childs[c], key, i+1)
	}

	// 判断是否可以释放当前结点：当前结点是否没有val&&没有孩子，是则可以删除
	if n.val != nil {
		return n
	}
	for i := range n.childs {
		if n.childs[i] != nil {
			return n
		}
	}
	return nil
}
