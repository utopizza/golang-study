package main

import "sync"

// Memo 缓存
type Memo struct {
	f     Func
	mux   sync.Mutex
	cache map[string]*entry
}

// Func 需要被缓存的函数
type Func func(key string) (interface{}, error)

// Result 函数的执行结果
type Result struct {
	value interface{}
	err   error
}

type entry struct {
	res   Result
	ready chan struct{}
}

// NewMemo 创建缓存
func NewMemo(f Func) *Memo {
	return &Memo{
		f:     f,
		cache: make(map[string]*entry),
	}
}

// Get 查询缓存
func (m *Memo) Get(key string) (interface{}, error) {
	m.mux.Lock()
	e := m.cache[key]
	if e == nil {
		// e为nil说明当前go程是第一次请求
		// 先构造一个指针写回缓存，释放锁，再执行耗时的f函数
		e = &entry{ready: make(chan struct{})}
		m.cache[key] = e
		m.mux.Unlock()

		// 执行f
		e.res.value, e.res.err = m.f(key)

		// 广播ready信息
		close(e.ready)
	} else {
		// 重复请求
		m.mux.Unlock()
		<-e.ready // 等待这个数据ready
	}

	return e.res.value, e.res.err
}
