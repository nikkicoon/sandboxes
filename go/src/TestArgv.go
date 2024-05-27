package main

import (
	"github.com/cosiner/argv"
	"testing"
	"os"
	"reflect"
)

func TestArgv(t *testing.T) {
	args, err := argv.Argv([]rune(" ls   `echo /`   |  wc  -l "), os.Environ(), argv.Run)
	if err != nil {
		t.Fatal(err)
	}
	expects := [][]string {
		[]string{"ls", "/"},
		[]string{"wc", "-l"},
	}
	if !reflect.DeepEqual(args, expects) {
		t.Fatal(args)
	}
}
