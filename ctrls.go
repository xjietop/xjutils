package xjutils

import (
	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context){
	c.String(200,"ok")
}

func Refresh(c *gin.Context){
	//models.Loadconf()
	local_url := Config.Section("local").Key("url").String()
	local_port := Config.Section("local").Key("port").String()
	register_url := Config.Section("register").Key("url").String()
	register_port := Config.Section("register").Key("port").String()
	local_appname := Config.Section("local").Key("appname").String()
	registerurl := "http://"+register_url + ":" + register_port +"/put/" + local_appname+ "/"+local_url+":"+local_port
	HttpGetStr(registerurl)
	c.String(200,"refresh ok.")
}
