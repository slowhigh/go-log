package clone

import "fmt"

type folder struct {
	name     string
	children []inode
}

func (f *folder) print(indentation string) {
	fmt.Println(indentation + f.name)

	for _, i := range f.children {
		i.print(indentation + indentation)
	}
}

// prototype pattern
func (f *folder) clone() inode {
	clonedFolder := &folder{name: f.name}

	var clonedChildren []inode
	for _, i := range f.children {
		clonedChild := i.clone()
		clonedChildren = append(clonedChildren, clonedChild)
	}
	clonedFolder.children = clonedChildren

	return clonedFolder
}

func (f *folder) equalsClone(n inode) bool {
	otherFolder, ok := n.(*folder)

	if !ok {
		return false
	}

	if f.name != otherFolder.name {
		return false
	}

	fmt.Printf("%p vs %p\n", f, otherFolder)

	if fmt.Sprintf("%p", f) == fmt.Sprintf("%p", otherFolder) {
		return false
	}

	if len(f.children) != len(otherFolder.children) {
		return false
	}

	for i := 0; i < len(f.children); i++ {
		if !f.children[i].equalsClone(otherFolder.children[i]) {
			return false
		}
	}

	return true
}
