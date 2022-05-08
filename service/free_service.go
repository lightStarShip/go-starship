package service

type FreeService struct {
}

func newFreeService() *FreeService {
	fs := &FreeService{}
	return fs
}

func (fs *FreeService) StartUp() error {

	return nil
}

func (fs *FreeService) ShutDown() {

}
