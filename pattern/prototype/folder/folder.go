package folder

import "fmt"
import "design_pattern/domain/prototype"

type Folder struct {
	Children []prototype.Node
	Name     string
}

func (f *Folder) Print(indentation string) {
	fmt.Println(indentation + f.Name)
	for _, i := range f.Children {
		i.Print(indentation + indentation)
	}
}

func (f *Folder) Clone() prototype.Node {
	cloneFolder := &Folder{Name: f.Name + "_clone"}
	var tempChildren []prototype.Node
	for _, i := range f.Children {
		copyNose := i.Clone()
		tempChildren = append(tempChildren, copyNose)
	}
	cloneFolder.Children = tempChildren
	return cloneFolder
}

