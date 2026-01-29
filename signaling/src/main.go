package main

import (
	"flag"
	"signaling/src/framework"
)

func main() {
	flag.Parse()

	err := framework.Init("./conf/framework.conf")
	if err != nil {
		panic(err)
	}

	framework.RegisterStaticUrl()

	go startHttp()

	startHttps()
}

func startHttp() {
	err := framework.StartHttp()
	if err != nil {
		panic(err)
	}
}

func startHttps() {
	err := framework.StartHttps()
	if err != nil {
		panic(err)
	}
}
