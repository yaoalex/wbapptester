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
		funcInfos, packageName := parseFunctions(filePath)
		if len(funcInfos) > 0 {
			fmt.Println("Creating tests for the following http handlers:")
			for i, v := range funcInfos {
				fmt.Printf("%d. %s\n", i+1, v.Name)
			}
			generateTestFile(packageName, filePath, funcInfos)
		} else {
			fmt.Println("Could not find any http handler functions in the file")
		}
	} else if os.IsNotExist(err) {
		fmt.Printf("Could not find the file at: %s\n", filePath)

	} else {
		fmt.Printf("Error looking for the file at: %s\n", filePath)
	}

}
