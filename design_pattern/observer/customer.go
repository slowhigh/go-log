// Concrete observer
package observer

import "fmt"

type customer struct {
	id string
	inbox []string
}

func (c *customer) update(itemName string) {
	mail := fmt.Sprintf("Sending email to customer %s for item %s\n", c.id, itemName)
	fmt.Print(mail)
	c.inbox = append(c.inbox, mail)
}

func (c *customer) getID() string {
	return c.id
}