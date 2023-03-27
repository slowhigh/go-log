// Concrete state
package state

import (
	"errors"
	"fmt"
)

type itemRequestedState struct {
	vendingMachine *vendingMachine
}

func (i *itemRequestedState) requestItem() error {
	return errors.New("item already requested")
}

func (i *itemRequestedState) addItem(count int) error {
	return errors.New("item Dispense")
}

func (i *itemRequestedState) insertMoney(money int) error {
	if money < i.vendingMachine.itemPrice {
		return fmt.Errorf("inserted money is less. please insert %d", i.vendingMachine.itemPrice)
	}

	fmt.Println("Money entered is ok")
	i.vendingMachine.setState(i.vendingMachine.hasMoney)
	return nil
}

func (i *itemRequestedState) dispenseItem() error {
	return errors.New("please insert money first")
}