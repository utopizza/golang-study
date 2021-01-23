package main

// Memo 缓存
type Memo struct {
	requestChan chan request
}

// Func 需要被缓存的函数
type Func func(key string) (interface{}, error)

// Result 函数的执行结果
type Result struct {
	value interface{}
	err   error
}

// request 一次请求，包含一个返回通道
type request struct {
	key          string
	responseChan chan Result
}

type entry struct {
	res   Result
	ready chan struct{}
}

// NewMemo 创建缓存
func NewMemo(f Func) *Memo {
	m := &Memo{
		requestChan: make(chan request),
	}
	go m.Server(f)
	return m
}

// Get 查询缓存
func (m *Memo) Get(key string) (interface{}, error) {
	// 创建一个请求，请求中包含一个返回通道response
	r := request{
		key:          key,
		responseChan: make(chan Result),
	}

	// 通过通道，发送请求到 m.Server
	m.requestChan <- r

	// 等待 m.Server 将结果通过返回通道推送回来
	res := <-r.responseChan

	return res.value, res.err
}

// Server 在此方法中维护缓存，只通过channel对外返回数据
func (m *Memo) Server(f Func) {
	// 缓存
	cache := make(map[string]*entry)

	// 监听通道中的请求
	for req := range m.requestChan {
		e := cache[req.key]
		if e == nil {
			e = &entry{
				ready: make(chan struct{}),
			}
			cache[req.key] = e
			go e.call(f, req.key) // 开go程执行f
		}
		go e.deliver(req.responseChan) // 开go程返回数据
	}
}

func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	close(e.ready)
}

func (e *entry) deliver(response chan<- Result) {
	<-e.ready // 等待数据ready
	response <- e.res
}
