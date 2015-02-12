package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	ApiRoot = "http://developer.echonest.com/api/v4/artist"
)

var Usage = func() {
	fmt.Fprintf(os.Stderr, `
NAME:
  musicli - It does something related to music

USAGE:
  musicli [--options]

VERSION:
  %s

COMMANDS:
  help, h Shows a list of commands or help

EXAMPLES:
  musicli -artist="Paul Mccartney"

OPTIONS:
`, Version)
	flag.PrintDefaults()
}

var Version = "0.0.1"

func Run(artist string) {

	urls, err := FetchUrls(artist)
	if err != nil {
		log.Println("Ops. No urls")
	}

	fmt.Printf("%v\n", urls)

}

func main() {

	version := flag.Bool("version", false, "display the version number")
	artist := flag.String("artist", "", "artist's name for searching")

	flag.Usage = Usage
	flag.Parse()

	if *version {
		fmt.Printf("version %s\n", Version)
		os.Exit(1)
	}

	if *artist != "" {
		fmt.Fprintf(os.Stdout, "You have selected the artist: %s\n", *artist)
		Run(*artist)
	} else {
		fmt.Println("You should enter an artist name in order to proceed.")
	}

}
