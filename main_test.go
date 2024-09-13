package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

var exampleOK = `1
2
3
4
5`

func TestOk(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(exampleOK))
	out := new(bytes.Buffer)
	err := unig(in, out)
	if err != nil {
		t.Errorf("TestOk Failed")
	}
}
