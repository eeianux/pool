# eeianux/pool
[![LICENSE](https://img.shields.io/github/license/eeianux/pool.svg)](https://github.com/eeianux/pool/blob/main/LICENSE)
[![Language](https://img.shields.io/badge/Language-Go-blue.svg)](https://golang.org/)
[![Go Report Card](https://goreportcard.com/badge/github.com/eeianux/pool)](https://goreportcard.com/report/github.com/eeianux/pool)
[![Coverage](https://codecov.io/gh/eeianux/pool/branch/main/graph/badge.svg)](https://codecov.io/gh/eeianux/pool)

 a simple go routine pool

# Installation
```shell
go get github.com/eeianux/pool
go mod tidy
```

# Quick start
Make a test, for example:

```go
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

```
