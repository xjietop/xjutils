package xjutils

import (
	"encoding/json"
	"gitee.com/xjieinfo/xjutils/entity"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

func GetAppConfig(profile string) (error, AppConfig) {
	filename := "./conf/conf.yml"
	if profile != "" {
		filename = "./conf/conf-" + profile + ".yml"
	}
	data, _ := ioutil.ReadFile(filename)
	t := AppConfig{}
	//把yaml形式的字符串解析成struct类型
	err := yaml.Unmarshal(data, &t)
	return err, t
}

type AppConfig struct {
	App struct {
		Name string
		Url  string
		Port string
	}
	Register struct {
		Url  string
		Port string
	}
}

func GetNacAppConfig(AppConfig AppConfig, profile string) (error, NacAppConfig) {
	filename := AppConfig.App.Name + ".yml"
	if profile != "" {
		filename = AppConfig.App.Name + "-" + profile + ".yml"
	}
	url := "http://" + AppConfig.Register.Url + ":" + AppConfig.Register.Port + "/config/get?DataId=" + filename
	data := HttpGetStr(url)
	t := NacAppConfig{}
	var r entity.R
	json.Unmarshal([]byte(data), &r)
	if r.Data == nil {
		log.Println("nac配置文件读取失败:", url)
		os.Exit(1)
	}
	d := []byte(r.Data.(string))
	//把yaml形式的字符串解析成struct类型
	err := yaml.Unmarshal(d, &t)
	return err, t
}

type NacAppConfig struct {
	Etcd struct {
		Url       []string
		LeaseTime int64
	}
	Datasource struct {
		Drivername string
		Url        string
		Port       string
		Username   string
		Password   string
		Database   string
	}
}
