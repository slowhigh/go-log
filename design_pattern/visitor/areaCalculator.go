// Concrete visitor
package visitor

type areaCalculator struct{
	area float32
}

func (a *areaCalculator) visitForTriangle(t *triangle) {
	a.area = t.width * t.height / 2
}

func (a *areaCalculator) visitForCircle(c *circle) {
	a.area = c.radius * c.radius * 3.14
}

func (a *areaCalculator) visitForRectangle(r *rectangle) {
	a.area = r.width * r.height
}