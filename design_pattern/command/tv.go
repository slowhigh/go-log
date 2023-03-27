// Concrete recevier
package command

type tv struct {
	isRunning bool
}

func (t *tv) on() {
	t.isRunning = true
}

func (t *tv) off() {
	t.isRunning = false
}