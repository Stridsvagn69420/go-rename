package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	// presets
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	// set cli flags
	flag.StringVar(&path, "D", path, "Directory of files to rename")
	flag.String("E", "*", "File extension to filter")
	flag.String("S", "", "Substring to replace")
	flag.String("R", "", "Replacement string")
	flag.Parse()

	fmt.Print(path)
}
