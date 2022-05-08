package config

import (
	"fmt"
	"github.com/lightStarShip/go-server/config/fdlimit"
	"math/rand"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"time"
)

const (
	BaseDirName = ".starship"
	PidFileName = "pid"
	FreeSrvPort = 30001
	VipSrvPort  = 40001
)

var (
	BasDir = BaseDirName
)

type StartParam struct {
	Version  bool
	FreeMode bool
	LogLevel int8
}

var SysConfig = &StartParam{}

func init() {
	BasDir = BaseUsrDir()
}

func PidFilePath() string {
	return filepath.Join(BasDir, string(filepath.Separator), PidFileName)
}

func BaseUsrDir() string {

	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	baseDir := filepath.Join(usr.HomeDir, string(filepath.Separator), BaseDirName)
	return baseDir
}

func SetupFDLimit() error {
	if err := os.Setenv("GODEBUG", "netdns=go"); err != nil {
		return err
	}

	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(int64(time.Now().Nanosecond()))
	limit, err := fdlimit.Maximum()
	if err != nil {
		return fmt.Errorf("failed to retrieve file descriptor allowance:%s", err)
	}
	_, err = fdlimit.Raise(uint64(limit))
	if err != nil {
		return fmt.Errorf("failed to raise file descriptor allowance:%s", err)
	}
	return nil
}
