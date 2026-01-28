package main

import (
	"flag"
	"signaling/src/framework"

	"github.com/golang/glog"
)

func main() {
	flag.Parse()

	err := framework.Init()
	if err != nil {
		panic(err)
	}

	glog.SetLogToStderr(true)
	glog.Info("hello go go go")

	err = framework.StartHttp()
	if err != nil {
		panic(err)
	}
}
