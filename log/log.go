package log

import (
	"fmt"
	"log"
	"os"
)

const (
	Prefix = "Spider Log: "
)

var (
	// 全局logger
	logger *log.Logger
)

func NewLog() {
	l, err := os.Create("/opt/log/spider.log")
	if err != nil {
		fmt.Errorf("os.Create is error, %v", err)
		return
	}

	logger = log.New(l, Prefix, log.Lshortfile|log.LstdFlags)
}

func Info(v ...interface{}) {
	logger.Print(append([]interface{}{"Info: "}, v...))
}

func Infof(format string, v ...interface{}) {
	logger.Printf("Info: "+format, v...)
}

func Warn(v ...interface{}) {
	logger.Print(append([]interface{}{"Warn: "}, v...))
}

func Warnf(format string, v ...interface{}) {
	logger.Printf("Warn: "+format, v...)
}

func Error(v ...interface{}) {
	logger.Print(append([]interface{}{"ERROR: "}, v...))
}

func Errorf(format string, v ...interface{}) {
	logger.Printf("ERROR: "+format, v...)
}
