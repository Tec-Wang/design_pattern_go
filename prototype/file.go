package prototype

import "fmt"

type File struct {
	name string
}

func (f *File) Print(indentation string) {
	fmt.Println(indentation + "My Name Is  " + f.name)
}

func (f *File) Clone() Inode {
	return &File{
		name: f.name + "_clone",
	}
}
