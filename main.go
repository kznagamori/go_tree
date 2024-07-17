package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var (
	directoriesOnly bool
	maxLevel        int
)

func init() {
	flag.BoolVar(&directoriesOnly, "d", false, "Display directories only")
	flag.IntVar(&maxLevel, "L", 0, "Specify the depth of directories to display")
}

func main() {
	flag.Parse()

	rootPath := "."
	if flag.NArg() > 0 {
		rootPath = flag.Arg(0)
	}

	printTree(rootPath, "", 0)
}

func printTree(path, indent string, level int) {
	if maxLevel != 0 && level > maxLevel {
		return
	}

	fileOrDir, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening directory:", err)
		return
	}
	defer fileOrDir.Close()

	entries, err := fileOrDir.Readdir(-1)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for i, entry := range entries {
		if directoriesOnly && !entry.IsDir() {
			continue
		}

		prefix := "├──"
		if i == len(entries)-1 {
			prefix = "└──"
		}

		fmt.Println(indent + prefix + " " + entry.Name())

		if entry.IsDir() {
			newIndent := indent
			if i != len(entries)-1 {
				newIndent += "│   "
			} else {
				newIndent += "    "
			}
			printTree(filepath.Join(path, entry.Name()), newIndent, level+1)
		}
	}
}

