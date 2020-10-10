package test

import (

	"mylog/mylogConsole"

	"testing"
)

func TestA(t *testing.T){
	con := mylogConsole.GetCon(mylogConsole.Debug)
	con.Info("%s","diaoniamde")
}