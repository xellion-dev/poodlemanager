package main

import (
	"flag"
	"fmt"
)

func main() {
	//var pkgname string
	var install bool
	flag.BoolVar(&install, "install", false, "Publish the article")
	flag.Parse()
	if install == true {
		fmt.Printf("PoodleManager Version 1.0.0\n===================================================================================\nNo Package Specified.\n")
	} else {
		fmt.Printf("PoodleManager Version 1.0.0\n===================================================================================\nNo Option Specified. Since you did not specify an option we will provide you some basic knoledge.\nThe basic commands are:\n Install - Installs package\n Remove - Removes package\n Uninstall - Same thing as remove but purges all directories made by the install\n Update - Updates package list and upgrades packages + dependencies.\n")
	}
	//fmt.Scan(&pkgname)
	//fmt.Println("Package is", pkgname)
}
