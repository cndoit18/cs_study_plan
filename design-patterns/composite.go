package main

import "fmt"

type Component interface {
	Search(string)
}

type File struct {
	name string
}

func (f *File) Search(keyword string) {
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.name)
}

type Folder struct {
	components []Component
	name       string
}

func (f *Folder) add(c ...Component) {
	f.components = append(f.components, c...)
}

func (f *Folder) Search(keyword string) {
	fmt.Printf("Searching for keyword %s in Folder %s\n", keyword, f.name)

	for _, component := range f.components {
		component.Search(keyword)
	}
}
