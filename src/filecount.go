package main

import (
	"fmt"
	"log"
	"path/filepath"
)

func main() {
	fmt.Println("Application Ran")
	files, err := filepath.Glob("*")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("File Count:", len(files))
	fmt.Println(files) // contains a list of all files in the current directory
}
