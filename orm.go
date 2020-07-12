package xjutils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	//"github.com/go-xorm/xorm"
	"github.com/xormplus/xorm"
)

var (
	Db *xorm.Engine
)

func OrmInit(Conf NacAppConfig) *xorm.Engine {
	dataSourceName := Conf.Datasource.Username + ":" + Conf.Datasource.Password + "@tcp(" + Conf.Datasource.Url + ":" + Conf.Datasource.Port + ")/" + Conf.Datasource.Database + "?charset=utf8&parseTime=True&loc=Asia%2fShanghai"
	Db, err := xorm.NewEngine(Conf.Datasource.Drivername, dataSourceName)
	if err = Db.Ping(); err != nil {
		fmt.Println("数据库连接失败")
		os.Exit(1)
	}
	Db.ShowSQL(true)
	//lst := make([]map[string]string,0)
	//sql := "select * from m_user limit ?,?"
	//err = Db.SQL(sql,0,10).Find(&lst)
	//if err!=nil{
	//	fmt.Println(err)
	//}
	//fmt.Println(lst)
	return Db
}
func GormInit(Conf NacAppConfig) *gorm.DB {
	dataSourceName := Conf.Datasource.Username + ":" + Conf.Datasource.Password + "@tcp(" + Conf.Datasource.Url + ":" + Conf.Datasource.Port + ")/" + Conf.Datasource.Database + "?charset=utf8&parseTime=True&loc=Asia%2fShanghai"
	log.Println("开始连接数据库...")
	Gorm, err := gorm.Open(Conf.Datasource.Drivername, dataSourceName)
	if err != nil {
		log.Println("数据库连接失败")
		log.Println(err)
		os.Exit(1)
	}
	Gorm.SingularTable(true)
	Gorm.LogMode(true)
	log.Println("连接数据库...OK")
	return Gorm
}
func Init() {
	Loadconf()
	var err error
	register_url := Config.Section("register").Key("url").String()
	register_port := Config.Section("register").Key("port").String()
	//appname := Config.Section("local").Key("appname").String()
	dbfname := Config.Section("local").Key("dbfname").String()
	//regurl := "http://"+register_url+":"+register_port+"/item/"+appname
	//HttpGetStr(regurl)
	confurl := "http://" + register_url + ":" + register_port + "/item/config"
	confstr := HttpGetStr(confurl)
	var confitem []Srv
	json.Unmarshal([]byte(confstr), &confitem)
	url0 := confitem[0].Url
	surl := "http://" + url0 + "/config/dbf/" + dbfname
	println(surl)
	dataSourceName := HttpGetStr(surl)
	if dataSourceName == "" {
		fmt.Println(surl + "获取配置失败")
		os.Exit(1)
	}
	Db, err = xorm.NewEngine("mysql", dataSourceName)
	if err = Db.Ping(); err != nil {
		fmt.Println("数据库连接失败")
		os.Exit(1)
	}
	Db.ShowSQL(true)
}

type Srv struct {
	Name      string
	Url       string
	StartTime string
}
