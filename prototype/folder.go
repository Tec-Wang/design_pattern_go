package prototype

import "fmt"

type Folder struct {
	name     string
	childern []Inode
}

func (f *Folder) Print(indentation string) {
	fmt.Println(indentation + "My Name Is  " + f.name)
	for _, child := range f.childern {
		child.Print(indentation)
	}
}

func (f *Folder) Clone() Inode {
	newChildern := make([]Inode, 0, len(f.childern))
	for _, child := range f.childern {
		newChildern = append(newChildern, child.Clone())
	}

	return &Folder{
		name:     f.name + "_clone",
		childern: newChildern,
	}

}
