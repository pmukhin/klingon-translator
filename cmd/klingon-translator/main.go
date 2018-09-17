package main

import (
	"fmt"
	"github.com/pmukhin/klingon-translator/pkg/klingon"
	"os"
)

var (
	AppName    = "klingon"
	AppVersion = "0.0.1"
)

func usage() {
	fmt.Println("Usage:", AppName, "<charachterName>")
}

func version() {
	fmt.Println("Version: ", AppVersion)
}

func main() {
	args := os.Args
	// no args then usage
	if len(args) != 2 {
		usage()
		return
	}

	postfix := args[1]
	switch postfix {
	case "-v", "-version", "version":
		version()
	case "-h", "-help", "help", "usage":
		usage()
	default:
		err := klingon.Main(postfix)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(-1)
		}
	}
}
