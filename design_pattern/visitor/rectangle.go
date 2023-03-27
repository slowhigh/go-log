// Concrete element
package visitor

type rectangle struct {
	width  float32
	height float32
}

func (r *rectangle) accept(v visitor) {
	v.visitForRectangle(r)
}

func (r *rectangle) getType() string {
	return "rectangle"
}
