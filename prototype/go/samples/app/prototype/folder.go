package prototype

import (
	"fmt"
	"prototype/samples/domain"
)

type Folder struct {
	Children []domain.Inode
	Name     string
}

func (f *Folder) Print(indentation string) {
	fmt.Println(indentation + f.Name)
	for _, i := range f.Children {
		i.Print(indentation + indentation)
	}
}

func (f *Folder) Clone() domain.Inode {
	cloneFolder := &Folder{Name: f.Name + "_clone"}
	var tempChildren []domain.Inode
	for _, i := range f.Children {
		copy := i.Clone()
		tempChildren = append(tempChildren, copy)
	}
	cloneFolder.Children = tempChildren
	return cloneFolder
}