package composite


import (
	"fmt"

	compositeDomain "design_pattern/domain/composite"
	)


type Folder struct {
	Components []compositeDomain.Component
	Name       string
}

func (f *Folder) Search(keyword string) {
	fmt.Printf("Serching recursively for keyword %s in folder %s\n", keyword, f.Name)
	for _, composite := range f.Components {
		composite.Search(keyword)
	}
}

func (f *Folder) Add(c compositeDomain.Component) {
	f.Components = append(f.Components, c)
}