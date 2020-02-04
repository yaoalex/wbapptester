package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

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
		parseFile(filePath, *file)
	} else if os.IsNotExist(err) {
		fmt.Printf("Could not find the file at: %s\n", filePath)

	} else {
		fmt.Printf("Error looking for the file at: %s\n", filePath)
		log.Fatal(err)
	}
}

func buildTestFileName(file string) string {
	extension := filepath.Ext(file)
	testFileName := file[0:len(file)-len(extension)] + "_test.go"
	return testFileName
}

func parseFile(filePath, file string) {
	funcInfos, packageName := parseFunctions(filePath)
	if len(*funcInfos) > 0 {
		fmt.Println("Creating tests for the following http handlers:")
		for i, v := range *funcInfos {
			fmt.Printf("%d. %s\n", i+1, v.Name)
		}
		testFileName := buildTestFileName(file)
		_, err := os.Stat(testFileName)
		for err == nil {
			fmt.Printf("File already exists at the location: %s\n", testFileName)
			fmt.Println("Please enter a new location to store generated test file")
			fmt.Scanln(&testFileName)
			if _, err := os.Stat(testFileName); os.IsNotExist(err) {
				break
			}
		}
		err = generateTestFile(packageName, testFileName, funcInfos)
		if err != nil {
			fmt.Println("Error trying to create the test file")
			log.Fatal(err)
		}
		cmd := exec.Command("gofmt", "-w", testFileName)
		if _, err := cmd.CombinedOutput(); err != nil {
			fmt.Println("Failed to run gofmt on the test file")
			log.Fatal(err)
		}
		fmt.Printf("Successfully generate test code at: %s\n", testFileName)

	} else {
		fmt.Println("Could not find any http handler functions in the file")
	}
}
