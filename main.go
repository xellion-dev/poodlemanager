package main

import (
	"fmt"
)

func main() {
	var pkgname string

	fmt.Print("PoodleManager Version 1.0.0")
	fmt.Scan(&pkgname)
	fmt.Println("Package is", pkgname)
}
