package utils

import (
	"log"
	"os"
)

var logger *log.Logger

func GetLog(pre string) *log.Logger {
	if nil == logger {
		logger = log.New(os.Stdout, pre+"|", 0)
	}
	return logger
}
