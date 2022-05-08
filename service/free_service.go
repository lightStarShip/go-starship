package service

import (
	"fmt"
	"github.com/lightStarShip/go-server/config"
	"net"
)

type FreeService struct {
	l net.Listener
}

func newFreeService() *FreeService {
	addr := fmt.Sprintf(":%d", config.FreeSrvPort)
	ls, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	fs := &FreeService{l: ls}
	return fs
}

func (fs *FreeService) StartUp() error {
	go fs.accepting()
	return nil
}

func (fs *FreeService) ShutDown() {

}
func (fs *FreeService) accepting() {

}
