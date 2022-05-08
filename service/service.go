package service

import "sync"

type Proxy interface {
	StartUp() error
	ShutDown()
}

var (
	_instance Proxy = nil
	once      sync.Once
)

func Inst() Proxy {
	once.Do(func() {
		_instance = newFreeService()
	})
	return _instance
}
