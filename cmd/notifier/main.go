package main

import (
	"flag"
	"github.com/golang/glog"

	"github.com/codeflavor/notifier/pkg/controller"
)

func main() {
	defer glog.Flush()
	flag.Set("logtostderr", "true")
	flag.Parse()

	controller := controller.Controller{}
	controller.Load()
	controller.Start()
}
