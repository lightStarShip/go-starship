package service

type VipService struct {
}

func newVipService() *VipService {
	vs := &VipService{}
	return vs
}

func (vs *VipService) StartUp() error {

	return nil
}

func (vs *VipService) ShutDown() {

}
