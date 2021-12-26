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
	extension := flag.String("E", "*", "File extension to filter")
	substring := flag.String("S", "", "Substring to replace")
	replacement := flag.String("R", "", "Replacement string")
	flag.Parse()

	fmt.Println(path)
	fmt.Println(*extension)
	fmt.Println(*substring)
	fmt.Println(*replacement)

}
