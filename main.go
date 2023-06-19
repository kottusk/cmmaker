package main

import (
	"cmmaker/helper"
	"fmt"
	"os"
)

var Genre string
var Version string
var Db = "mysql"

func main() {
	fmt.Println("Start!")
	helper.HelperSign()
	if !parameterCheck() {
		Genre, Version, Db = helper.HelperScan()
	}
	for !configCheck() {
		Genre, Version, Db = helper.HelperScan()
	}
	helper.CheckoutAndBuild("http://svn.consol.de/svn/cm/cmas/core/branches/6.16-SNAPSHOT", "core")
}

func parameterCheck() bool {
	argsWithProg := os.Args
	if len(argsWithProg) < 3 {
		Genre, Version, Db = helper.HelperScan()
	} else {
		Genre = argsWithProg[1]
		Version = argsWithProg[2]
		if len(argsWithProg) > 3 {
			Db = argsWithProg[3]
		}
	}
	fmt.Printf("Parameters are: genre=%s, version=%s, db=%s\n", Genre, Version, Db)
	if len(Genre) > 1 && len(Version) > 1 && len(Db) > 1 {
		return true
	} else {
		return false
	}
}

func configCheck() bool {
	if len(Genre) > 1 && len(Version) > 1 && len(Db) > 1 {
		return true
	} else {
		return false
	}
}
