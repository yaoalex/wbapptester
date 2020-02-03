package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	app := kingpin.New("webapptester", "Generate boilerplate code to test your HTTP handlers")
	file := app.Arg("file", "Go file you would like to create tests for.").Required().String()
	kingpin.MustParse(app.Parse(os.Args[1:]))
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	filePath := dir + "\\" + *file

	if _, err := os.Stat(filePath); err == nil {
		fmt.Printf("Creating tests for the file at: %s\n", filePath)
		parseFunctions(filePath)

	} else if os.IsNotExist(err) {
		fmt.Printf("Could not find the file at: %s\n", filePath)

	} else {
		fmt.Printf("Error looking for the file at: %s\n", filePath)
	}

}
