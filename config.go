package utils

import (
	"encoding/json"
	"io/ioutil"

	"gitee.com/xjieinfo/xjutils/models"
	"gopkg.in/yaml.v2"
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
	var r models.R
	json.Unmarshal([]byte(data), &r)
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
