package main

import "sync"

// Memo 缓存
type Memo struct {
	f     Func
	mux   sync.Mutex
	cache map[string]Result
}

// Func 需要被缓存的函数
type Func func(key string) (interface{}, error)

// Result 函数的执行结果
type Result struct {
	value interface{}
	err   error
}

// NewMemo 创建缓存
func NewMemo(f Func) *Memo {
	return &Memo{
		f:     f,
		cache: make(map[string]Result),
	}
}

// Get 查询缓存
func (m *Memo) Get(key string) (interface{}, error) {

	// 临界区1：读
	m.mux.Lock()
	res, ok := m.cache[key]
	m.mux.Unlock()

	// 这样实现的问题是，可能有多个go程先后发现同一个url不在缓存中，然后重复更新缓存
	if !ok {
		res.value, res.err = m.f(key)

		// 临界区2：写
		m.mux.Lock()
		m.cache[key] = res
		m.mux.Unlock()
	}
	return res.value, res.err
}
