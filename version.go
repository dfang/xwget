package main

import (
	"io/ioutil"
	"log"
)

var xwgetVersion string
var goVersion string
var buildTimestamp string
var repo string

func init() {
	repo = "https://github.com/dfang/xwget"
	goVersion = "go 1.14.0"
	readVersion()
}

func readVersion() {
	content, err := ioutil.ReadFile("version.txt")
	if err != nil {
		log.Fatal(err)
	}
	xwgetVersion = string(content)
}
