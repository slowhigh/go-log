// Concrete element
package visitor

type triangle struct {
	width  float32
	height float32
}

func (t *triangle) accept(v visitor) {
	v.visitForTriangle(t)
}

func (t *triangle) getType() string {
	return "Triangle"
}
