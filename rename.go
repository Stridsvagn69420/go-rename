package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// presets
	path, err := os.Getwd()
	if err != nil {
		path = "./"
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
		os.Stderr.WriteString("\033[1;31mSubstring not set but required!\033[0m\n")
		errorFlags = true
	}
	if errorFlags {
		fmt.Print("To get help, run: \033[1;32mrename -h\033[0m")
		os.Exit(1)
	}

	// read files
	files, err := ioutil.ReadDir(path)
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
	}
	files = filterFiles(files, *extension)

	// rename files
	for _, file := range files {
		fmt.Print(file.Name())
	}

	// exit normally
	os.Exit(0)
}

func filterFiles(files []os.FileInfo, extension string) []os.FileInfo {
	filteredFiles := make([]os.FileInfo, 0)
	for _, file := range files {
		if !file.IsDir() {
			if extension != "*" {
				if file.Name()[len(file.Name())-len(extension):] == extension {
					filteredFiles = append(filteredFiles, file)
				}
			} else {
				filteredFiles = append(filteredFiles, file)
			}
		}
	}
	return filteredFiles
}
