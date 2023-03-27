// Concrete element
package visitor

type circle struct {
	radius float32
}

func (c *circle) accept(v visitor) {
	v.visitForCircle(c)
}

func (c *circle) getType() string {
	return "Circle"
}
