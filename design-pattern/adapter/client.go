// Client code
package adapter

import "fmt"

type client struct {}

func (c *client) insertLightningConnectorIntoComputer(com interface{}) bool {
	fmt.Println("Client inserts Lightning connector into computer.")

	computer, ok := com.(computer)
	if !ok {
		return false;
	}

	computer.insertIntoLightningPort()
	return true
}