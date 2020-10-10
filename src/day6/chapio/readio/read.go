package readio

import (
	"fmt"
	"os"
	"testing"
)

type A struct {
	Name string `json:name`
}

type B interface {
	run()
}

func (ac *A)run(){
	fmt.Println(ac.Name)
}

func ReaTest(t *testing.T){
	os.Open("")
	var AV A
	AV.Name = "wodiaoniamde"

	var an B = &AV
	an.run()

}