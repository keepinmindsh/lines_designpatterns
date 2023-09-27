package main

import (
	"fmt"
	"prototype/app/prototype"
	"prototype/domain"
)

func main() {
	file1 := &prototype.FileProto{Name: "File1"}
	file2 := &prototype.FileProto{Name: "File2"}
	file3 := &prototype.FileProto{Name: "File3"}

	folder1 := &prototype.Folder{
		Children: []domain.Inode{file1},
		Name:     "Folder1",
	}

	folder2 := &prototype.Folder{
		Children: []domain.Inode{folder1, file2, file3},
		Name:     "Folder2",
	}

	fmt.Println("\nPrinting hierarchy for Folder 2")
	folder2.Print("   ")

	cloneFolder := folder2.Clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")

	cloneFolder.Print("   ")
}
