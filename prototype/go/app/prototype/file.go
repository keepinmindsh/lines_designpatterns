package prototype

import (
	"fmt"
)

type FileProto struct {
	Name string
}

func (f *FileProto) print(indentation string) {
	fmt.Println(indentation + f.Name)
}

/*
// todo 해당 코드는 정상적으로 동작하지 않음
func (f *FileProto) clone() domain.Inode {
	return &FileProto{Name: f.Name}
}
*/

func Clone() *FileProto {
	fileProto := &FileProto{Name: "Test"}

	return fileProto
}
