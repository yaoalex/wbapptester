package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	app := kingpin.New("webapptester", "Generate tests for your HTTP handler functions")
	file := app.Arg("file", "Go file you would like to create tests for.").Required().String()
	kingpin.MustParse(app.Parse(os.Args[1:]))
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fullpath := dir + "\\" + *file
	if _, err := os.Stat(fullpath); err == nil {
		fmt.Printf("Would create tests for the file: %s", fullpath)

	} else if os.IsNotExist(err) {
		fmt.Printf("Could not find the file; %s", fullpath)

	} else {
		fmt.Printf("Error looking for the file; %s", fullpath)
	}

}
