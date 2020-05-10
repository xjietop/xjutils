package xjutils

import (
	"log"
	"os"
	"time"
)

func LogOutput(str,prefix string,calldepin int){
	file := "./logs/" + time.Now().Format("2006010215") + ".txt"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer logFile.Close()
	if nil != err {
		panic(err)
	}
	loger := log.New(logFile, prefix, log.Ldate | log.Lmicroseconds | log.Llongfile)
	loger.Output(calldepin,str)
}
func LogDebug(str string){
	LogOutput(str,"[DEBUG]",3)
}
func LogInfo(str string){
	LogOutput(str,"[INFO]",3)
}
func LogWarn(str string){
	LogOutput(str,"[WARN]",3)
}
func LogError(str string){
	LogOutput(str,"[ERROR]",3)
}
func LogPrefix(str,prefix string){
	LogOutput(str,prefix,3)
}

