package utils

import "fmt"

const (
	MaxVer   = 1
	MiniVer  = 0
	BuildVer = 0
)

func VerStr() string {
	return fmt.Sprintf("%d.%d.%d", MaxVer, MiniVer, BuildVer)
}
