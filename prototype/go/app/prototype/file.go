package prototype

import (
	"fmt"
	"prototype/domain"
)

type File struct {
	name string
}

func (f *File) print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *File) clone() domain.Inode {
	return &File{name: f.name + "_clone"}
}
