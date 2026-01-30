package framework

import (
	"fmt"

	"github.com/golang/glog"
)

var gconf *FrameworkConf

func Init(confFile string) error {
	var err error
	gconf, err = loadConf(confFile)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", gconf)

	glog.SetLogDir(gconf.logDir)
	glog.SetLogFilename(gconf.logFile)
	glog.SetLogToStderr(gconf.logToStderr)

	err = loadWrpc()
	if err != nil {
		return err
	}

	return nil
}
