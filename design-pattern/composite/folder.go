// Composite
package composite

import "fmt"

type folder struct {
	components []component
	name string
}

func (f *folder) search(keyword string) bool {
	fmt.Printf("foler name is %s\n", f.name)

	for _, composite := range f.components {
        if composite.search(keyword) {
			return true
		}
    }

	return false
}

func (f *folder) add(c component) {
	f.components = append(f.components, c)
}

