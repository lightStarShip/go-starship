package utils

import (
	"os/user"
	"path/filepath"
)

const (
	BaseDirName = ".starship"
	PidFileName = "pid"
)

var (
	BasDir = BaseDirName
)

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
