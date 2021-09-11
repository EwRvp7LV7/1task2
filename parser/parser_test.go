package parser

import (
	"fmt"
	"testing"
)


type testpair struct {
    raw string
    ready []string
}

var tests = []testpair{
    { "ggg!hhh?ddd",  []string{"ggg","hhh","ddd"}},
    { "жжж 123[kkkkk",  []string{"жжж","123","kkkkk"}},
}

func TestSplitText(t *testing.T) {
	// fmt.Print("error1")
	// t.Error("error")
    for _, pair := range tests {
        v := SplitText(pair.raw)
        if fmt.Sprint(v) != fmt.Sprint(pair.ready) {
            t.Error(
                "For", pair.ready, 
                "expected", pair.ready,
                "got", v,
            )
        }
    }
}