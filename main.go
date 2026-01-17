package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// n version
const VERSION = "0.1.0"

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("n", VERSION)
		fmt.Println("n programming language compiler.")
		fmt.Println("For help:")
		fmt.Println("  n -h")
	} else if len(args) == 1 {
		switch args[0] {
		case "--version", "-v":
			fmt.Println("n", VERSION)
		case "init":
			initProject()
		case "--help", "-h":
			fmt.Println("n compiler commands:")
			fmt.Println("   n               -  show about message")
			fmt.Println("   n -v,--version  -  show n version")
			fmt.Println("   n -h,--help     -  show compiler commands list")
		}
	}
}

func initProject() {
	module_name := filepath.Base(func() string { d, _ := os.Getwd(); return d }())

	config_file_name := "n.mod"
	config_content := "module " + module_name + "\n\n" +
		"n " + VERSION + "\n"

	config_file, err := os.OpenFile(config_file_name, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)

	if err != nil {
		if os.IsExist(err) {
			fmt.Printf("Error: %s file already exists!\n", config_file_name)
			os.Exit(1)
		}
		fmt.Println("Error creating file:\n", config_file_name)
		os.Exit(1)
	}
	defer config_file.Close()

	_, err = config_file.WriteString(config_content)
	if err != nil {
		fmt.Println("Error writing to file:", config_file_name)
		os.Exit(1)
	}

	fmt.Printf("n: created new %s: module %s\n", config_file_name, module_name)

}
