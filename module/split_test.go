package module

import (
	"strings"
	"testing"
)

func TestHandleServer(t *testing.T) {
	s:= "ls"
	ss := strings.Split(s," ")
	st :=  ss[0]=="ls"
	t.Log(ss[0],len(ss),st)//go test -v split_test.go
}
