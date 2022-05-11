// Component interface
package composite

import "fmt"

type file struct {
	name string
}

func (f *file) search(keyword string) bool {
	fmt.Printf("file name is %s\n", f.name)

	return f.name == keyword
}

func (f *file) getName() string {
	return f.name
}
