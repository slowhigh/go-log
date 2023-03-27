// Element
package visitor

type Shape interface {
	getType() string
	accept(visitor)
}
