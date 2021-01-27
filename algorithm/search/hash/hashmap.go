package search

import "hash/fnv"

const defaultCap uint64 = 1 << 5

// HashMap 拉链法自实现的简单hashmap
type HashMap struct {
	cap, size uint64
	table     []*node
}

type node struct {
	key   string
	value interface{}
	next  *node
}

// NewHashMap 创建hashmap
func NewHashMap() *HashMap {
	return &HashMap{
		cap:   defaultCap,
		table: make([]*node, defaultCap),
	}
}

// Put 插入元素
func (m *HashMap) Put(key string, value interface{}) {
	i := m.getTableIndex(key)
	for n := m.table[i]; n != nil; n = n.next {
		if n.key == key {
			n.value = value
			return
		}
	}
	m.table[i] = &node{
		key:   key,
		value: value,
		next:  m.table[i],
	}
}

// Get 查找元素
func (m *HashMap) Get(key string) interface{} {
	i := m.getTableIndex(key)
	for n := m.table[i]; n != nil; n = n.next {
		if n.key == key {
			return n.value
		}
	}
	return nil
}

// Delete 删除元素
func (m *HashMap) Delete(key string) bool {
	i := m.getTableIndex(key)
	n := m.table[i]
	if n == nil {
		return false
	} else if n.key == key {
		m.table[i] = n.next
		n.next = nil
		return true
	} else {
		for p, q := n, n.next; q != nil; p, q = p.next, q.next {
			if q.key == key {
				p.next = q.next
				q.next = nil
				return true
			}
		}
	}
	return false
}

func (m *HashMap) getTableIndex(key string) uint64 {
	hashValue := hash(key)
	tableIndex := (m.cap - 1) & (hashValue ^ (hashValue >> 16))
	return tableIndex
}

func hash(key string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(key))
	return h.Sum64()
}
