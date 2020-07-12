package xjutils

import (
	"fmt"
	"os"
	"time"

	"github.com/go-ini/ini"
)

var Config *ini.File

func Loadconf() {
	var err error
	Config, err = ini.Load("./conf.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
}

func Register() {
	local_url := Config.Section("local").Key("url").String()
	local_port := Config.Section("local").Key("port").String()
	register_url := Config.Section("register").Key("url").String()
	register_port := Config.Section("register").Key("port").String()
	local_appname := Config.Section("local").Key("appname").String()
	registerurl := "http://" + register_url + ":" + register_port + "/put/" + local_appname + "/" + local_url + ":" + local_port
	HttpGetStr(registerurl)

	ser := NewServiceReg([]string{"127.0.0.1:2379"}, 5)
	ser.PutService(local_appname+"/"+local_url+":"+local_port, "")
}

type RegItem struct {
	Url        string
	Status     string
	Createtime time.Time
	Updatetime time.Time
}
