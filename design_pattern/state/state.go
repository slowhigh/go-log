// State interface
package state

type state interface {
	addItem(int) error
	requestItem() error
	insertMoney(int) error
	dispenseItem() error
}