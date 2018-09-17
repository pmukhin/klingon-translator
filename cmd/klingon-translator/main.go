package main

import (
	"fmt"
	"github.com/pmukhin/klingon-translator/pkg/klingon"
	"os"
	"strings"
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
	if len(args) < 2 {
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
		name := postfix
		if len(args) > 2 {
			name = strings.Join(args[1:], " ")
		}

		if err := klingon.Main(name); err != nil {
			fmt.Println(err.Error())
			os.Exit(-1)
		}
	}
}
