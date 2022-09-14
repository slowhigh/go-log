// Concrete visitor
package visitor

import "fmt"

type infoPrinter struct {
	print string
}

func (i *infoPrinter) visitForTriangle(t *triangle) {
	i.print = fmt.Sprint(t.width * t.height / 2)
}

func (i *infoPrinter) visitForCircle(c *circle) {
	i.print = fmt.Sprint(c.radius * c.radius * 3.14)
}

func (i *infoPrinter) visitForRectangle(r *rectangle) {
	i.print = fmt.Sprint(r.width * r.height)
}
