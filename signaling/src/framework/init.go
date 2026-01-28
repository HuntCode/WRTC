package framework

import (
	"github.com/golang/glog"
)

func Init() error {
	glog.SetLogDir("./log")
	glog.SetLogFilename("signaling")
	glog.SetLogToStderr(true)
	return nil
}
