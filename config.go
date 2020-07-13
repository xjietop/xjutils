package xjutils

import (
	"encoding/json"
	"gitee.com/xjietop/xjutils/entity"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

func GetAppConfig(profile string) AppConfig {
	filename := "./conf/conf.yml"
	if profile != "" {
		filename = "./conf/conf-" + profile + ".yml"
	}
	data, _ := ioutil.ReadFile(filename)
	t := AppConfig{}
	//把yaml形式的字符串解析成struct类型
	err := yaml.Unmarshal(data, &t)
	if err != nil {
		log.Println("配置文件有误：", err)
		os.Exit(1)
	}
	return t
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

func GetNacAppConfig(AppConfig AppConfig, profile string) NacAppConfig {
	log.Println("开始读取Nac配置信息...")
	t := NacAppConfig{}
	filename := AppConfig.App.Name + ".yml"
	if profile != "" {
		filename = AppConfig.App.Name + "-" + profile + ".yml"
	}
	url := "http://" + AppConfig.Register.Url + ":" + AppConfig.Register.Port + "/config/get?DataId=" + filename
	data, err := HttpGetStr(url)
	if err != nil {
		log.Println("读取Nac配置信息出错了")
		log.Println(err)
		os.Exit(1)
	}
	var r entity.R
	err = json.Unmarshal([]byte(data), &r)
	if err != nil {
		log.Println("读取Nac配置信息出错了")
		log.Println(err)
		os.Exit(1)
	}
	if r.Data == nil {
		log.Println("Nac配置信息为空")
		os.Exit(1)
	}
	d := []byte(r.Data.(string))
	//把yaml形式的字符串解析成struct类型
	err = yaml.Unmarshal(d, &t)
	if err != nil {
		log.Println("读取Nac配置信息出错了")
		log.Println(err)
		os.Exit(1)
	}
	log.Println("读取Nac配置信息...OK")
	return t
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
	Redis struct {
		Addr     string
		Password string
		DB       int
	}
}
