package main

import (
	"fmt"
	"os"
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
		case "--help", "-h":
			fmt.Println("n compiler commands:")
			fmt.Println("   n               -  show about message")
			fmt.Println("   n -v,--version  -  show n version")
			fmt.Println("   n -h,--help     -  show compiler commands list")
		}
	}
}
