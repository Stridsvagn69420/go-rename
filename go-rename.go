package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func main() {
	// presets
	dir, err := os.Getwd()
	if err != nil {
		dir = "./"
	}
	// set cli flags
	flag.StringVar(&dir, "D", dir, "Directory of files to rename")
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
		os.Stderr.WriteString("\033[31mSubstring not set but required!\033[0m\n")
		errorFlags = true
	}
	if errorFlags {
		fmt.Print("To get help, run: \033[32mgo-rename -h\033[0m\n")
		os.Exit(1)
	}

	// read files
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
	}
	files = filterFiles(files, *extension)

	// rename files
	fmt.Println("Renaming files in \033[1m" + dir + "\033[0m...")
	ranOnce := false
	for _, file := range files {
		newName := strings.ReplaceAll(file.Name(), *substring, *replacement)
		if newName != file.Name() {
			if !ranOnce {
				ranOnce = true
			}
			err := os.Rename(
				path.Join(dir, file.Name()),
				path.Join(dir, newName),
			)
			if err != nil {
				os.Stderr.WriteString(err.Error() + "\n")
				os.Exit(2)
			}
			fmt.Println("\033[31m" + file.Name() + "\033[0m => \033[32m" + newName + "\033[0m")
		}
	}
	if !ranOnce {
		fmt.Println("No files replaced.")
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
