package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseLists(t *testing.T) {
	byte, err := ioutil.ReadFile("lists_test.html")
	if err != nil {
		panic(err)
	}

	parserResult := ParseLists(byte)
	for _, item := range parserResult.Items {
		t.Logf("Parser result: %v\n", item)
	}

}
