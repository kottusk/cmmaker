package main

import (
	"cmmaker/helper"
	"fmt"
	"os"
)

var genre string
var version string
var db = "mysql"

func main() {
	fmt.Println("Start!")
	helper.HelperSign()
	if !parameterCheck() {
		helper.HelperScan()
	}
}

func parameterCheck() bool {
	argsWithProg := os.Args
	if len(argsWithProg) < 3 {
		helper.HelperScan()
	} else {
		genre = argsWithProg[1]
		version = argsWithProg[2]
		if len(argsWithProg) > 3 {
			db = argsWithProg[3]
		}
	}
	fmt.Printf("Parameters are: genre=%s, version=%s, db=%s\n", genre, version, db)
	if len(genre) > 0 && len(version) > 1 && len(db) > 1 {
		return true
	} else {
		return false
	}
}
