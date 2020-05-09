package helper

import (
	"log"
	"os"
	"time"
)

func LogInfo() *log.Logger {
	//创建io对象，日志的格式为当前时间.log;
	file := "./" + time.Now().Format("2006-01-02") + ".log"

	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if nil != err {
		panic(err)
	}

	//创建一个Logger：参数1：日志写入目的地 ； 参数2：每条日志的前缀 ；参数3：日志属性
	return  log.New(logFile, "Service:",log.Ldate| log.Ltime )
}

