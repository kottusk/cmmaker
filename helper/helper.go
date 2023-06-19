package helper

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

var Genre string
var Version string
var Db string

func HelperSign() {
	fmt.Println("XXXXXXXXXXXXXXX")
	fmt.Println("XXXXXXXXXXXXXXX")
}

func HelperScan() (string, string, string) {
	fmt.Println("Please give parameters")
	fmt.Println("Please give genre (trunk/branch/tag)")
	fmt.Scan(&Genre)
	fmt.Println("Please give version ")
	fmt.Scan(&Version)
	fmt.Println("Please give db (mysql/oracle/mssql/postgress)")
	fmt.Scan(&Db)
	return Genre, Version, Db
}

func CheckoutAndBuild(path string, directory string) bool {
	//cmd := exec.Command("svn", "checkout", path, directory)
	//err := cmd.Run()
	//// if there is an error with our execution
	//// handle it here
	//if err != nil {
	//	fmt.Printf("%s", err)
	//}
	ppath := filepath.Join("core", "pom.xml")
	fmt.Printf("Path %s", ppath)
	out, err := exec.Command("mvn", "clean", "install", "-f", ppath).Output()
	//err := cmd.Run()
	// if there is an error with our execution
	// handle it here
	if err != nil {
		fmt.Printf("%s", err)
	}
	output := string(out[:])
	fmt.Println(output)

	return false
}
