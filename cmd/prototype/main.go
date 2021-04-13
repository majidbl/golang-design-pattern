package main

import (
	"fmt"

	"design_pattern/pattern/prototype/file"
	"design_pattern/pattern/prototype/folder"
	"design_pattern/domain/prototype"
)

func main() {
	file1 := &file.File{Name: "File1"}
	file2 := &file.File{Name: "File2"}
	file3 := &file.File{Name: "File3"}

	folder1 := &folder.Folder{
		Children: []prototype.Node{file1},
		Name:      "Folder1",
	}

	folder2 := &folder.Folder{
		Children: []prototype.Node{folder1, file2, file3},
		Name:      "Folder2",
	}
	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.Print("  ")

	cloneFolder := folder2.Clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.Print("  ")
}
