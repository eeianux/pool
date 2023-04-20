package pool

import (
	. "github.com/smartystreets/goconvey/convey"
	"sync/atomic"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	Convey("pool", t, func() {
		const cnt = 1000
		pool := NewPool(cnt / 2)
		start := time.Now()
		for i := 0; i < cnt; i++ {
			pool.Go(func() {
				time.Sleep(time.Second)
			})
		}
		pool.Close()
		So(time.Since(start)/time.Millisecond-time.Second*2/time.Millisecond, ShouldBeLessThan, 10)
	})
	Convey("panic", t, func() {
		const cnt = 10
		pool := NewPool(cnt / 2)
		start := time.Now()
		for i := 0; i < cnt; i++ {
			pool.Go(func() {
				time.Sleep(time.Second)
				panic("test")
			})
		}
		pool.Close()
		So(time.Since(start)/time.Millisecond-time.Second*2/time.Millisecond, ShouldBeLessThan, 10)
	})
	Convey("wait", t, func() {
		const cnt = 10
		tt := atomic.Int64{}
		tt.Store(0)
		pool := NewPool(cnt)
		for i := 0; i < cnt; i++ {
			pool.Go(func() {
				tt.Add(1)
			})
		}
		pool.Close()
		So(tt.Load(), ShouldEqual, cnt)
	})
}
