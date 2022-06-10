// Concrete visitor
package visitor

import "fmt"

type infoPrinter struct{
	print string
}

func (i *infoPrinter) visitForTriangle(t *triangle) {
	i.print = fmt.Sprintf() t.width * t.height / 2
}

func (i *infoPrinter) visitForCircle(c *circle) {
	i.print = c.radius * c.radius * 3.14
}

func (i *infoPrinter) visitForRectangle(r *rectangle) {
	i.print = r.width * r.height
}