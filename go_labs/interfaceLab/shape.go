package interface_lib

type Shape interface {
	GetAreaSize() int
	AddValue(value int)
}

type Rectangle struct {
	width int
	height int
}

type RightTriangle struct {
	width int
	height int
}

func (rect Rectangle) GetAreaSize() int {
	return (rect.width * rect.height)
}

func (tri RightTriangle) GetAreaSize() int {
	return (tri.width * tri.height / 2)
}
