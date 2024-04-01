package main

import "fmt"

type File struct {
	name string
}

func (f *File) print(ind string) {
	fmt.Println(ind + f.name)
}

func (f *File) clone() Inode {
	return &File{name: f.name + "_clone"}
}
