package file

import (
	"design_pattern/domain/prototype"
	"fmt"
)

type File struct {
	Name string
}

func (f *File) Print(indentation string) {
	fmt.Println(indentation + f.Name)
}

func (f *File) Clone() prototype.Node {
	return &File{Name: f.Name + "_clone"}
}
