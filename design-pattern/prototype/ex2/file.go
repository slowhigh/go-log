package clone

import "fmt"

type file struct {
	name string
}

func (f *file) print(indentation string) {
	fmt.Println(indentation + f.name)
}

// prototype pattern
func (f *file) clone() inode {
	return &file{name: f.name}
}

func (f *file) equalsClone(n inode) bool {
	otherFile, ok := n.(*file)

	if !ok {
		return false
	}

	if f.name != otherFile.name {
		return false
	}

	fmt.Printf("%p vs %p \n", f, otherFile)

	if fmt.Sprintf("%p", f) == fmt.Sprintf("%p", otherFile) {
		return false
	}

	return true
}
