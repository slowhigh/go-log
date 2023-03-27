// Concrete state
package state

import (
	"errors"
)

type noItemState struct {
	vendingMachine *vendingMachine
}

func (i *noItemState) requestItem() error {
	return errors.New("item out of stock")
}

func (i *noItemState) addItem(count int) error {
	i.vendingMachine.incrementItemCount(count)
	i.vendingMachine.setState(i.vendingMachine.hasItem)

	return nil
}

func (i *noItemState) insertMoney(money int) error {
	return errors.New("item out of stock")
}

func (i *noItemState) dispenseItem() error {
	return errors.New("item out of stock")
}
