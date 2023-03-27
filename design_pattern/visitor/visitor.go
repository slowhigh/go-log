// Visitor
package visitor

type visitor interface {
	visitForTriangle(*triangle)
	visitForCircle(*circle)
	visitForRectangle(*rectangle)
}
