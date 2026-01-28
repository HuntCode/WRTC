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

	glog.Info("hello go")
	glog.Flush()

	err = framework.StartHttp()
	if err != nil {
		panic(err)
	}
}
