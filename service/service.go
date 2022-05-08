package service

import (
	"github.com/lightStarShip/go-server/config"
	"sync"
)

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
		if config.SysConfig.FreeMode {
			_instance = newFreeService()
		} else {
			_instance = newVipService()
		}
	})
	return _instance
}
