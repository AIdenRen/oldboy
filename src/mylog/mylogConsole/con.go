package mylogConsole

import (
	"fmt"
	"runtime"
	"time"
)

type logLevel  uint8

const(
	Debug logLevel = iota
	Info
	Warning
	Error
	Fatal
)

type Con struct {
	lever logLevel
}

func GetCon(level logLevel) *Con{
	var con *Con = new(Con)
	con.lever = level
	return con
}

func (con *Con) Debug(str string,args...interface{}){
	con.switch1(Debug,str,args)
}

func (con *Con) Info(str string,args...interface{}){
	con.switch1(Info,str,args)
}

func (con *Con) Warning(str string,args...interface{}){
	con.switch1(Warning,str,args)
}

func (con *Con) Error(str string,args...interface{}){
	con.switch1(Error,str,args)
}

func (con *Con) Fatal(str string,args...interface{}){
	con.switch1(Fatal,str,args)
}


func timeUtil() string{
	now := time.Now()
	return now.Format("2006-01-02 15:04:05")
}

func (con *Con) switch1(lev logLevel,str string,args...interface{}){
	switch lev>=con.lever {
	case true :
		fmt.Print(Fatal,"\t")
		localtion, s, i := callerLocaltion(2)
		fmt.Printf("[%s] [%s] [%d] ",localtion,s,i)
		fmt.Printf("[%s]",timeUtil())
		fmt.Printf(str,args...)
	case false:
		fmt.Println("级别错误")
	}
}

func callerLocaltion(skip int) (string,string,int){//file method.name line
	caller, file, line, ok := runtime.Caller(skip)
	if !ok{
		fmt.Println("未能获取信息")
	}
	pc := runtime.FuncForPC(caller)
	name := pc.Name()

	return file,name,line

}


