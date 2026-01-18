package main

import "fmt"

func buildProject() {
	fmt.Println("building...")

	//get project config info
	config := parseConfigFile("n.config")
	fmt.Println("module:", config["module"])
	fmt.Println("n-version:", config["n-version"])
}
