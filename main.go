package main

import (
	"flag"
	"fmt"
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

CATEGORIES:
  urls

COMMANDS:
  help, h Shows a list of commands or help

EXAMPLES:
  musicli -artist="Paul Mccartney" -category="urls"

OPTIONS:
`, Version)
	flag.PrintDefaults()
}

var Version = "0.0.1"
var Categories = []string{
	"urls",
}

func Run(artist, category string) {

	switch category {
	case "urls":
		RunUrls(artist)
	}

}

func main() {

	version := flag.Bool("version", false, "display the version number")
	artist := flag.String("artist", "", "artist's name for searching")
	category := flag.String("category", "", "categories to search for")

	flag.Usage = Usage
	flag.Parse()

	if *version {
		fmt.Printf("version %s\n", Version)
		os.Exit(1)
	}

	if ok := SliceContains(Categories, *category); *artist != "" && ok {
		fmt.Fprintf(os.Stdout, "You have selected the artist %s and the category %s.\n", *artist, *category)
		Run(*artist, *category)
	} else {
		fmt.Fprintln(os.Stdout, "WARNING: You should enter an artist name and a category in order to proceed.\nType musicli --help in order to get the availables categories.")
	}

}
