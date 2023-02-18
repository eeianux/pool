package pool

import "runtime/debug"

type pool struct {
	pool chan func()
}

func (r pool) handle() {
	defer func() {
		if e := recover(); e != nil {
			// 避免所有协程退出导致通道阻塞，引起死锁
			go r.handle()
			debug.PrintStack()
		}
	}()
	for {
		if f, ok := <-r.pool; ok {
			f()
		} else {
			break
		}
	}
}

func NewPool(cnt int) Pool {
	pool := pool{
		pool: make(chan func(), cnt),
	}
	for i := 0; i < cnt; i++ {
		go pool.handle()
	}
	return pool
}

func (r pool) Go(f func()) {
	r.pool <- f
}

func (r pool) Close() {
	close(r.pool)
}
