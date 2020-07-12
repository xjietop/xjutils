package xjutils

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	//"github.com/go-xorm/xorm"
	"github.com/xormplus/xorm"
)

var (
	Db      *xorm.Engine
	Gorm    *gorm.DB
	Redisdb *redis.Client
)

func XormInit(Conf NacAppConfig) *xorm.Engine {
	dataSourceName := Conf.Datasource.Username + ":" + Conf.Datasource.Password + "@tcp(" + Conf.Datasource.Url + ":" + Conf.Datasource.Port + ")/" + Conf.Datasource.Database + "?charset=utf8&parseTime=True&loc=Asia%2fShanghai"
	Db, err := xorm.NewEngine(Conf.Datasource.Drivername, dataSourceName)
	if err = Db.Ping(); err != nil {
		fmt.Println("数据库连接失败")
		os.Exit(1)
	}
	Db.ShowSQL(true)
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
	confstr, _ := HttpGetStr(confurl)
	var confitem []Srv
	json.Unmarshal([]byte(confstr), &confitem)
	url0 := confitem[0].Url
	surl := "http://" + url0 + "/config/dbf/" + dbfname
	println(surl)
	dataSourceName, _ := HttpGetStr(surl)
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

func Count(db *gorm.DB, sql string) (total int64) {
	count := struct{ Count int64 }{}
	err := db.Raw(sql).First(&count).Error
	if err != nil {
		log.Println(err)
		return
	}
	total = count.Count
	return
}

func RedisInit(Conf NacAppConfig) *redis.Client {
	Redisdb = redis.NewClient(&redis.Options{
		Addr:     Conf.Redis.Addr,     // use default Addr
		Password: Conf.Redis.Password, // no password set
		DB:       Conf.Redis.DB,       // use default DB
	})
	//心跳
	pong, err := Redisdb.Ping().Result()
	if err != nil {
		log.Println("连接redis出错了")
		log.Println(err)
		os.Exit(0)
	}
	log.Println("连接redis...", pong, err) // Output: PONG <nil>
	return Redisdb
}
