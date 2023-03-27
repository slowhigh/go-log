// Concrete component
package decorator

type veggieMania struct{}

func (p *veggieMania) getPrice() int {
	return 15
}
