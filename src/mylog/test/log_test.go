package test

import (
	"mylog/mylogConsole"
	"os"
	"testing"
)

func TestA(t *testing.T) {
	con := mylogConsole.GetCon(mylogConsole.Debug)
	con.Info("%s", "diaoniamde")
}

func TestB(t *testing.T) {
	open, _ := os.Open("/Users/renchen/IdeaProjects/基础/ssmBuild2/web/WEB-INF/web.xml")
	stat, _ := open.Stat()
	println(stat.Size())
	print(stat.Name())

	//location, _ := time.LoadLocation("Asia/ShangHai")

}
