package utils

import (
	"fmt"
	"github.com/lightStarShip/go-server/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
	"os"
	"time"
)

var _logInst *zerolog.Logger

func InitLog() {

	writer := diode.NewWriter(os.Stderr, 1000, 10*time.Millisecond, func(missed int) {
		fmt.Printf("Logger Dropped %d messages", missed)
	})
	out := zerolog.ConsoleWriter{Out: writer}
	out.TimeFormat = time.StampMilli

	logger := zerolog.New(out).
		Level(zerolog.Level(config.SysConfig.LogLevel)).
		With().
		Caller().
		Timestamp().
		Logger()
	_logInst = &logger
}
