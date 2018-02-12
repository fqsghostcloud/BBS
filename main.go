package main

import (
	_ "bbs/routers"
	"flag"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/golang/glog"
)

func main() {
	// use glog
	flag.Parse()
	defer glog.Flush()

	orm.Debug = true                                 // database debug model
	beego.BConfig.WebConfig.Session.SessionOn = true // session on
	beego.Run()
	glog.Flush()
}
