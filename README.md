# rool, a simple go routine pool

# Installation
```shell
go get github.com/eeianux/pool
go mod tidy
```

# Quick start
Make a test, for example:

```go
package rool

import (
	. "github.com/smartystreets/goconvey/convey"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	Convey("pool", t, func() {
		const cnt = 1000
		pool := NewPool(cnt / 2)
		defer pool.Close()
		start := time.Now()
		wg := sync.WaitGroup{}
		for i := 0; i < cnt; i++ {
			wg.Add(1)
			pool.Go(func() {
				defer wg.Done()
				time.Sleep(time.Second)
			})
		}
		wg.Wait()
		So(time.Since(start)/time.Millisecond-time.Second*2/time.Millisecond, ShouldBeLessThan, 10)
	})
	Convey("panic", t, func() {
		const cnt = 10
		pool := NewPool(cnt / 2)
		defer pool.Close()
		start := time.Now()
		wg := sync.WaitGroup{}
		for i := 0; i < cnt; i++ {
			wg.Add(1)
			pool.Go(func() {
				defer wg.Done()
				time.Sleep(time.Second)
				panic("test")
			})
		}
		wg.Wait()
		So(time.Since(start)/time.Millisecond-time.Second*2/time.Millisecond, ShouldBeLessThan, 10)
	})
}

```