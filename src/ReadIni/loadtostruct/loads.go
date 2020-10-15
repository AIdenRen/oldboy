package loadtostruct

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

type MySQL struct {
	Url      string `ini:url`
	Driver   string `ini:driver`
	User     string `ini:user`
	Password string `ini:password`
}

type Redis struct {
	Url      string `ini:url`
	Driver   string `ini:driver`
	User     string `ini:user`
	Password string `ini:password`
}

func loadIni(fileName string, data interface{}) error {
	//调用isStruct 传入值是否正常
	if !isStruct(data) {
		return fmt.Errorf("%v", "需要传入结构体指针")
	}
	//读取文件
	file, err := ioutil.ReadFile("/Users/renchen/oldboy/src/ReadIni/configuration.ini")
	if err != nil {
		fmt.Errorf("%v", "打开文件错误")
	}
	configurationInfo := string(file)
	splitSlice := strings.Split(configurationInfo, "\n")
	notes := trimNotes(splitSlice)
	info := findInfo(notes, findSection(findSec(notes), "mysql"))
	fmt.Println(info)
	return nil
}

//判断传入的值是否符合指针结构体
func isStruct(data interface{}) bool {
	//判断是否为指针
	t := reflect.TypeOf(data)
	if t.Kind() != reflect.Ptr {
		return false
	}
	//判断是否为结构体
	if t.Elem().Kind() != reflect.Struct {
		return false
	}

	return true
}

//去除读出内容中的注释
func trimNotes(str []string) []string {
	resStr := make([]string, 0, 10)
	for _, value := range str {
		value = strings.TrimSpace(value)
		if strings.HasPrefix(value, "#") || strings.HasPrefix(value, ";") {
			continue
		}
		resStr = append(resStr, value)
	}
	return resStr
}

//是否传入结
func findSection(str []string, sec string) string {
	var secstr string
	for _, value := range findSec(str) {
		if strings.TrimRight(strings.TrimLeft(sec, "["), "]") == value {
			secstr = value
		}
	}
	return secstr
}

//找到结
func findSec(str []string) []string {
	resSli := make([]string, 0, 10)
	for _, value := range str {
		if strings.HasPrefix(value, "[") && strings.HasSuffix(value, "]") {
			value = strings.TrimRight(strings.TrimLeft(value, "["), "]")
			resSli = append(resSli, value)
		}
	}
	return resSli
}

//找到结相关信息
func findInfo(strs []string, str string) []string {
	res := make([]string, 0, 10)
	for _, value := range strs {
		if strings.TrimRight(strings.TrimLeft(value, "["), "]") == str {
			res = append(res, value)
		}
		if strings.HasPrefix(value, "[") && strings.HasSuffix(value, "]") {
			return res
		}
	}
	return res
}
