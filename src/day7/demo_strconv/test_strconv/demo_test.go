package test_strconv

import (
	"strconv"
	"testing"
)

func TestA(t *testing.T) {
	str := "1000"
	parseInt, _ := strconv.ParseInt(str, 10, 64)
	t.Logf("%T,%v", parseInt, parseInt)
}
