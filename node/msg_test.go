package node

import (
	"encoding/json"
	"fmt"
	"github.com/lightStarShip/go-starship/network"
	"testing"
)

func TestSetupMsg(t *testing.T) {
	sr := &SetupReq{
		IV:      *network.NewSalt(),
		SubAddr: "SV4a9kD9PdTLihJhD3CgZbTkMSoc528sXG1Tupz2PCwDZ3",
	}

	bts, _ := json.Marshal(sr)
	fmt.Println(string(bts))
}
