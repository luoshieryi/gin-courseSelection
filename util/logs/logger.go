package logs

import (
	"log"
	"runtime"
)

/*
 @Author: as
 @Date: Creat in 15:27 2022/2/12
 @Description: 简单的日志封装打印
*/

type Level string

var (
	DB      Level = "DB"
	Service Level = "Service"
	API     Level = "Api"
)

const (
	green   = "\033[97;42m"
	white   = "\033[90;47m"
	yellow  = "\033[90;43m"
	red     = "\033[97;41m"
	blue    = "\033[97;44m"
	magenta = "\033[97;45m"
	cyan    = "\033[97;46m"
	reset   = "\033[0m"
)

func PrintLogErr(level Level, msg string, err interface{}) {
	pc, _, _, _ := runtime.Caller(1)
	log.Printf("%s[%s]%s | [%s] %s : %v\n",
		red, level, reset,
		runtime.FuncForPC(pc).Name(),
		msg,
		err,
	)
}

func PrintLog(msg ...interface{}) {
	log.Println(msg...)
}
