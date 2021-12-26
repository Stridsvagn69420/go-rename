package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// presets
	path, err := os.Getwd()
	if err != nil {
		path = "."
	}
	// set cli flags
	flag.StringVar(&path, "D", path, "Directory of files to rename")
	extension := flag.String("E", "*", "File extension to filter")
	substring := flag.String("S", "", "Substring to replace")
	replacement := flag.String("R", "\u200b", "Replacement string\n")
	flag.Parse()

	// check if substring and replacement are set
	errorFlags := false
	if *replacement == "\u200b" {
		os.Stderr.WriteString("Although the replacement string can be empty, it's still required to be explicitly set!\nThis is to prevent accidental renames.\n")
		errorFlags = true
	}
	if *substring == "" {
		os.Stderr.WriteString("\033[1;31mSubstring not set but required!\033[0m")
		errorFlags = true
	}
	if errorFlags {
		os.Exit(1)
	}

	fmt.Println(path)
	fmt.Println(*extension)
	fmt.Println(*substring)
	fmt.Println(*replacement)

}
