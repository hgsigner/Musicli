package main

import (
	"flag"
	"fmt"
	"os"
)

var Usage = func() {
	fmt.Fprintf(os.Stderr, `
NAME:
  musicli - It does something related to music

USAGE:
  musicli bla bla bla

VERSION:
  %s

COMMANDS:
  help, h Shows a list of commands or help

EXAMPLES:
  musicli bla bla bla 

OPTIONS:
`, Version)
	flag.PrintDefaults()
}

var Version = "0.0.1"

func main() {

	version := flag.Bool("version", false, "display the version number")
	flag.Usage = Usage
	flag.Parse()

	if *version {
		fmt.Printf("version %s\n", Version)
		os.Exit(1)
	}

}
