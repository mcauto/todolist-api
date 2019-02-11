package logger

import (
	"fmt"
	"todo-list-back/config"

	"github.com/amoghe/distillog"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var (
	// Std Standard File logger
	Std distillog.Logger
	// Err Error file logger
	Err distillog.Logger
)

func init() {
	// stdFileName standard file log path
	stdFileName := fmt.Sprintf("/tmp/log/%s/%s.log", config.Conf.Name, config.Conf.Name)
	// errFileName error file log path
	errFileName := fmt.Sprintf("/tmp/log/%s/%s_error.log", config.Conf.Name, config.Conf.Name)

	stdLumberjackHandle := &lumberjack.Logger{
		Filename:   stdFileName,
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
	}

	errLumberjackHandle := &lumberjack.Logger{
		Filename:   errFileName,
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
	}
	Std = distillog.NewStreamLogger("", stdLumberjackHandle)
	Err = distillog.NewStreamLogger("", errLumberjackHandle)
}
