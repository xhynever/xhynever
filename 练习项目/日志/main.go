package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

var Logger *log.Entry

func SetLogger() {
	Logger = log.WithFields(log.Fields{
		"sysId":  "sysid",
		"module": "module",
	})
}

func (Logger *log.Entry) Info(args ...interface{}) {
	fmt.Println("Info")
	Logger.WithFields(log.Fields{"globalSeqNo": globalSeqNo}).Infof(format, args...)
}

// func Infof(globalSeqNo, format string, args ...interface{}) {

//
// }

func main() {
	SetLogger()
	Logger.Infof("Starting")

}
