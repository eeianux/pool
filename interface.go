package pool

type Pool interface {
	Go(func())
	Close()
}
