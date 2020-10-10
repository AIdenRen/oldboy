package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestTime(t *testing.T){
	now := time.Now()
	t.Log(now)
	t.Log(now.UnixNano())
	rand.Seed(now.UnixNano())
	intn := rand.Intn(100)
	t.Log(intn)
}

func TestSub(t *testing.T){
	now := time.Now()
	parse, _ := time.Parse("2006-01-02 15:04", "2020-10-11 12:00")
	t.Log(parse)
	duration := now.Sub(parse)
	fmt.Println(duration)
}

func TestLocation(t *testing.T){
	location, _ := time.LoadLocation("Asia/ShangHai")
	inLocation, _ := time.ParseInLocation("2006-01-02 15:04", "2020-10-11 11:31", location)
	now := time.Now()
	sub := now.Sub(inLocation)
	t.Log(sub)
}

