package main

import (
	"fmt"
	prototype2 "prototype/samples/app/prototype"
	"prototype/samples/domain"
)

func main() {
	file1 := &prototype2.FileProto{Name: "File1"}
	file2 := &prototype2.FileProto{Name: "File2"}
	file3 := &prototype2.FileProto{Name: "File3"}

	folder1 := &prototype2.Folder{
		Children: []domain.Inode{file1},
		Name:     "Folder1",
	}

	folder2 := &prototype2.Folder{
		Children: []domain.Inode{folder1, file2, file3},
		Name:     "Folder2",
	}

	fmt.Println("\nPrinting hierarchy for Folder 2")
	folder2.Print("   ")

	cloneFolder := folder2.Clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")

	cloneFolder.Print("   ")
}
