# rename
Simple command line tool to rename all files in a directory by extension with given substring and replacer

# Installation
## Go SDK (recommended)
If you already have `go` installed, you can easily just clone the repository and run `go install`. Go will compile and install the project for you. The executable will be located in `~/go/bin`, so make sure the folder is inside your PATH (it automatically should be on Windows).  
You can dwonload the Go SDK [here](https://go.dev/dl/).

## Download
Alternatively, you can download the [latest release](https://github.com/Stridsvagn69420/rename/releases) for your platform and manually put it in a directory listed in your path variable.

# Usage
[optional] `-D`:  
The directory of the files you want to rename (defaults to currect directory), e.g. `/home/joe/mama`, `C:\Users\Stuff\Pictures`

[optional] `-E`:  
The file extension to filter (default: `*` (all files)), e.g. `webm`, `.jpg`, `tar.gz`, `.tar.xz`

`-S`:  
The substring/section you want to rename, e.g. `Bloat` (to be removed), `this`, `.m4v` (file extension)

`-R`:  
The replacement for the substring, e.g. `​` (empty string), `that`, `.mp4`
