package prototype

import (
	"fmt"
	"prototype/domain"
)

type FileProto struct {
	Name string
}

func (f *FileProto) Print(indentation string) {
	fmt.Println(indentation + f.Name)
}

func (f *FileProto) Clone() domain.Inode {
	return &FileProto{Name: f.Name}
}
