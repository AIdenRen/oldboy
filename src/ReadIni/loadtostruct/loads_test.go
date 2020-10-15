package loadtostruct

import "testing"

func TestLoad(t *testing.T) {

	loadIni("", &MySQL{"", "", "", ""})
}
